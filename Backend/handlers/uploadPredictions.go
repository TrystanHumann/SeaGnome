package handlers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
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
	parentContext := context.TODO()
	ctx, ctxCancel := context.WithTimeout(parentContext, time.Hour*2)
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

		// ID of the event that we will be updating data for
		eventID := r.FormValue("eventID")

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

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
		_, gamesCache := getGames(columnNames)
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
			var id int
			var matchID int
			err := u.Data.GetContext(ctx, &id, insertGameQuery, game)
			if err != nil {
				fmt.Println(err)
			}
			gameIDSlice = append(gameIDSlice, id)

			//upload a match
			err = u.Data.GetContext(ctx, &matchID, insertMatchQuery, eventID, id)
			if err != nil {
				fmt.Println(err)
			}
			if matchIDCache[game] == 0 {
				matchIDCache[game] = matchID
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
				err := u.Data.GetContext(ctx, &userID, insertUserQuery, strings.ToLower(currentRowSplit[1]), strings.ToLower(currentRowSplit[2]))
				if err != nil {
					fmt.Println(err)
				}
				userIDSlice = append(userIDSlice, userID)

				// Create competitors
				for colIndex, column := range currentRowSplit {
					if colIndex < len(columnNames) {
						if strings.Contains(columnNames[colIndex], "Predictions") {
							if competitorCache[column] == 0 {
								err := u.Data.GetContext(ctx, &competitorID, insertCompetitorQuery, column)
								if err != nil {
									fmt.Println(err)
								}
								competitorCache[column] = competitorID

							}
							// match exist, map the competitor to the match
							if matchIDCache[gamesCache[colIndex-3]] != 0 {
								err := u.Data.GetContext(ctx, &participantID, insertParticipantQuery, matchIDCache[gamesCache[colIndex-3]], competitorCache[column])
								if err != nil {
									fmt.Println(err)
								}

								//create a prediction
								err = u.Data.GetContext(ctx, &predictionID, insertPredictionQuery, userID, participantID)
								if err != nil {
									fmt.Println(err)
								}
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
			col = strings.Replace(col, "]", "", -1)
			col = strings.TrimSpace(col)
			games = append(games, col)
		}
		row[i] = col
	}
	return row, games
}
