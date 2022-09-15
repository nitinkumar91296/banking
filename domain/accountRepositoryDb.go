package domain

import (
	"database/sql"
	"strconv"

	"github.com/jmoiron/sqlx"
	errs "github.com/nitinkumar91296/banking/errors"
	"github.com/nitinkumar91296/banking/logger"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{client: dbClient}
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?,?,?,?,?)"

	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last inserted ID for new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	a.AccountId = strconv.FormatInt(id, 10)

	return &a, nil
}

func (d AccountRepositoryDb) FindBy(accountId string) (*Account, *errs.AppError) {
	findAccountSql := "select account_id, customer_id, opening_date, account_type, amount, status from accounts where account_id = ?"

	var a Account
	err := d.client.Get(&a, findAccountSql, accountId)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("account not found" + err.Error())
			return nil, errs.NewNotFoundError("account Not Found")
		} else {
			logger.Error("Error while scanning account by ID" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected db error")
		}
	}

	return &a, nil
}

func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	// start db txn block
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting the new transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected Database error")
	}

	//inserting account balance
	result, err := tx.Exec(`INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values(?,?,?,?)`, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	// updating the balance
	if t.IsWithdrawl() {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount - ? where account_id = ?`, t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + ? where account_id = ?`, t.Amount, t.AccountId)
	}
	// in case of error rollback
	if err != nil {
		tx.Rollback()
		logger.Error("error while saving transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	// commit the transaction when all good
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("error while committing the transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("error while getting last transactionID: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	account, appErr := d.FindBy(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}
	t.TransactionId = strconv.FormatInt(transactionId, 10)
	t.Amount = account.Amount

	return &t, nil
}
