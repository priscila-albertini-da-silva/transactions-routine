package usecase_test

import (
	"errors"
	"testing"

	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/entity"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/model"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/usecase"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AccountUseCaseTestSuite struct {
	suite.Suite
	mockRepo      *MockAccountRepository
	mockValidator *MockAccountValidator
	useCase       usecase.AccountUseCase
}

func (suite *AccountUseCaseTestSuite) SetupTest() {
	suite.mockRepo = new(MockAccountRepository)
	suite.mockValidator = new(MockAccountValidator)
	suite.useCase = usecase.NewAccountUseCase(suite.mockRepo, suite.mockValidator)
}

func (suite *AccountUseCaseTestSuite) TestCreateAccount_Success() {
	documentNumber := "12345678901"
	account := entity.NewAccount(documentNumber)

	// Set up the mocks
	suite.mockValidator.On("Validate", account).Return(nil)
	suite.mockRepo.On("Create", mock.Anything).Return(&model.Account{
		ID:             1,
		DocumentNumber: documentNumber,
	}, nil)

	// Call CreateAccount
	result, err := suite.useCase.CreateAccount(*account)

	// Assert results
	suite.NoError(err)
	suite.NotNil(result)
	suite.Equal(uint64(1), result.ID)
	suite.Equal(documentNumber, result.DocumentNumber)

	// Verify mocks were called
	suite.mockValidator.AssertExpectations(suite.T())
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *AccountUseCaseTestSuite) TestCreateAccount_ValidationError() {
	documentNumber := "12345678901"
	account := entity.NewAccount(documentNumber)

	// Set up the mocks
	suite.mockValidator.On("Validate", account).Return(errors.New("validation error"))

	// Call CreateAccount
	result, err := suite.useCase.CreateAccount(*account)

	// Assert results
	suite.Error(err)
	suite.Nil(result)

	// Verify mocks were called
	suite.mockValidator.AssertExpectations(suite.T())
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *AccountUseCaseTestSuite) TestGetAccountByID_Success() {
	accountModel := &model.Account{
		ID:             1,
		DocumentNumber: "12345678901",
	}

	// Set up the mock
	suite.mockRepo.On("FindByID", uint64(1)).Return(accountModel, nil)

	// Call GetAccountByID
	result, err := suite.useCase.GetAccountByID(1)

	// Assert results
	suite.NoError(err)
	suite.NotNil(result)
	suite.Equal(uint64(1), result.ID)
	suite.Equal("12345678901", result.DocumentNumber)

	// Verify mocks were called
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *AccountUseCaseTestSuite) TestGetAccountByID_NotFound() {
	// Set up the mock
	suite.mockRepo.On("FindByID", uint64(1)).Return(&model.Account{}, errors.New("account not found"))

	// Call GetAccountByID
	result, err := suite.useCase.GetAccountByID(1)

	// Assert results
	suite.Error(err)
	suite.Nil(result)

	// Verify mocks were called
	suite.mockRepo.AssertExpectations(suite.T())
}

func TestAccountUseCase(t *testing.T) {
	suite.Run(t, new(AccountUseCaseTestSuite))
}

type MockAccountRepository struct {
	mock.Mock
}

func (m *MockAccountRepository) Create(account *model.Account) (*model.Account, error) {
	args := m.Called(account)
	return args.Get(0).(*model.Account), args.Error(1)
}

func (m *MockAccountRepository) FindByID(id uint64) (*model.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Account), args.Error(1)
}

type MockAccountValidator struct {
	mock.Mock
}

func (m *MockAccountValidator) Validate(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}
