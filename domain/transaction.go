package domain

import "github.com/nitinkumar91296/banking/dto"

const WITHDRAWL = "widthdrawl"

type Transaction struct {
	TransactionId   string  `db:"transaction_id"`
	AccountId       string  `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	TransactionDate string  `db:"transaction_date"`
}

func (t Transaction) IsWithdrawl() bool {
	if t.TransactionType == WITHDRAWL {
		return true
	}
	return false
}

func (t Transaction) ToDto() *dto.TransactionResponse {
	return &dto.TransactionResponse{
		AccountId:       t.AccountId,
		TransactionType: t.TransactionType,
		Amount:          t.Amount,
		TransactionId:   t.TransactionId,
		TransactionDate: t.TransactionDate,
	}
}
