package connector

import (
	"example/gin-vinyl-store/types"
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectToDatabase() {
	var err error
	db, err = gorm.Open(sqlite.Open("database.sqlite"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")

	db.AutoMigrate(&types.Album{})

	fmt.Println("Migrated successfully!")
}

func GetDB() *gorm.DB {
	if db == nil {
		ConnectToDatabase()
	}

	return db
}
