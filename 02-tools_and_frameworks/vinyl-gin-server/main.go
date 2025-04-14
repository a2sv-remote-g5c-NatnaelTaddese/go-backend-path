package main

import (
	"net/http"
	"vinyl-server/model"

	"github.com/gin-gonic/gin"
)

func HandleGetAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, model.Albums)
}

func HandleGetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	album, err := model.GetAlbumByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, album)
}

func main() {
	router := gin.Default()
	router.GET("/albums", HandleGetAlbums)
	router.GET("/albums/:id", HandleGetAlbumByID)

	router.Run(":8080")
}
