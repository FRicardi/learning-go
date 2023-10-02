package albumController

import (
	albumService "example/gin-vinyl-store/services"
	types "example/gin-vinyl-store/types"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAlbums(ctx *gin.Context) {
	if albums, error := albumService.GetAlbums(); error == nil {
		ctx.IndentedJSON(http.StatusOK, albums)
	} else {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": error})
	}
}

func GetAlbumByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	if album, error := albumService.GetAlbumByID(int64(id)); error == nil {
		ctx.IndentedJSON(http.StatusOK, *album)
	} else {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": error})
	}
}

func PostAlbums(ctx *gin.Context) {
	var newAlbum types.Album

	if err := ctx.BindJSON(&newAlbum); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Could not create Album"})
		return
	}

	if rowsAffected, error := albumService.AddAlbum(newAlbum); error == nil {
		ctx.IndentedJSON(http.StatusCreated, rowsAffected)
	} else {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": error})
	}
}
