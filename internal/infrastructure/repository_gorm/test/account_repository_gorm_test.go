package repository_gorm_test

import (
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/model"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/infrastructure/repository"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/infrastructure/repository_gorm"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type AccountRepositoryGormTestSuite struct {
	suite.Suite
	db      *gorm.DB
	repo    repository.AccountRepository
	sqlmock sqlmock.Sqlmock
	dbmock  *gorm.DB
}

func (suite *AccountRepositoryGormTestSuite) SetupTest() {
	suite.configureTestWithMock()
}

func (suite *AccountRepositoryGormTestSuite) configureTestWithMock() {
	// Mock database setup
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp)) // Use regular expression matching
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
	suite.repo = repository_gorm.NewAccountRepository(suite.dbmock)

}

func buildAccount() *model.Account {
	return &model.Account{
		// ID:             1,
		DocumentNumber: "123456789",
		CreatedAt:      time.Date(2024, time.December, 21, 10, 30, 0, 0, time.UTC),
		UpdatedAt:      time.Date(2024, time.December, 21, 10, 30, 0, 0, time.UTC),
	}
}

func (suite *AccountRepositoryGormTestSuite) TestCreate_Success() {
	// Given: prepare a valid account
	account := buildAccount()

	// When: mock the CREATE query
	suite.sqlmock.ExpectBegin()
	suite.sqlmock.ExpectQuery("INSERT INTO \"accounts\"").
		WithArgs(account.DocumentNumber, account.CreatedAt, account.UpdatedAt).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(account.ID))
	suite.sqlmock.ExpectCommit()

	// Then: execute the Create method
	result, err := suite.repo.Create(account)

	// Assert
	suite.NoError(err)
	suite.NotNil(result)
	suite.Equal(account.ID, result.ID)
	suite.Equal(account.DocumentNumber, result.DocumentNumber)
	suite.Equal(account.CreatedAt, result.CreatedAt)
	suite.Equal(account.UpdatedAt, result.UpdatedAt)
}

func (suite *AccountRepositoryGormTestSuite) TestCreate_Fail() {
	// Given: prepare an invalid account
	account := &model.Account{
		ID:             0,
		DocumentNumber: "",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
	}

	// When: mock the CREATE query with an error
	suite.sqlmock.ExpectBegin()
	suite.sqlmock.ExpectQuery("INSERT INTO \"accounts\"").
		WithArgs(account.DocumentNumber, account.CreatedAt, account.UpdatedAt).
		WillReturnError(gorm.ErrInvalidData)
	suite.sqlmock.ExpectRollback()

	// Then: execute the Create method and expect an error
	result, err := suite.repo.Create(account)

	// Assert
	suite.Error(err)
	suite.Nil(result)
}

func (suite *AccountRepositoryGormTestSuite) TestFindByID_Success() {
	// Given: prepare a valid account
	account := buildAccount()

	suite.sqlmock.ExpectQuery(`SELECT \* FROM "accounts" WHERE "accounts"."id" = \$1 ORDER BY "accounts"."id" LIMIT \$2`).
		WithArgs(account.ID, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "document_number", "created_at", "updated_at"}).
			AddRow(account.ID, account.DocumentNumber, account.CreatedAt, account.UpdatedAt))

	// Then: execute the FindByID method
	result, err := suite.repo.FindByID(account.ID)

	// Assert
	suite.NoError(err)
	suite.NotNil(result)
	suite.Equal(account.ID, result.ID)
	suite.Equal(account.DocumentNumber, result.DocumentNumber)
	suite.Equal(account.CreatedAt, result.CreatedAt)
	suite.Equal(account.UpdatedAt, result.UpdatedAt)
}

func (suite *AccountRepositoryGormTestSuite) TestFindByID_Fail() {
	// Given: an invalid account ID
	accountID := uint64(999)

	// When: mock the SELECT query with an error
	suite.sqlmock.ExpectQuery("SELECT * FROM \"accounts\" WHERE \"accounts\".\"id\" = ?").
		WithArgs(accountID).
		WillReturnError(gorm.ErrRecordNotFound)

	// Then: execute the FindByID method and expect an error
	result, err := suite.repo.FindByID(accountID)

	// Assert
	suite.Error(err)
	suite.Nil(result)
}

func TestAccountRepositoryGormTestSuite(t *testing.T) {
	suite.Run(t, new(AccountRepositoryGormTestSuite))
}
