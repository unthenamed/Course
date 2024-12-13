package app

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BillDetails struct {
	Id           int `json:"id"`
	BillId       int `json:"billId"`
	ProductId    int `json:"productId"`
	ProductPrice int `json:"productPrice"`
	Qty          int `json:"qty"`
}

type tBills struct {
	Id          int           `json:"id"`
	BillDate    JsonBirthDate `json:"billDate"`
	EntryDate   JsonBirthDate `json:"entryDate"`
	FinishDate  JsonBirthDate `json:"finishDate"`
	EmployeeId  int           `json:"employeeId"`
	CustomerId  int           `json:"customerId"`
	BillDetails []BillDetails `json:"billDetails"`
}

type ResponseBillDetails struct {
	Id           int      `json:"id"`
	BillId       int      `json:"billId"`
	ProductId    tProduct `json:"product"`
	ProductPrice int      `json:"productPrice"`
	Qty          int      `json:"qty"`
}

type ResponseBills struct {
	Id          int                   `json:"id"`
	BillDate    JsonBirthDate         `json:"billDate"`
	EntryDate   JsonBirthDate         `json:"entryDate"`
	FinishDate  JsonBirthDate         `json:"finishDate"`
	EmployeeId  tEmployees            `json:"employee"`
	CustomerId  tCustomers            `json:"customer"`
	BillDetails []ResponseBillDetails `json:"billDetails"`
	TotalBill   int                   `json:"totalBill"`
}

type tQuery struct {
	startDate   JsonBirthDate
	endDate     JsonBirthDate
	productName string
}

type tTransaction struct {
	Bills     tBills
	nProduct  tProduct
	nEmployee tEmployees
	nCustomer tCustomers
}

// Method untuk mengambil data dari database
func (t *tBills) read() error {
	t.BillDetails = make([]BillDetails, 1)

	// transaction table
	statment := "SELECT * FROM transaction_detail WHERE bill_id = $1;"
	err := DB.QueryRow(statment, t.Id).Scan(
		&t.BillDetails[0].Id,
		&t.BillDetails[0].BillId,
		&t.BillDetails[0].ProductId,
		&t.BillDetails[0].ProductPrice,
		&t.BillDetails[0].Qty,
	)

	// transaction table
	statment = "SELECT * FROM transaction WHERE id = $1"
	err = DB.QueryRow(statment, t.Id).Scan(
		&t.Id,
		&t.BillDate,
		&t.EntryDate,
		&t.FinishDate,
		&t.EmployeeId,
		&t.CustomerId,
	)
	return err
}

