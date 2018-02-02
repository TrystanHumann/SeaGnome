package data

import (
	"fmt"
	"os"

	"github.com/gocraft/dbr"
	"github.com/trystanhumann/SeaGnome/Backend/util"

	// driver for postgres connections
	_ "github.com/lib/pq"
)

var conn *dbr.Connection
var event dbr.EventReceiver

// ConnectToDB : Initialise the connection to the database
func init() {
	var err error
	conn, err = dbr.Open("postgres", util.GenerateConnectionString(), event)
	if err != nil {
		fmt.Println("Failed to open connection with postgres: " + err.Error())
		os.Exit(1)
	}
	if err := conn.Ping(); err != nil {
		fmt.Println("Failed open ping postgres: " + err.Error())
		os.Exit(1)
	} else {
		fmt.Println("Database connection established.")
	}
}

// Test : Temp function to force the init function to be called.
func Test() {
	fmt.Println("Success!")
}
