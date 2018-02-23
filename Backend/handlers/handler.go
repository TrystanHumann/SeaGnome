package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/tealeg/xlsx"
	models "github.com/twitchguy/SeaGnome/Backend/Models"
)

//Ping ... Test
func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong!"))
}

//Excel ... Format excel
func Excel(w http.ResponseWriter, r *http.Request) {
	excelFileName := "../test1.csv"

	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println(err)
	}

	var users []models.User
	for i, sheet := range xlFile.Sheets {
		if i < 1 {
			for j := range sheet.Rows {
				var user models.User
				for k := range sheet.Cols {
					column := sheet.Cell(0, k).String()
					data := sheet.Cell(j, k).String()

					if column == "Timestamp" {
						user.Timestamp = data
					}
					if strings.Contains(column, "Twitch Username") {
						user.Twitch = data
					}
					if strings.Contains(column, "Twitter") {
						user.Twitter = data
					}
					if strings.Contains(column, "TIE BREAKER") {
						user.TieBreakerTime = data
					}
					if strings.Contains(column, "Predictions") {
						user.Predictions = append(user.Predictions, data)

					}
					if strings.Contains(column, "Bonus Question") {
						user.BonusQuestion = data
					}
				}
				users = append(users, user)
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
