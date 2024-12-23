package mapper

import (
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/entity"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/model"
)

func AccountToModel(account entity.Account) model.Account {
	return model.Account{
		ID:             account.ID,
		DocumentNumber: account.DocumentNumber,
	}
}

func AccountToEntity(account model.Account) entity.Account {
	return entity.Account{
		ID:             account.ID,
		DocumentNumber: account.DocumentNumber,
	}
}
