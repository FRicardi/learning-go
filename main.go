package main

import (
	albumController "example/gin-vinyl-store/controllers"
	sqlite "example/gin-vinyl-store/database"

	"github.com/gin-gonic/gin"
)

func main() {
	sqlite.ConnectToDatabase()

	router := gin.Default()
	router.GET("/albums", albumController.GetAlbums)
	router.GET("/albums/:id", albumController.GetAlbumByID)
	router.POST("/albums", albumController.PostAlbums)

	router.Run("localhost:8080")
}
