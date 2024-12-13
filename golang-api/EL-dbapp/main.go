package main

import (
	"el-database/app"

	"github.com/gin-gonic/gin"
)

func main() {
	defer app.DB.Close()
	router := gin.Default()

	// Group Custommers
	customers := router.Group("/customers")
	{
		customers.POST("/", app.CreateCustomers)
		customers.GET("/:id", app.GetCustomersById)
		customers.PUT("/:id", app.UpdateCustomersById)
		customers.DELETE("/:id", app.DeleteCustomersById)
	}

	// Group employees
	employees := router.Group("/employees")
	{
		employees.POST("/", app.CreateEmployees)
		employees.GET("/:id", app.GetEmployeesById)
		employees.PUT("/:id", app.UpdateEmployees)
		employees.DELETE("/:id", app.DeleteEmployees)
	}

	products := router.Group("/products")
	{
		products.POST("/", app.CreateProduct)
		products.GET("/", app.ListProduct)
		products.GET("/:id", app.GetProductById)
		products.PUT("/:id", app.UpdateProduct)
		products.DELETE("/:id", app.DeleteProduct)
	}

	transaction := router.Group("/transactions")
	{
		transaction.POST("/", app.CreateTrx)
		transaction.GET("/:id_bill", app.GetTransactionById)
		transaction.GET("/", app.GetTransaction)
	}

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
