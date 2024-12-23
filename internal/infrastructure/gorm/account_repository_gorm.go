package gorm

import (
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/model"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/infrastructure/repository"
	"github.com/priscila-albertini-da-silva/transactions-routine/pkg/gormfx"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type AccountRepositoryGorm struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) repository.AccountRepository {
	return &AccountRepositoryGorm{db}
}

func (r *AccountRepositoryGorm) Create(account *model.Account) (*model.Account, error) {
	if err := r.db.Create(&account).Error; err != nil {
		return nil, err
	}
	return account, nil
}

func (r *AccountRepositoryGorm) FindByID(id uint64) (*model.Account, error) {
	var account model.Account
	if err := r.db.First(&account, id).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

var ModuleAccountRepositoryGorm = fx.Options(
	gormfx.ModuleGorm,
	fx.Provide(NewAccountRepository),
)
