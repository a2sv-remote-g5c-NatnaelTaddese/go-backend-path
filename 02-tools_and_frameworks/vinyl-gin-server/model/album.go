package model

import (
	"fmt"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var Albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbumByID(id string) (*Album, error) {
	for _, album := range Albums {
		if album.ID == id {
			return &album, nil
		}
	}
	return nil, fmt.Errorf("album not found")
}

func CreateAlbum(album Album) (string, error) {
	id := fmt.Sprintf("%d", len(Albums)+1)
	album.ID = id
	Albums = append(Albums, album)
	return id, nil
}
