package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Order struct {
	OrderId    int
	ReceivedBy string
}

var OrderS = []Order{
	{1, "Abdul"},
	{2, "Jalil"},
	{3, "Khoironi"},
}

func main() {
	router := gin.Default()

	router.GET("/order", getAllOrder)
	router.GET("/order/:id", getOrderByParam)

	router.POST("/order", createOrder)

	router.PUT("/order/:id", updateOrder)

	router.DELETE("/order/:id", deleteOrder)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}

}

func getAllOrder(c *gin.Context) {
	c.JSON(http.StatusOK, OrderS)
}

func getOrderByParam(c *gin.Context) {
	var id Order
	var err error
	id.OrderId, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id!"})
		return
	}

	for _, value := range OrderS {
		if value.OrderId == id.OrderId {
			c.JSON(http.StatusFound, value)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "id not found!"})
}

func createOrder(c *gin.Context) {
	var newOrder Order
	if err := c.ShouldBind(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	for _, value := range OrderS {
		if value.OrderId == newOrder.OrderId {
			c.JSON(http.StatusConflict, gin.H{"error": "id already exist!"})
			return
		}
	}

	OrderS = append(OrderS, newOrder)
	c.JSON(http.StatusOK, newOrder)
}

func updateOrder(c *gin.Context) {
	var newUpdateOrder Order
	orderId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id!"})
		return
	}

	if err = c.ShouldBind(&newUpdateOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "request bad!"})
		return
	}

	for i, v := range OrderS {
		if v.OrderId == orderId {
			if strings.TrimSpace(newUpdateOrder.ReceivedBy) != "" {
				OrderS[i].ReceivedBy = newUpdateOrder.ReceivedBy
				c.JSON(http.StatusOK, OrderS[i])
				return
			}
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "id not found!"})

}

func deleteOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id!"})
		return
	}

	newResult := []Order{}
	var index int
	for i, v := range OrderS {
		if v.OrderId != id {
			newResult = append(newResult, v)
		} else {
			index = i
		}
	}

	if len(OrderS) == len(newResult) {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	} else {
		c.JSON(http.StatusOK, OrderS[index])
		OrderS = newResult
	}
}
