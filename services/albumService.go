package albumService

import (
	"errors"
	connector "example/gin-vinyl-store/database"
	types "example/gin-vinyl-store/types"
	"fmt"
)

func GetAlbums() ([]types.Album, error) {
	var albums []types.Album

	result := connector.GetDB().Find(&albums)

	fmt.Println(albums)
	fmt.Println(result)

	if result.Error != nil {
		return nil, errors.New("Could not find albums")
	}

	return albums, nil
}

func AddAlbum(newAlbum types.Album) (int64, error) {
	result := connector.GetDB().Create(&newAlbum)

	if result.Error != nil {
		return 0, errors.New("Error creating album")
	}

	return result.RowsAffected, nil
}

func GetAlbumByID(id int64) (*types.Album, error) {
	var album types.Album
	result := connector.GetDB().First(&album, id)

	fmt.Println(album)
	fmt.Println(result)
	if result.Error != nil {
		return nil, errors.New("Album not found")
	}

	return &album, nil
}
