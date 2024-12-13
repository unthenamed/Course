package app

import (
	"el-database/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Deklarasi Database sebagai global variable
var DB = config.ConnectDB()


// string validasi
func stringNull(obj string) (state bool) {
	if strings.TrimSpace(obj) == "" {
		state = true
	}
	return
}

func strContains(obj, matching string) (state bool) {
	if strings.Contains(strings.ToLower(obj), strings.ToLower(matching)) {
		state = true
	}
	return
}


// syarat format respon
type Respons struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func responsMap(message string, t any) any {
	return Respons{Message: message, Data: t}
}


// shortcut json function
func okRespons(data any, c *gin.Context) {
	c.JSON(http.StatusOK, responsMap("OK", data))
}

func createdRespons(data any, c *gin.Context) {
	c.JSON(http.StatusCreated, responsMap("Created", data))
}

func notFoundRespons(errorFrom string, c *gin.Context) {
	c.JSON(http.StatusNotFound, responsMap("error", errorFrom+" not found!"))
}

func badGatewayRespons(data any, c *gin.Context) {
	c.JSON(http.StatusBadGateway, responsMap("error", data))
}

func badRequestRespons(errorFrom string, c *gin.Context) {
	c.JSON(http.StatusBadRequest, responsMap("error", "invalid "+errorFrom+"!"))
}
