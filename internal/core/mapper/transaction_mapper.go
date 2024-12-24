package mapper

import (
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/entity"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/model"
)

func TransactionToEntity(m model.Transaction) entity.Transaction {
	return entity.Transaction{
		ID:              m.ID,
		AccountID:       m.AccountID,
		OperationTypeID: m.OperationTypeID,
		Amount:          m.Amount,
		EventDate:       m.EventDate,
	}
}

func TransactionToModel(e entity.Transaction) model.Transaction {
	return model.Transaction{
		ID:              e.ID,
		AccountID:       e.AccountID,
		OperationTypeID: e.OperationTypeID,
		Amount:          e.Amount,
		EventDate:       e.EventDate,
	}
}
