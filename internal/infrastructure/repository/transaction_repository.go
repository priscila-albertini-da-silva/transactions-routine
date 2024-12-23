package repository

import (
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/model"
)

type TransactionRepository interface {
	Create(transaction *model.Transaction) (*model.Transaction, error)
}
