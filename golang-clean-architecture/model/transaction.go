package model

type BillDetails struct {
	Id           int `json:"id"`
	BillId       int `json:"billId"`
	ProductId    int `json:"productId"`
	ProductPrice int `json:"productPrice"`
	Qty          int `json:"qty"`
}

type Bill struct {
	Id         int    `json:"id"`
	BillDate   string `json:"billDate"`
	EntryDate  string `json:"entryDate"`
	FinishDate string `json:"finishDate"`
	EmployeeId int    `json:"employeeId"`
	CustomerId int    `json:"customerId"`
}

type Transaction struct {
	Bill
	BillDetails []BillDetails
}
