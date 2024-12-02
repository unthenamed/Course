package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Customer struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

var Customers = []Customer{{Id: 1, Name: "Abdul", Phone: "082310607461", Address: "Kp.Nusa"}}

func main() {
	router := gin.Default()

	// get all customer
	router.GET("/customer", getAllCustomer)

	// create new customer
	router.POST("/customer", createCustomer)

	// run api
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func getAllCustomer(c *gin.Context) {
	c.JSON(http.StatusOK, Customers)
}

func createCustomer(c *gin.Context) {
	var newCustomer Customer
	err := c.ShouldBind(&newCustomer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Customers = append(Customers, newCustomer)
	c.JSON(http.StatusCreated, newCustomer)
}
