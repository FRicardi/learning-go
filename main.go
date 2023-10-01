package main

import (
	albumService "example/gin-vinyl-store/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", albumService.GetAlbums)
	router.GET("/albums/:id", albumService.GetAlbumByID)
	router.POST("/albums", albumService.AddAlbums)

	router.Run("localhost:8080")
}
