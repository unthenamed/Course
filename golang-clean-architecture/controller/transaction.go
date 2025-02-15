package controller

import (
	"api-laundy/model"
	"api-laundy/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	tService service.TransactionService
	rg       *gin.RouterGroup
}

func (t *TransactionController) Route() {
	t.rg.POST("/transaction", t.createNewTransaction)
}

func (t *TransactionController) createNewTransaction(c *gin.Context) {
	var newModel model.Transaction

	err := c.ShouldBind(&newModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rModel, err := t.tService.CreateNewTransaction(newModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Created": rModel})
}

func ConstructorTransactionController(service service.TransactionService, rg *gin.RouterGroup) *TransactionController {
	return &TransactionController{tService: service, rg: rg}
}
