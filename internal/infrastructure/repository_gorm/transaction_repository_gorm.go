package repository_gorm

import (
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/model"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/infrastructure/repository"
	"github.com/priscila-albertini-da-silva/transactions-routine/pkg/gormfx"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type TransactionRepositoryGorm struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) repository.TransactionRepository {
	return &TransactionRepositoryGorm{db}
}

func (r *TransactionRepositoryGorm) Create(transaction *model.Transaction) (*model.Transaction, error) {
	if err := r.db.Create(&transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

var ModuleTransactionRepositoryGorm = fx.Options(
	gormfx.ModuleGorm,
	fx.Provide(NewTransactionRepository),
)
