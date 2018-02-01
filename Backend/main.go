package main

import (
	"fmt"
	"strings"

	"github.com/tealeg/xlsx"

	"github.com/trystanhumann/SeaGnome/Backend/models"
)

func main() {
	excelFileName := "../test.xlsx"

	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println(err)
	}

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
				err := stormDB.Save(&user)
				if err != nil {
					// fmt.Println(err)
				}

			}

		}
	}

	var users []models.User
	err = stormDB.All(&users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(users)
}
