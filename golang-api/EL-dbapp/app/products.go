package app

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type tProduct struct {
	Id    int
	Name  string
	Price int
	Unit  string
}

func (t *tProduct) create() error {
	statment := `
		INSERT INTO
		product (name, price, unit) 
		VALUES ($1,$2,$3) RETURNING id`
	err := DB.QueryRow(statment, t.Name, t.Price, t.Unit).Scan(&t.Id)
	return err
}

func (t *tProduct) update() error {
	statment := `UPDATE product SET name=$2, price=$3, unit=$4 WHERE id=$1`
	_, err := DB.Exec(statment, t.Id, t.Name, t.Price, t.Unit)
	return err
}

func (t *tProduct) read(filed string, value any) error {
	statment := `SELECT * FROM product WHERE ` + filed + "=$1"
	err := DB.QueryRow(statment, value).Scan(&t.Id, &t.Name, &t.Price, &t.Unit)
	return err
}

func readRows(s *[]tProduct, statment string) error {
	rows, err := DB.Query(statment)
	if err != nil {
		return err
	}

	for rows.Next() {
		var nProduct tProduct
		err = rows.Scan(&nProduct.Id, &nProduct.Name, &nProduct.Price, &nProduct.Unit)
		fmt.Println(nProduct)
		if err != nil {
			return err
		}
		*s = append(*s, nProduct)
	}
	err = rows.Close()
	return err
}

func (t *tProduct) delete() error {
	statment := `DELETE FROM product WHERE id = $1`
	result, err := DB.Exec(statment, t.Id)
	r, _ := result.RowsAffected()
	if r == 0 && err == nil {
		return errors.New("Not Found!")
	}
	return err

}

/*
   Create Product
*/

func CreateProduct(c *gin.Context) {
	var nProduct tProduct
	var err error
	if err = c.ShouldBind(&nProduct); err != nil {
		badRequestRespons("body", c)
		return
	}

	if err = nProduct.create(); err != nil {
		badGatewayRespons(err.Error(), c)
		return
	}

	createdRespons(nProduct, c)
}

/*
   List Product
*/

func ListProduct(c *gin.Context) {
	var sProduct []tProduct
	var err error

	query := c.Query("productName")
	statment := `SELECT * FROM product `
	if strings.TrimSpace(query) != "" {
		statment += `WHERE name ILIKE '%` + query + `%'`
	}

	err = readRows(&sProduct, statment)
	if err != nil {
		badGatewayRespons(err.Error(), c)
		return
	}

	if len(sProduct) > 0 {
		okRespons(sProduct, c)
	} else {
		if strings.TrimSpace(query) != "" {
			notFoundRespons("productName", c)
		} else {
			notFoundRespons("data empty, product", c)
		}
	}
}

/*
Product By Id
*/
func GetProductById(c *gin.Context) {
	var nProduct tProduct

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		badRequestRespons("id", c)
		return
	}

	if err = nProduct.read("id", id); err != nil {
		if strContains(err.Error(), "no rows") {
			notFoundRespons("id", c)
		} else {
			badGatewayRespons(err.Error(), c)
		}
		return
	}

	okRespons(nProduct, c)
}

/*
   Update Product
*/

func UpdateProduct(c *gin.Context) {
	var nProduct tProduct
	var uProduct tProduct

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		badRequestRespons("id", c)
		return
	}

	err = c.ShouldBind(&uProduct)
	if err != nil {
		badRequestRespons("body", c)
		return
	}

	if err = nProduct.read("id", id); err != nil {
		if strContains(err.Error(), "no rows") {
			notFoundRespons("id", c)
		} else {
			badGatewayRespons(err.Error(), c)
		}
		return
	}

	if strings.TrimSpace(uProduct.Name) != "" {
		nProduct.Name = uProduct.Name
	}
	if uProduct.Price != 0 {
		nProduct.Price = uProduct.Price
	}
	if strings.TrimSpace(uProduct.Unit) != "" {
		nProduct.Unit = uProduct.Unit
	}
	if err = nProduct.update(); err != nil {
		badGatewayRespons(err.Error(), c)
		return
	}

	okRespons(nProduct, c)
}

/*
   Delete Product
*/

func DeleteProduct(c *gin.Context) {
	var nProduct tProduct
	var err error
	nProduct.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		badRequestRespons("id", c)
		return
	}

	if err = nProduct.delete(); err != nil {
		if strContains(err.Error(), "Not Found") {
			notFoundRespons("id", c)
		} else {
			badGatewayRespons(err.Error(), c)
		}
		return
	}

	okRespons("deleted", c)
}
