package repo

import (
	"api-laundy/model"
	"database/sql"
)

type transactionRepo struct {
	db *sql.DB
}

// Method untuk membuat transaksi baru dari sebuah model transaksi
func (c *transactionRepo) CreateNewTransaction(transaction model.Transaction) (model.Transaction, error) {

	// Membuat koneksi transaksi
	tx, err := c.db.Begin()
	if err != nil {
		return model.Transaction{}, err
	}

	// Membuat transaksi
	err = tx.QueryRow(
		`
		INSERT INTO transaction 
		(bill_date, entry_date, finish_date, employee_id, customer_id) 
		VALUES($1, $2, $3, $4, $5) RETURNING id ;
		`,
		transaction.Bill.BillDate,
		transaction.Bill.EntryDate,
		transaction.Bill.FinishDate,
		transaction.Bill.EmployeeId,
		transaction.Bill.BillDate,
	).Scan(&transaction.Bill.Id)

	if err != nil {
		tx.Rollback()
		return model.Transaction{}, err
	}

	// Membuat Detail transaksi
	for index, _ := range transaction.BillDetails {
		// Mengambil bill id
		transaction.BillDetails[index].BillId = transaction.Bill.Id

		// Mengambil  harga produk
		err = tx.QueryRow(
			"SELECT price FROM product WHERE id = $1",
			transaction.BillDetails[index].ProductId,
		).Scan(&transaction.BillDetails[index].ProductPrice)

		if err != nil {
			tx.Rollback()
			return model.Transaction{}, err
		}

		// Membuat Detail transaksi baru
		err = tx.QueryRow(
			`
			INSERT INTO transaction_detail 
			(bill_id, product_id, product_price, qty) 
			VALUES($1, $2, $3, $4) RETURNING id ;
			`,
			transaction.BillDetails[index].BillId,
			transaction.BillDetails[index].ProductId,
			transaction.BillDetails[index].ProductPrice,
			transaction.BillDetails[index].Qty,
		).Scan(&transaction.BillDetails[index].Id)

		if err != nil {
			tx.Rollback()
			return model.Transaction{}, err
		}

	}

	// commit ke database
	if err = tx.Commit(); err != nil {
		return model.Transaction{}, err
	}

	return transaction, nil
}

// interface yang menampung semua method transaksi
type TransactionRepo interface {
	CreateNewTransaction(model.Transaction) (model.Transaction, error)
}

// fungsi untuk pengeksekusian transaksi
func NewObjTransaction(db *sql.DB) TransactionRepo {
	return &transactionRepo{db: db}
}
