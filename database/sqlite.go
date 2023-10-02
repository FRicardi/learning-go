package sqlite

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func ConnectToDatabase() {
	var err error
	db, err = sql.Open("sqlite3", "database.sqlite")

	if err != nil {
		log.Fatal(err)
	}

	if pingErr := db.Ping(); pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")
}

func GetDB() *sql.DB {
	if db == nil {
		ConnectToDatabase()
	}

	return db
}
