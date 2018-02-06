package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/tealeg/xlsx"
	"github.com/trystanhumann/SeaGnome/Backend/types"
)

// Uploads : Handles requests involved with uploads
type Uploads struct {
	Data *sqlx.DB
}

// ServeHTTP : Listens for a request and creates a response
func (h *Uploads) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	excelFileName := "../test1.csv"

	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println(err)
	}

	var users []types.User
	for i, sheet := range xlFile.Sheets {
		if i < 1 {
			for j := range sheet.Rows {
				var user types.User
				for k := range sheet.Cols {
					column := sheet.Cell(0, k).String()
					value := sheet.Cell(j, k).String()

					if column == "Timestamp" {
						user.Timestamp = value
					}
					if strings.Contains(column, "Twitch Username") {
						user.Twitch = value
					}
					if strings.Contains(column, "Twitter") {
						user.Twitter = value
					}
					if strings.Contains(column, "TIE BREAKER") {
						user.TieBreakerTime = value
					}
					if strings.Contains(column, "Predictions") {
						prediction := types.Prediction{
							Column:     column,
							Prediction: value,
						}
						user.Predictions = append(user.Predictions, prediction)
					}
					if strings.Contains(column, "Bonus Question") {
						user.BonusQuestion = value
					}
				}
				//users = append(users, user)
				// go data.
			}
		}
	}
	// fmt.Println(users[1])
	output, err := json.Marshal(users[1])
	if err != nil {
		fmt.Println(err)
	}
	w.Write(output)
}
