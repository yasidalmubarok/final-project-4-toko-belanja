package transaction_history_pg

import (
	"database/sql"
	"toko-belanja-app/dto"
	"toko-belanja-app/entity"
	"toko-belanja-app/pkg/errs"
	"toko-belanja-app/repository/transaction_history_repository"
)

type transactionHistoryPg struct {
	db *sql.DB
}

func NewTransactionHistoryPg(db *sql.DB) transaction_history_repository.TransactionHistoryRepository {
	return &transactionHistoryPg{db: db}
}

const (
	createTransaction = `
		INSERT INTO transaction_histories (user_id, product_id, quantity, total_price)
		SELECT
			$1,
			$2,
			$3,
			(p.price * $3) AS total_price
		FROM products p
		WHERE p.id = $1
		RETURNING
			total_price,
			quantity,
			p.title AS product_title;
	`
	getTransaction = `
		SELECT
			t.id,
			t.product_id,
			t.user_id,
			t.quantity,
			t.total_price,
			p.id AS product_id,
			p.title AS product_title,
			p.price,
			p.stock,
			p.category_id,
			p.created_at AS product_created_at,
			p.updated_at AS product_updated_at,
			u.id AS user_id,
			u.email,
			u.full_name,
			u.balance,
			u.created_at AS user_created_at,
			u.updated_at AS user_updated_at
		FROM
			transaction_histories AS t
		LEFT JOIN
			products AS p
		ON
			t.product_id = p.id
		LEFT JOIN
			users AS u
		ON
			t.user_id = u.id
		WHERE
			t.deleted_at IS NULL
		ORDER BY
			t.id ASC;
	`

	getMyTransaction = `
		SELECT t.id, 
			t.product_id,
			t.user_id, 
			t.quantity, 
			t.total_price, 
			p.id, 
			p.title, 
			p.price, 
			p.stock, 
			p.category_Id, 
			p.created_at, 
			p.updated_at
		FROM 
			transaction_histories as t
		LEFT JOIN
			products as p
		ON
			t.product_id = p.id
		WHERE t.user_id = $1 
		t.deleted_at IS NULL
		ORDER BY 
		t.id ASC
		`
)

func (t *transactionHistoryPg) CreateNewTransaction(transactionPayLoad *entity.TransactionHistory) (*dto.TransactionBill, errs.Error) {
	tx, err := t.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	var transaction dto.TransactionBill

	row := tx.QueryRow(
		createTransaction,
		transactionPayLoad.UserId,
		transactionPayLoad.ProductId,
		transactionPayLoad.Quantity,
	)
	err = row.Scan(
		&transaction.TotalPrice,
		&transaction.Quantity,
		&transaction.ProductTitle,
	)

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong" + err.Error())
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	return &transaction, nil
}

func (t *transactionHistoryPg) getTransaction() ([]transaction_history_repository.TransactionProductMapped, errs.Error) {
	transactionProducts := []transaction_history_repository.TransactionProduct{}
	rows, err := t.db.Query(getTransaction)

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong" + err.Error())
	}

	for rows.Next() {
		var transactionProduct transaction_history_repository.TransactionProduct

		err := rows.Scan(
			&transactionProduct.TransactionHistory.Id,
			&transactionProduct.TransactionHistory.ProductId,
			&transactionProduct.TransactionHistory.UserId,
			&transactionProduct.TransactionHistory.Quantity,
			&transactionProduct.TransactionHistory.TotalPrice,
			&transactionProduct.Product.Id,
			&transactionProduct.Product.Title,
			&transactionProduct.Product.Price,
			&transactionProduct.Product.Stock,
			&transactionProduct.Product.CategoryId,
			&transactionProduct.Product.CreatedAt,
			&transactionProduct.Product.UpdatedAt,
			&transactionProduct.User.Id,
			&transactionProduct.User.Email,
			&transactionProduct.User.FullName,
			&transactionProduct.User.Balance,
			&transactionProduct.User.CreatedAt,
			&transactionProduct.User.UpdatedAt,
		)

		if err != nil {
			return nil, errs.NewInternalServerError("something went wrong " + err.Error())
		}

		transactionProducts = append(transactionProducts, transactionProduct)
	}

	result := transaction_history_repository.TransactionProductMapped{}
	return result.HandleMappingTransactionWithProduct(transactionProducts), nil
}
