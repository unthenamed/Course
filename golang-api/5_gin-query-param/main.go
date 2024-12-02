package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

var UserS = []User{
	{1, "Abdul", "BlokM"},
	{2, "Jalil", "BlokD"},
	{3, "Khoironi", "BlokC"},
	{4, "Usman", "BlokC"},
}

func main() {
	router := gin.Default()

	//Get all user
	router.GET("/user", getUser)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}

func getUser(c *gin.Context) {
	address := c.Query("address")
	if address == "" {
		c.JSON(http.StatusOK, UserS)
		return
	}

	matched := []User{}
	for _, user := range UserS {
		if strings.Contains(strings.ToLower(user.Address), strings.ToLower(address)) {
			matched = append(matched, user)
		}
	}

	if len(matched) > 0 {
		c.JSON(http.StatusOK, matched)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "address not found!"})
	}

}
