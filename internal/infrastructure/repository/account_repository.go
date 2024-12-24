package repository

import "github.com/priscila-albertini-da-silva/transactions-routine/internal/core/model"

type AccountRepository interface {
	Create(account *model.Account) (*model.Account, error)
	FindByID(id uint64) (*model.Account, error)
}
