package domain

import (
	"database/sql"
	"strconv"

	"github.com/caiquenoboa/go-banking/errs"
	"github.com/caiquenoboa/go-banking/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDB struct {
	client *sqlx.DB
}

func (d AccountRepositoryDB) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"

	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)

	if err != nil {
		logger.Error("Error while creating the account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()

	if err != nil {
		logger.Error("Error while getting last insert id from new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	a.AccountId = strconv.FormatInt(id, 10)

	return &a, nil
}

func (d AccountRepositoryDB) GetAmountById(id string) (*Account, *errs.AppError) {
	FindById := "select amount from accounts where account_id = ?"

	var a Account
	err := d.client.Get(&a, FindById, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Account not found")
		} else {
			logger.Error("Error while scanning account " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &a, nil
}

func (d AccountRepositoryDB) UpdateBalanceAmountById(accountId string, amount float64) *errs.AppError {
	UpdateQuery := "UPDATE accounts SET amount = :amount WHERE account_id = :account"

	_, err := d.client.NamedExec(UpdateQuery,
		map[string]interface{}{
			"amount":  amount,
			"account": accountId,
		})

	if err != nil {
		if err == sql.ErrNoRows {
			return errs.NewNotFoundError("Account not found")
		} else {
			logger.Error("Error while updating account amount " + err.Error())
			return errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return nil
}

func NewAccountRepositoryDb(client *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{client}
}
