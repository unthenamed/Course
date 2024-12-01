package library

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

func validateTransaction(err error, tx *sql.Tx) (result bool) {
	if err != nil {
		tx.Rollback()
		result = true
		fmt.Println(err)
	}
	return
}

func CreateOrder(db *sql.DB) {
	var tx *sql.Tx
	db, err := ConnectDb()
	if err != nil {
		return
	}
	defer db.Close()

	tx, err = db.Begin()
	if err != nil {
		fmt.Println("Creating new transaction error !")
		return
	}

	iOrder, iDetail, status := OrderInstance(tx)
	if status {
		return
	}
	iOrder.insertOrder(tx)
	iDetail.insertDetail(tx)

	if err = tx.Commit(); err != nil {
		fmt.Println("Transaction Rollback!")
	} else {
		fmt.Println("Transaction Commited")
	}

}

func CompleteOrder() {

	var tx *sql.Tx
	db, err := ConnectDb()
	if err != nil {
		return
	}
	defer db.Close()

	tx, err = db.Begin()
	if err != nil {
		return
	}

	sqlStatment := `SELECT order_id FROM "order" WHERE completion_date = $1`
	var date time.Time
	var rows *sql.Rows
	var orders = []Order{}
	date, err = time.Parse("2006-01-02", InputText("Completion Date YYYY-MM-DD :"))
	if err != nil {
		fmt.Println("Format salah! YYYY-MM-DF")
		return
	}
	rows, err = tx.Query(sqlStatment, date)

	for rows.Next() {
		order := Order{}
		err := rows.Scan(&order.order_id)
		if err != nil {
			return
		}
		orders = append(orders, order)
	}

	if len(orders) == 0 {
		fmt.Println("Order Not Fund!")
		return
	}

	for index, value := range orders {
		fmt.Println(index+1, value.order_id, value.order_date, value.customer_id, value.received_by)
		value.deleteDetail(tx)
		value.deleteOrder(tx)
	}
	err = tx.Commit()
	if !validateTransaction(err, tx) {
		fmt.Println("Order Terhapus!")
	}

}

func ViewOfListOrder() {
	db, err := ConnectDb()
	if err != nil {
		return
	}
	defer db.Close()

	sqlStatment := `
SELECT a.order_id, a.customer_id, b.order_detail_id, a.order_date, a.completion_date, b.service_id, b.qty
FROM "order" AS a
JOIN order_detail AS b
ON b.order_id = b.order_id 
ORDER BY a.order_id ASC;
`

	var rows *sql.Rows
	rows, err = db.Query(sqlStatment)
	if err != nil {
		return
	}

	tbSlice := convertRowsToTable(rows)
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"order_id", "customer_id", "order_detail_id", "order_date", "completion_date", "service_id", "qty"})
	t.AppendRows(tbSlice)
	t.AppendSeparator()
	//t.AppendFooter(table.Row{})
	t.Render()

}

func ViewOfListOrderById(id int) {
	db, err := ConnectDb()
	if err != nil {
		return
	}
	defer db.Close()

	sqlStatment := `
SELECT a.order_id, a.customer_id, b.order_detail_id, a.order_date, a.completion_date, b.service_id, b.qty
FROM "order" AS a
JOIN order_detail AS b
ON b.order_id = b.order_id 
WHERE a.order_id = $1;
`

	var rows *sql.Rows
	rows, err = db.Query(sqlStatment, id)
	if err != nil {
		return
	}

	tbSlice := convertRowsToTable(rows)
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"order_id", "customer_id", "order_detail_id", "order_date", "completion_date", "service_id", "qty"})
	t.AppendRows(tbSlice)
	t.AppendSeparator()
	//t.AppendFooter(table.Row{})
	t.Render()

}

func convertRowsToTable(rows *sql.Rows) (result []table.Row) {

	for rows.Next() {
		a := Order{}
		b := Order_detail{}
		err := rows.Scan(&a.order_id, &a.customer_id, &b.order_detail_id, &a.order_date, &a.completion_date, &b.service_id, &b.qty)
		if err != nil {
			return
		}

		newTable := table.Row{a.order_id, a.customer_id, b.order_detail_id, a.order_date.Format("2006-01-02"), a.completion_date.Format("2006-01-02"), b.service_id, b.qty}
		result = append(result, newTable)
	}
	defer rows.Close()
	return
}
