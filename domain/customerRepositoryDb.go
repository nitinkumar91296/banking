package domain

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	errs "github.com/nitinkumar91296/banking/errors"
	"github.com/nitinkumar91296/banking/logger"
)

type CustomerReposityDb struct {
	client *sqlx.DB
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerReposityDb {
	return CustomerReposityDb{client: dbClient}
}

func (d CustomerReposityDb) FindAll(status string) ([]Customer, *errs.AppError) {

	customers := make([]Customer, 0)
	var err error

	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)
		// rows, err = d.client.Query(findAllSql) // used with sql package
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		logger.Error("error while querying customer table: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected db error")
	}

	return customers, nil
}

func (d CustomerReposityDb) FindCustomerById(id string) (*Customer, *errs.AppError) {

	findCustomerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	var c Customer
	err := d.client.Get(&c, findCustomerSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Customer not found" + err.Error())
			return nil, errs.NewNotFoundError("customer Not Found")
		} else {
			logger.Error("Error while scanning customer by ID" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected db error")
		}
	}

	return &c, nil
}
