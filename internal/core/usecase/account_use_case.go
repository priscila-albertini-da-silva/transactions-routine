package usecase

import (
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/entity"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/mapper"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/infrastructure/repository"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/infrastructure/repositorygorm"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/validator"
	"go.uber.org/fx"
)

type AccountUseCase interface {
	CreateAccount(account entity.Account) (*entity.Account, error)
	GetAccountByID(id uint64) (*entity.Account, error)
}

type AccountUseCaseImpl struct {
	accountRepository repository.AccountRepository
	validator         validator.AccountValidator
}

func NewAccountUseCase(accountRepo repository.AccountRepository, validator validator.AccountValidator) AccountUseCase {
	return &AccountUseCaseImpl{
		accountRepository: accountRepo,
		validator:         validator,
	}
}

func (uc *AccountUseCaseImpl) CreateAccount(account entity.Account) (*entity.Account, error) {
	err := uc.validator.Validate(&account)
	if err != nil {
		return nil, err
	}

	accountModel := mapper.AccountToModel(account)

	savedAccountModel, err := uc.accountRepository.Create(&accountModel)
	if err != nil {
		return nil, err
	}

	savedAccountEntity := mapper.AccountToEntity(*savedAccountModel)

	return &savedAccountEntity, nil
}

func (uc *AccountUseCaseImpl) GetAccountByID(id uint64) (*entity.Account, error) {
	accountModel, err := uc.accountRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	accountEntity := mapper.AccountToEntity(*accountModel)

	return &accountEntity, nil
}

var ModuleAccountUseCase = fx.Options(
	repositorygorm.ModuleAccountRepositoryGorm,
	validator.ModuleAccountValidator,
	fx.Provide(NewAccountUseCase),
)
