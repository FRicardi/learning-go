package albumService

import (
	types "example/gin-vinyl-store/packages"
	"net/http"

	"github.com/gin-gonic/gin"
)

var albums = []types.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums(ginContext *gin.Context) {
	ginContext.IndentedJSON(http.StatusOK, albums)
}

func AddAlbums(ginContext *gin.Context) {
	var newAlbum types.Album

	// makes use of executing something before IF method
	if err := ginContext.BindJSON(&newAlbum); err != nil {
		// kills the method
		return
	}

	albums = append(albums, newAlbum)
	ginContext.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumByID(ginContext *gin.Context) {
	id := ginContext.Param("id")

	for _, album := range albums {
		if album.ID == id {
			ginContext.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	ginContext.IndentedJSON(http.StatusNotFound, gin.H{"message": "album nottfound"})
}
