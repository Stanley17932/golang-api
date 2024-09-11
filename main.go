package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json: "id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json: "quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "1", Title: "The Great Gatsby", Author: "Scott Fitzgerald", Quantity: 3},
	{ID: "1", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 4},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.Run("localhost:8000")

}
