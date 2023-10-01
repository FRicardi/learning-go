package main

import (
	"database/sql"
	albumController "example/gin-vinyl-store/controllers"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func main() {
	connectToDatabase()

	router := gin.Default()
	router.GET("/albums", albumController.GetAlbums)
	router.GET("/albums/:id", albumController.GetAlbumByID)
	router.POST("/albums", albumController.PostAlbums)

	router.Run("localhost:8080")
}

func connectToDatabase() {
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
