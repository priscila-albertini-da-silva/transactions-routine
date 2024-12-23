package entity

import "time"

type Transaction struct {
	ID              uint64
	AccountID       uint64
	OperationTypeID uint64
	Amount          float64
	EventDate       time.Time
}

func NewTransaction(accountID, operationTypeID uint64, amount float64) *Transaction {
	return &Transaction{
		AccountID:       accountID,
		OperationTypeID: operationTypeID,
		Amount:          amount,
	}
}
