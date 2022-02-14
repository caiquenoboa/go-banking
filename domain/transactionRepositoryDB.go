package domain

import (
	"strconv"

	"github.com/caiquenoboa/go-banking/errs"
	"github.com/caiquenoboa/go-banking/logger"
	"github.com/jmoiron/sqlx"
)

type TransactionRepositoryDB struct {
	client *sqlx.DB
}

func (d TransactionRepositoryDB) NewTransaction(t Transaction) (*Transaction, *errs.AppError) {
	sqlInsert := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values (?, ?, ?, ?)"

	result, err := d.client.Exec(sqlInsert, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	if err != nil {
		logger.Error("Error while creating the transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()

	if err != nil {
		logger.Error("Error while getting last insert id from new transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	t.TransactionId = strconv.FormatInt(id, 10)

	return &t, nil
}

func NewTransactionRepositoryDB(client *sqlx.DB) TransactionRepositoryDB {
	return TransactionRepositoryDB{client}
}
