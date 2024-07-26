package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.DELETE("/albums/:id", removeAlbumByID)
	router.PATCH("/albums/:id", updateAlbumByID)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func removeAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for index, album := range albums {
		if album.ID == id {
			albums = append(albums[:index], albums[index+1:]...)
			c.IndentedJSON(http.StatusOK, albums)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album to delete not found"})
}

func updateAlbumByID(c *gin.Context) {
	id := c.Param("id")

	var albumToUpdate album

	for _, album := range albums {
		if album.ID == id {
			albumToUpdate = album
		}
	}

	req, err := c.Request.GetBody()

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "some error occured"})
	}
}
