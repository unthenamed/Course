package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id     int    `json:"id"`
	Judul  string `json:"judul"`
	Terbit int    `json:"terbit"`
}

var books = []Book{{Id: 1, Judul: "mentari pajar", Terbit: 2005}, {Id: 2, Judul: "sinar alam", Terbit: 2015}, {Id: 3, Judul: "kaidah islam", Terbit: 2017}}

func main() {
	router := gin.Default()

	router.GET("/books", getBooks)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}
