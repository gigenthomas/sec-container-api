package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type pageMeta struct {
	CurrentPage     uint64 `json:"currentPage"`
	Limit           uint64 `json:"limit"`
	TotalItems      uint64 `json:"totalItems"`
	TotalPages      uint64 `json:"totalPages"`
	HasPreviousPage bool   `json:"hasPreviousPage"`
	HasNextPage     bool   `json:"hasNextPage"`
}

type page[T any] struct {
	Items []T      `json:"items"`
	Meta  pageMeta `json:"meta"`
}

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type user struct {
	ID          string `json:"id"`
	Avatar      string `json:"avatar"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Jobtitle    string `json:"jobtitle"`
	Username    string `json:"username"`
	Location    string `json:"location"`
	Role        string `json:"role"`
	Posts       string `json:"posts"`
	CoverImg    string `json:"coverImg"`
	Followers   string `json:"followers"`
	Description string `json:"description"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

var users = []user{
	{ID: "1",
		Avatar:      "/static/images/avatars/1.jpg",
		Email:       "Monte.Auer31@yahoo.com",
		Name:        "Rafael Kunde",
		Jobtitle:    "Product Infrastructure Associate",
		Username:    "Delphia22",
		Location:    "Gislasonchester",
		Role:        "admin",
		Posts:       "",
		CoverImg:    "/static/images/placeholders/covers/1.jpg",
		Followers:   "667",
		Description: "Vestibulum rutrum rutrum neque. Aenean auctor gravida sem quam pede lobortis ligula, sit amet eleifend",
	},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getUsersByPage(c *gin.Context) {
	var currentMeta pageMeta
	currentMeta.CurrentPage = 1
	currentMeta.TotalItems = 1
	currentMeta.HasNextPage = false
	currentMeta.HasPreviousPage = false
	currentMeta.Limit = 10

	var userDoamin page[user]
	userDoamin.Items = users
	userDoamin.Meta = currentMeta

	c.IndentedJSON(http.StatusOK, userDoamin)
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

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

func getUserByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range users {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.GET("/users", getUsersByPage)
	router.GET("/users/:id", getUserByID)
	router.Run("localhost:8080")
}
