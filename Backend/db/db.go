package db

import (
	"fmt"

	"github.com/asdine/storm"
	models "github.com/twitchguy/SeaGnome/Backend/Models"
)

// InitDB ... Initializes storm database and returns a pointer to the db obj
func InitDB(file string) (*storm.DB, error) {
	db, err := storm.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	return db, err
}

// InitBuckets ... Initializes the bucket objs
func InitBuckets(db *storm.DB) error {
	err := db.Init(&models.User{})
	if err != nil {
		fmt.Println(err)
	}

	err = db.Init(&models.Results{})
	if err != nil {
		fmt.Println(err)
	}
	return err
}
