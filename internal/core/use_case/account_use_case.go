package usecase

import (
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/entity"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/mapper"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/infrastructure/repository"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/infrastructure/repository_gorm"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/validator"
	"go.uber.org/fx"
)

type AccountUseCase struct {
	accountRepo repository.AccountRepository
	validator   *validator.AccountValidator
}

func NewAccountUseCase(accountRepo repository.AccountRepository, validator *validator.AccountValidator) *AccountUseCase {
	return &AccountUseCase{
		accountRepo: accountRepo,
		validator:   validator,
	}
}

func (uc *AccountUseCase) CreateAccount(documentNumber string) (*entity.Account, error) {
	account := entity.NewAccount(documentNumber)

	err := uc.validator.Validate(account)
	if err != nil {
		return nil, err
	}

	accountModel := mapper.AccountToModel(*account)

	savedAccountModel, err := uc.accountRepo.Create(&accountModel)
	if err != nil {
		return nil, err
	}

	savedAccountEntity := mapper.AccountToEntity(*savedAccountModel)

	return &savedAccountEntity, nil
}

func (uc *AccountUseCase) GetAccountByID(id uint64) (*entity.Account, error) {
	accountModel, err := uc.accountRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	accountEntity := mapper.AccountToEntity(*accountModel)

	return &accountEntity, nil
}

var ModuleGenerateInvoiceExtractUseCase = fx.Options(
	repository_gorm.ModuleAccountRepositoryGorm,
	validator.ModuleAccountValidator,
)
