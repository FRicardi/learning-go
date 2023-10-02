package albumService

import (
	"database/sql"
	"errors"
	sqlite "example/gin-vinyl-store/database"
	types "example/gin-vinyl-store/types"
)

func GetAlbums() ([]types.Album, error) {
	rows, err := sqlite.GetDB().Query("SELECT * FROM albums")

	if err != nil {
		return nil, errors.New("Could not find albums")
	}

	var albums []types.Album
	defer rows.Close()

	for rows.Next() {
		var album types.Album

		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			return nil, errors.New("Albums not properly formatted")
		}

		albums = append(albums, album)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.New("Error reading albums")
	}

	return albums, nil
}

func AddAlbum(newAlbum types.Album) (int64, error) {
	result, err := sqlite.GetDB().Exec("INSERT INTO album (title, artist, price) values (?, ?, ?)", newAlbum.Title, newAlbum.Artist, newAlbum.Price)

	if err != nil {
		return 0, errors.New("Error creating album")
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, errors.New("Error fetching album inserted")
	}

	return id, nil
}

func GetAlbumByID(id string) (*types.Album, error) {
	row := sqlite.GetDB().QueryRow("SELECT * FROM albums WHERE id = ?", id)

	var album types.Album

	if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Album not found")
		}

		return nil, errors.New("Album not properly formatted")
	}

	return &album, nil
}
