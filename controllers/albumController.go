package albumController

import (
	types "example/gin-vinyl-store/packages"
	albumService "example/gin-vinyl-store/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAlbums(ctx *gin.Context) {
	albums := albumService.GetAlbums()

	ctx.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumByID(ctx *gin.Context) {
	id := ctx.Param("id")

	if album, error := albumService.GetAlbumByID(id); error == nil {
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

	createdAlbum := albumService.AddAlbum(newAlbum)
	ctx.IndentedJSON(http.StatusCreated, createdAlbum)
}
