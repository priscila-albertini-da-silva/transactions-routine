package usecase_test

import (
	"errors"
	"testing"

	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/entity"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/mapper"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/model"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/usecase"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TransactionUseCaseTestSuite struct {
	suite.Suite
	mockTransactionRepo *MockTransactionRepository
	mockAccountRepo     *MockAccountRepository
	useCase             usecase.TransactionUseCase
}

func (suite *TransactionUseCaseTestSuite) SetupTest() {
	suite.mockTransactionRepo = new(MockTransactionRepository)
	suite.mockAccountRepo = new(MockAccountRepository)
	suite.useCase = usecase.NewTransactionUseCase(suite.mockTransactionRepo, suite.mockAccountRepo)
}

func (suite *TransactionUseCaseTestSuite) TestCreate_Success() {
	transaction := &entity.Transaction{
		AccountID:       1,
		OperationTypeID: 2,
		Amount:          100.0,
	}

	// Prepare the expected transaction entity and model
	transactionModel := mapper.TransactionToModel(*transaction)

	savedTransactionModel := &transactionModel
	savedTransactionEntity := mapper.TransactionToEntity(*savedTransactionModel)

	// Set up mocks
	suite.mockTransactionRepo.On("Create", mock.Anything).Return(savedTransactionModel, nil)

	// Call Create
	result, err := suite.useCase.Create(*transaction)

	// Assert results
	suite.NoError(err)
	suite.NotNil(result)
	suite.Equal(savedTransactionEntity.ID, result.ID)
	suite.Equal(savedTransactionEntity.Amount, result.Amount)

	// Verify mocks were called
	suite.mockTransactionRepo.AssertExpectations(suite.T())
	suite.mockAccountRepo.AssertExpectations(suite.T())
}

func (suite *TransactionUseCaseTestSuite) TestCreate_RepositoryError() {
	transaction := &entity.Transaction{
		AccountID:       1,
		OperationTypeID: 2,
		Amount:          100.0,
	}

	// Set up the mock to return an error from the repository
	suite.mockTransactionRepo.On("Create", mock.Anything).Return(&model.Transaction{}, errors.New("repository error"))

	// Call Create
	result, err := suite.useCase.Create(*transaction)

	// Assert results
	suite.Error(err)
	suite.Nil(result)

	// Verify mocks were called
	suite.mockTransactionRepo.AssertExpectations(suite.T())
	suite.mockAccountRepo.AssertExpectations(suite.T())
}

func TestTransactionUseCase(t *testing.T) {
	suite.Run(t, new(TransactionUseCaseTestSuite))
}

type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) Create(transaction *model.Transaction) (*model.Transaction, error) {
	args := m.Called(transaction)
	return args.Get(0).(*model.Transaction), args.Error(1)
}
