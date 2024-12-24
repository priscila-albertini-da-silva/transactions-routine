package usecase

import (
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/entity"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/mapper"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/infrastructure/repository"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/infrastructure/repositorygorm"
	"go.uber.org/fx"
)

type TransactionUseCase interface {
	Create(entity.Transaction) (*entity.Transaction, error)
}

type TransactionUseCaseImpl struct {
	transactionRepository repository.TransactionRepository
	accountRepository     repository.AccountRepository
}

func NewTransactionUseCase(transactionRepository repository.TransactionRepository, accountRepository repository.AccountRepository) TransactionUseCase {
	return &TransactionUseCaseImpl{
		transactionRepository,
		accountRepository,
	}
}

func (u *TransactionUseCaseImpl) Create(transaction entity.Transaction) (*entity.Transaction, error) {
	accountModel := mapper.TransactionToModel(transaction)

	savedTransactionModel, err := u.transactionRepository.Create(&accountModel)
	if err != nil {
		return nil, err
	}

	savedTransactionEntity := mapper.TransactionToEntity(*savedTransactionModel)

	return &savedTransactionEntity, nil
}

var ModuleTransactionUseCase = fx.Options(
	repositorygorm.ModuleAccountRepositoryGorm,
	repositorygorm.ModuleTransactionRepositoryGorm,
	fx.Provide(NewTransactionUseCase),
)
