package handlers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/trystanhumann/SeaGnome/Backend/types"
)

// UploadResults : Handles requests involved with uploads
type UploadResults struct {
	Data *sqlx.DB
}

// ServeHTTP : Listens for a request and creates a response
func (u *UploadResults) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	ctx, ctxCancel := context.WithTimeout(r.Context(), time.Second*1800000)
	defer ctxCancel()
	switch r.Method {
	case http.MethodPut:
		var buffer bytes.Buffer

		file, header, err := r.FormFile("uploadFile")

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fn := strings.Split(header.Filename, ".")
		if len(fn) <= 0 {
			http.Error(w, "Invalid file type", http.StatusBadRequest)
			return
		}

		if fn[len(fn)-1] != "csv" {
			http.Error(w, "Invalid file type.  File must be a csv file", http.StatusBadRequest)
			return
		}

		// transfer contents of the file to our buffer
		io.Copy(&buffer, file)
		// Determine whether it is prediction or results
		// Getting the string version of our buffer
		contents := strings.Split(buffer.String(), "\n")

		if len(contents) < 2 {
			http.Error(w, "No contents found in CSV", http.StatusBadRequest)
			return
		}

		// ID of the event that we will be updating data for
		eventID := r.FormValue("eventID")

		_, err = strconv.Atoi(eventID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		const GAMENAME = "Game Name"
		const WINNER = "Winner"
		const SCHEDULEDDATE = "Scheduled Date"

		var games []types.MatchesForResults
		gameCache := make(map[string]int)
		compCache := make(map[int]map[string]int)
		cols := make(map[string]int)

		getMatchQuery := "select * from public.getmatchesbyeventforresults($1::int4);"
		insertMatchQuery := "select * from public.insert_match_sp($1::int4,$2::int4,$3::timestamp,$4::int4)"

		err = u.Data.SelectContext(ctx, &games, getMatchQuery, eventID)
		if err != nil {
			fmt.Println(err)
		}

		for index := range games {
			gameCache[games[index].Game] = games[index].GameID
			compMap, ok := compCache[gameCache[games[index].Game]]
			if !ok {
				compMap = make(map[string]int)
				compCache[gameCache[games[index].Game]] = compMap
			}
			compMap[games[index].Competitor] = games[index].CompetitorID
		}

		firstRow := strings.Split(contents[0], ",")
		for i := range firstRow {
			if strings.EqualFold(GAMENAME, firstRow[i]) {
				cols[GAMENAME] = i
			} else if strings.EqualFold(WINNER, firstRow[i]) {
				cols[WINNER] = i
			} else if strings.EqualFold(SCHEDULEDDATE, firstRow[i]) {
				cols[SCHEDULEDDATE] = i
			}
		}

		if len(cols) < 3 {
			http.Error(w, "invalid upload format", http.StatusBadRequest)
		}

		for index := 1; index < len(contents); index++ {
			// split by commas and begin extracting the values
			rowData := strings.Split(contents[index], ",")
			game := trimGame(rowData[cols[GAMENAME]])
			gameID := gameCache[game]

			if gameID == 0 {
				fmt.Println("game not found during results upload: " + game)
				continue
			}

			_, err = u.Data.ExecContext(ctx, insertMatchQuery,
				eventID, // event id
				gameID,  // game name
				rowData[cols[SCHEDULEDDATE]],             // schedule date
				compCache[gameID][rowData[cols[WINNER]]]) // winner id

			if err != nil {
				fmt.Println(err)
			}
		}

		// Cleaning up buffer memory
		buffer.Reset()
		w.Write([]byte("Yes"))

	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func trimGame(game string) string {
	game = strings.Replace(game, "Predictions [", "", -1)
	game = strings.Replace(game, "]", "", -1)
	game = strings.TrimSpace(game)
	return game
}
