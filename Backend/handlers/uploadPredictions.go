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
)

// UploadPredictions : Handles requests involved with uploads
type UploadPredictions struct {
	Data *sqlx.DB
}

// ServeHTTP : Listens for a request and creates a response
func (u *UploadPredictions) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

		fmt.Printf("%i\n", id)

		// transfer contents of the file to our buffer
		io.Copy(&buffer, file)

		// Determine whether it is prediction or results
		// Getting the string version of our buffer
		contents := strings.Split(buffer.String(), "\n")

		if len(contents) <= 0 {
			http.Error(w, "No contents found in CSV", http.StatusBadRequest)
			return
		}
		columnNames := strings.Split(contents[0], ",")
		firstRow, gamesCache := getGames(columnNames)
		fmt.Println(firstRow)
		var gameIDSlice []int
		var userIDSlice []int
		competitorCache := make(map[string]int)
		matchIDCache := make(map[string]int)
		insertGameQuery := "select * from public.insert_game_sp($1::text);"
		insertUserQuery := "select * from public.insert_or_update_user_sp($1::text,$2::text);"
		insertMatchQuery := "select * from public.insert_match_sp($1::int4,$2::int4);"
		insertCompetitorQuery := "select * from public.insert_competitor_sp($1::text);"
		insertParticipantQuery := "select * from public.insert_participant_sp($1::int4,$2::int4);"
		insertPredictionQuery := "select * from public.insert_prediction_sp($1::int4,$2::int4);"

		for _, game := range gamesCache {
			// get game Id from hitting postgres
			// fmt.Println(game)
			var id int
			rows, err := u.Data.QueryContext(ctx, insertGameQuery, game)
			if err != nil {
				fmt.Println(err)
			}
			if rows.Next() {
				rows.Scan(&id)
			}
			rows.Close()
			// fmt.Println(id)
			gameIDSlice = append(gameIDSlice, id)

			//upload a match
			rows, err = u.Data.QueryContext(ctx, insertMatchQuery, eventID, id)
			if err != nil {
				fmt.Println(err)
			}
			if rows.Next() {
				rows.Scan(&id)
			}
			rows.Close()
			if matchIDCache[game] == 0 {
				matchIDCache[game] = id
			}
		}

		columnNames = strings.Split(contents[0], ",")
		for l, v := range contents {
			currentRowSplit := strings.Split(v, ",")
			//skip first row cuz its whack
			if l != 0 {
				var userID int
				var competitorID int
				var participantID int
				var predictionID int
				// user created(replace with postgres)
				rows, err := u.Data.QueryContext(ctx, insertUserQuery, strings.ToLower(currentRowSplit[1]), strings.ToLower(currentRowSplit[2]))

				if err != nil {
					fmt.Println(err)
				}
				if rows.Next() {
					rows.Scan(&userID)
				}
				rows.Close()
				userIDSlice = append(userIDSlice, userID)

				// Create competitors
				for colIndex, column := range currentRowSplit {
					if colIndex < len(columnNames) {
						if strings.Contains(columnNames[colIndex], "Predictions") {
							// fmt.Println(len(currentRowSplit))
							if competitorCache[column] == 0 {
								rows, err := u.Data.QueryContext(ctx, insertCompetitorQuery, column)

								if err != nil {
									fmt.Println(err)
								}
								if rows.Next() {
									rows.Scan(&competitorID)
								}
								rows.Close()
								competitorCache[column] = competitorID

							}
							// match exist, map the competitor to the match
							if matchIDCache[gamesCache[colIndex-3]] != 0 {
								rows, err := u.Data.QueryContext(ctx, insertParticipantQuery, matchIDCache[gamesCache[colIndex-3]], competitorCache[column])
								if err != nil {
									fmt.Println(err)
								}
								if rows.Next() {
									rows.Scan(&participantID)
								}
								rows.Close()

								//create a prediction
								rows, err = u.Data.QueryContext(ctx, insertPredictionQuery, userID, participantID)
								if err != nil {
									fmt.Println(err)
								}
								if rows.Next() {
									rows.Scan(&predictionID)
								}
								rows.Close()
							}
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

func getGames(row []string) ([]string, []string) {
	var games []string
	for i, col := range row {
		col = strings.Replace(col, "\"", "", -1)
		if strings.Contains(col, "Predictions") {
			col = strings.Replace(col, "Predictions [", "", -1)
			col = strings.Trim(col, " ")
			col = strings.Replace(col, "]", "", -1)
			games = append(games, col)
		}
		row[i] = col
	}
	return row, games
}
