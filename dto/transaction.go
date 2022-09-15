package dto

import (
	"strings"

	errs "github.com/nitinkumar91296/banking/errors"
)

type TransactionRequest struct {
	AccountId       string
	CustomerId      string
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
}

type TransactionResponse struct {
	TransactionId   string  `json:"transaction_id"`
	AccountId       string  `json:"account_id"`
	TransactionDate string  `json:"transaction_date"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"new_balance"`
}

func (r TransactionRequest) Validate() *errs.AppError {

	if r.Amount <= 0 {
		return errs.NewValidationError("Amount can't be less than 0")
	}
	if strings.ToLower(r.TransactionType) != "deposit" && strings.ToLower(r.TransactionType) != "widthdrawl" {
		return errs.NewValidationError("Transaction Type should be deposit or withdrawl")

	}
	return nil
}

func (r TransactionRequest) IsTransactionTypeWithdrawl() bool {
	return strings.ToLower(r.TransactionType) == "withdrawl"
}
