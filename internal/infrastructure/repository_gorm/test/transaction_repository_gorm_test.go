package repository_gorm_test

import (
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/model"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/infrastructure/repository"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/infrastructure/repository_gorm"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/infrastructure/repository_gorm/test/database"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type TransactionRepositoryGormTestSuite struct {
	suite.Suite
	db      *gorm.DB
	repo    repository.TransactionRepository
	sqlmock sqlmock.Sqlmock
	dbmock  *gorm.DB
}

func (suite *TransactionRepositoryGormTestSuite) SetupTest() {
	suite.db = database.NewDatabaseConnection("user='postgres' dbname='transactions_routine' host='localhost' password='postgres' port='5432' sslmode='disable'").DB
	err := suite.db.AutoMigrate(&model.Account{})
	if err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}
	suite.repo = repository_gorm.NewTransactionRepository(suite.db)
	suite.configureTestWithMock()
}

func (suite *TransactionRepositoryGormTestSuite) configureTestWithMock() {
	// Mock database setup
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	dialector := postgres.New(postgres.Config{
		Conn: db,
	})

	suite.dbmock, _ = gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	suite.sqlmock = mock
	suite.repo = repository_gorm.NewTransactionRepository(suite.dbmock)

	// Expect Begin
	suite.sqlmock.ExpectBegin()

	// Expect the INSERT query with the correct parameters
	suite.sqlmock.ExpectQuery("INSERT INTO \"transactions\"").
		WithArgs(123, 456, 100.0, time.Date(2024, time.December, 21, 10, 30, 0, 0, time.UTC), 1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	// Expect Commit after the INSERT
	suite.sqlmock.ExpectCommit()

	// Expect Close for the connection
	suite.sqlmock.ExpectClose()
}

func buildTransaction() *model.Transaction {
	return &model.Transaction{
		ID:              1,
		AccountID:       123,
		OperationTypeID: 456,
		Amount:          100.0,
		EventDate:       time.Date(2024, time.December, 21, 10, 30, 0, 0, time.UTC),
	}
}

func (suite *TransactionRepositoryGormTestSuite) TestCreate_Success() {
	// Given: prepare a valid transaction
	transaction := buildTransaction()

	// When: call the Create method
	result, err := suite.repo.Create(transaction)

	// Then: assert that there was no error and the transaction was created correctly
	suite.NoError(err)
	suite.NotNil(result)
	suite.Equal(transaction.ID, result.ID)
	suite.Equal(transaction.Amount, result.Amount)
	suite.Equal(transaction.AccountID, result.AccountID)
	suite.Equal(transaction.OperationTypeID, result.OperationTypeID)
	suite.True(transaction.EventDate.Equal(result.EventDate))
}

func (suite *TransactionRepositoryGormTestSuite) TestCreate_Fail_DuplicateID() {
	// Given: prepare a valid transaction
	transaction := buildTransaction()

	// Create the first transaction
	_, err := suite.repo.Create(transaction)
	suite.NoError(err)

	// When: try to create a transaction with the same ID
	_, err = suite.repo.Create(transaction)

	// Then: assert that a duplicate key error occurred
	suite.Error(err)
}

func (suite *TransactionRepositoryGormTestSuite) TestCreate_Fail_InvalidData() {
	// Given: prepare a transaction with invalid data
	transaction := &model.Transaction{
		ID:              0,
		AccountID:       0,
		OperationTypeID: 0,
		Amount:          0.0,
		EventDate:       time.Time{},
	}

	// When: try to create the transaction with invalid data
	result, err := suite.repo.Create(transaction)

	// Then: assert that an error occurred due to invalid data
	suite.Error(err)
	suite.Nil(result)
}

func TestTransactionRepositoryGormTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionRepositoryGormTestSuite))
}
