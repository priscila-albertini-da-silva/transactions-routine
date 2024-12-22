package entity

type Account struct {
	ID             uint64
	DocumentNumber string
}

func NewAccount(documentNumber string) *Account {
	return &Account{DocumentNumber: documentNumber}
}
