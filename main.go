package main

import (
	albumController "example/gin-vinyl-store/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", albumController.GetAlbums)
	router.GET("/albums/:id", albumController.GetAlbumByID)
	router.POST("/albums", albumController.PostAlbums)

	router.Run("localhost:8080")
}