// method untuk mempush instance ke database \
func (t *tBills) upload() error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}

	// bill transactions
	statment := `
		INSERT INTO transaction 
		(bill_date, entry_date, finish_date, employee_id, customer_id) 
		VALUES($1, $2, $3, $4, $5) RETURNING id ;`

	err = tx.QueryRow(
		statment,
		t.BillDate.Format("2006-01-01"),
		t.EntryDate.Format("2006-01-02"),
		t.FinishDate.Format("2006-01-02"),
		t.EmployeeId,
		t.CustomerId).Scan(&t.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	t.BillDetails[0].BillId = t.Id

	// get product price
	var price tProduct
	err = price.read("id", t.BillDetails[0].ProductId)
	if err != nil {
		tx.Rollback()
		return err
	}
	t.BillDetails[0].ProductPrice = price.Price

	/// detail transaction
	statment = `
		INSERT INTO transaction_detail 
		(bill_id, product_id, product_price, qty) 
		VALUES($1, $2, $3, $4) RETURNING id ;`

	err = tx.QueryRow(
		statment,
		t.BillDetails[0].BillId,
		t.BillDetails[0].ProductId,
		t.BillDetails[0].ProductPrice,
		t.BillDetails[0].Qty).Scan(&t.BillDetails[0].Id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil

}

/*
Create Transaction
*/
func CreateTrx(c *gin.Context) {
	var nBills tBills
	err := c.ShouldBind(&nBills)
	if err != nil {
		badRequestRespons(err.Error(), c)
		return
	}

	if err = nBills.upload(); err != nil {
		badGatewayRespons(err.Error(), c)
		return
	}

	c.JSON(http.StatusCreated, responsMap("Successfully Created!", nBills))
}

/*
   Get Transaction
*/

func GetTransactionById(c *gin.Context) {
	var nBills tBills
	var nProduct tProduct
	var nEmployee tEmployees
	var nCustomer tCustomers
	var err error

	// Mengambil path params
	nBills.Id, err = strconv.Atoi(c.Param("id_bill"))
	if err != nil {
		badRequestRespons(err.Error(), c)
		return
	}

	// Mengambil data bill
	if err = nBills.read(); err != nil {
		if strContains(err.Error(), "no rows") {
			notFoundRespons("id bill", c)
		} else {
			badGatewayRespons(err.Error(), c)
		}
		return
	}

	// Mengambil  data product
	if err = nProduct.read("id", nBills.BillDetails[0].ProductId); err != nil {
		if strContains(err.Error(), "no rows") {
			notFoundRespons("Id Product", c)
		} else {
			badGatewayRespons(err.Error(), c)
		}
		return
	}

	// Mengambil data employee
	nEmployee.Id = nBills.EmployeeId
	if err = nEmployee.read(); err != nil {
		if strContains(err.Error(), "no rows") {
			notFoundRespons("id employee", c)
		} else {
			badGatewayRespons(err.Error(), c)
		}
		return
	}

	// Mengambil data  customer
	nCustomer.Id = nBills.CustomerId
	if err = nCustomer.getById(); err != nil {
		if strContains(err.Error(), "no rows") {
			notFoundRespons("id customer", c)
		} else {
			badGatewayRespons(err.Error(), c)
		}
		return
	}

	//  Membuat respons dari data yang telah di ambilil
	detail := []ResponseBillDetails{{
		Id:           nBills.BillDetails[0].Id,
		BillId:       nBills.BillDetails[0].BillId,
		ProductId:    nProduct,
		ProductPrice: nBills.BillDetails[0].ProductPrice,
		Qty:          nBills.BillDetails[0].Qty,
	}}
	respon := ResponseBills{
		Id:          nBills.Id,
		BillDate:    nBills.BillDate,
		EntryDate:   nBills.EntryDate,
		FinishDate:  nBills.FinishDate,
		EmployeeId:  nEmployee,
		CustomerId:  nCustomer,
		BillDetails: detail,
		TotalBill:   detail[0].ProductPrice + detail[0].Qty,
	}

	// Memberikan respon dengan data  respon
	c.JSON(http.StatusOK, responsMap("Product Details!", respon))

}

/*
   List Transaction
*/

func GetTransaction(c *gin.Context) {
	fmt.Println(1)

	// Mengambil query param
	var timeZero JsonBirthDate
	var nQuery = tQuery{
		productName: c.Query("productName"),
	}
	nQuery.startDate.UnmarshalJSON([]byte(c.Query("startDate")))
	nQuery.endDate.UnmarshalJSON([]byte(c.Query("endDate")))

	// Membuat query
	var args []interface{}
	statment := `
		SELECT t.id, d.product_id, t.customer_id, t.employee_id
		FROM transaction_detail AS d
		JOIN transaction AS t ON d.bill_id = t.id
		JOIN product AS p ON d.product_id = p.id
		`
	query := ""
	counter := 1

	if nQuery.startDate != timeZero {
		if counter != 1 {
			query += " AND"
		}
		query += fmt.Sprintf(" t.entry_date >= $%d", counter)
		args = append(args, nQuery.startDate.Format("2006-01-02"))
		counter++
	}
	if nQuery.endDate != timeZero {
		if counter != 1 {
			query += " AND"
		}
		query += fmt.Sprintf(" t.finish_date <= $%d", counter)
		args = append(args, nQuery.endDate.Format("2006-01-02"))
		counter++
	}
	if nQuery.productName != "" {
		if counter != 1 {
			query += " AND"
		}
		query += fmt.Sprintf(" p.name LIKE $%d", counter)
		args = append(args, "%"+nQuery.productName+"%")
		counter++
	}
	if counter != 1 {
		statment += "WHERE" + query + " ;"
	}

	// Mengambil data dari database
	rows, err := DB.Query(statment, args...)
	if err != nil {
		badGatewayRespons(err.Error(), c)
		return
	}

	// Parsing rows
	sTransaction := []tTransaction{}
	for rows.Next() {
		newTx := tTransaction{}
		newTx.Bills.BillDetails = make([]BillDetails, 1)

		err = rows.Scan(
			&newTx.Bills.Id,
			&newTx.nProduct.Id,
			&newTx.nCustomer.Id,
			&newTx.nEmployee.Id,
		)

		if err != nil {
			badGatewayRespons(err.Error(), c)
			return
		}

		sTransaction = append(sTransaction, newTx)
	}
	defer rows.Close()

	// Memberikan respon
	var respons []ResponseBills
	for _, tx := range sTransaction {

		if tx.nEmployee.read() != nil {
			badGatewayRespons("fetch customer error!", c)
			return
		}
		if tx.nProduct.read("id", tx.nProduct.Id) != nil {
			badGatewayRespons("fetch customer error!", c)
			return
		}
		if tx.nCustomer.getById() != nil {
			badGatewayRespons("fetch customer error!", c)
			return
		}
		if tx.Bills.read() != nil {
			badGatewayRespons("fetch customer error!", c)
			return
		}

		detail := []ResponseBillDetails{{
			Id:           tx.Bills.BillDetails[0].Id,
			BillId:       tx.Bills.BillDetails[0].BillId,
			ProductId:    tx.nProduct,
			ProductPrice: tx.Bills.BillDetails[0].ProductPrice,
			Qty:          tx.Bills.BillDetails[0].Qty,
		}}
		respon := ResponseBills{
			Id:          tx.Bills.Id,
			BillDate:    tx.Bills.BillDate,
			EntryDate:   tx.Bills.EntryDate,
			FinishDate:  tx.Bills.FinishDate,
			EmployeeId:  tx.nEmployee,
			CustomerId:  tx.nCustomer,
			BillDetails: detail,
			TotalBill:   detail[0].ProductPrice + detail[0].Qty,
		}

		respons = append(respons, respon)
	}

	if len(sTransaction) == 0 {
		notFoundRespons("transactions", c)
		return
	}

	// Memberikan respon dengan data  respon
	c.JSON(http.StatusOK, responsMap("Product Details!", respons))
}
