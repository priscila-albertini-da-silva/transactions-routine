package model

import "time"

type Transaction struct {
	ID              uint64
	AccountID       uint64  `gorm:"not null"`
	OperationTypeID uint64  `gorm:"not null"`
	Amount          float64 `gorm:"not null"`
	EventDate       time.Time
}

func NewTransaction(accountID, operationTypeID uint64, amount float64) *Transaction {
	return &Transaction{
		AccountID:       accountID,
		OperationTypeID: operationTypeID,
		Amount:          amount,
		EventDate:       time.Now(),
	}
}
