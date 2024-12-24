package model

import (
	"time"
)

type Account struct {
	ID             uint64
	DocumentNumber string `gorm:"not null;unique"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewAccount(documentNumber string) *Account {
	return &Account{
		DocumentNumber: documentNumber,
	}
}
