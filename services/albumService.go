package albumService

import (
	"errors"
	types "example/gin-vinyl-store/types"
)

var albums = []types.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums() []types.Album {
	return albums
}

func AddAlbum(newAlbum types.Album) types.Album {
	albums = append(albums, newAlbum)
	return newAlbum
}

func GetAlbumByID(id string) (*types.Album, error) {

	for _, album := range albums {
		if album.ID == id {
			return &album, nil
		}
	}

	return &types.Album{}, errors.New("Album not found")
}
