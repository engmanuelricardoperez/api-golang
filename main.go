package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

//deleteAlbumsByID, deletes album especified by path, including id number
func deleteAlbumsByID(c *gin.Context) {
	id := c.Param("id")
	position := 0
	for _, a := range albums {

		if a.ID == id {
			fmt.Println("entry in Id condition: ID detected=", position)
			albums = append(albums[:position], albums[position+1:]...)

			position = 0
			return
		}
		position++
		fmt.Println("all counts of position", position)
	}

}

//deleteAlbums , delete all albums in slice albums
func deleteAlbums(c *gin.Context) {
	albums = albums[:0]
}

func main() {
	// a := []string{"A", "B", "C", "D", "E", "F"}
	// var position, i int = 0, 2
	// fmt.Println("slice albums:", albums)
	// fmt.Println(len(albums))
	// fmt.Println("slice a:", a)
	// fmt.Println(len(a))

	// //a = append(a[position:], a[position+1:]...)
	// a[i] = a[len(a)-1] // Copy last element to index i.
	// a[len(a)-1] = ""   // Erase last element (write zero value).
	// a = a[:len(a)-1]   // Truncate slice.

	// //albums[position] = albums[len(albums)-1]

	// albums = append(albums[:position], albums[position+1:]...)

	// // albums[i] = albums[len(albums)-1] // Copy last element to index i.
	// // albums[len(albums)-1] = :0 // Erase last element (write zero value).
	// // albums = albums[:len(albums)-1]   // Truncate slice.

	// fmt.Println("slice albums:", albums)
	// fmt.Println(len(albums))
	// fmt.Println("slice a:", a)
	// fmt.Println(len(a))

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.DELETE("/albums", deleteAlbums)
	router.DELETE("/albums/:id", deleteAlbumsByID)
	router.Run("localhost:8080")
}
