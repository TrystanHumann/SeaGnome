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
	// type matchIDInsertResult struct {
	// 	ID int64 `db:"returnID"`
	// }
	ctx, ctxCancel := context.WithTimeout(r.Context(), time.Second*1800000)
	defer ctxCancel()
	switch r.Method {
	case http.MethodPost:
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

		// ID of the event that we will be updating data for
		eventID := r.FormValue("eventID")

		id, err := strconv.Atoi(eventID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// transfer contents of the file to our buffer
		io.Copy(&buffer, file)

		var games []types.Match
		gamesCache := make(map[string]int)
		getMatchQuery := "select * from public.getmatchesbyevent($1::int4);"
		getGameIDQuery := "select * from public.get_gameid_sp($1::text);"
		getCompetitorQuery := "select * from public.get_competitorid_sp($1::text);"
		insertMatchQuery := "select * from public.insert_match_sp($1::int4,$2::int4,$3::timestamp,$4::int4)"
		err = u.Data.SelectContext(ctx, &games, getMatchQuery, eventID)
		if err != nil {
			fmt.Println(err)
		}

		//sent games get gameID
		for _, game := range games {
			var gameID []int
			err = u.Data.SelectContext(ctx, &gameID, getGameIDQuery, game.Game)
			if err != nil {
				fmt.Println(err)
			}
			if len(gameID) != 0 {
				gamesCache[strings.Trim(game.Game, " ")] = gameID[0]
			}

		}

		// Determine whether it is prediction or results
		// Getting the string version of our buffer
		contents := strings.Split(buffer.String(), "\n")

		if len(contents) <= 0 {
			http.Error(w, "No contents found in CSV", http.StatusBadRequest)
			return
		}
		firstRow := strings.Split(contents[0], ",")
		for l, data := range contents {
			// split by commas and begin extracting the values
			rowData := strings.Split(data, ",")
			if l != 0 {
				for i, col := range rowData {
					if firstRow[i] == "Winner" {
						var competitorID []int
						var matchID int
						if len(col) == 0 {
							continue
						}
						err := u.Data.SelectContext(ctx, &competitorID, getCompetitorQuery, col)
						if err != nil {
							fmt.Println(err)
						}
						err = u.Data.GetContext(ctx, &matchID, insertMatchQuery, eventID, gamesCache[trimGame(rowData[0])], time.Now(), competitorID[0])
						if err != nil {
							fmt.Println(err)
						}
					}
				}
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
