package mapper_test

import (
	"testing"
	"time"

	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/entity"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/mapper"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/model"
	"github.com/stretchr/testify/suite"
)

type TransactionMapperTestSuite struct {
	suite.Suite
}

func (suite *TransactionMapperTestSuite) TestTransactionToModel_Values() {
	// given: prepare transaction entity
	transaction := buildTransactionEntity()

	// when: map entity to model
	result := mapper.TransactionToModel(transaction)

	// then: assert the expected result
	expectedModel := buildTransactionModel()

	suite.Equal(expectedModel.ID, result.ID)
	suite.Equal(expectedModel.AccountID, result.AccountID)
	suite.Equal(expectedModel.OperationTypeID, result.OperationTypeID)
	suite.Equal(expectedModel.Amount, result.Amount)
	suite.Equal(expectedModel.EventDate, result.EventDate)
}

func (suite *TransactionMapperTestSuite) TestTransactionToEntity_Values() {
	// given: prepare transaction model
	transaction := buildTransactionModel()

	// when: map model to entity
	result := mapper.TransactionToEntity(transaction)

	// then: assert the expected result
	expectedEntity := buildTransactionEntity()

	suite.Equal(expectedEntity.ID, result.ID)
	suite.Equal(expectedEntity.AccountID, result.AccountID)
	suite.Equal(expectedEntity.OperationTypeID, result.OperationTypeID)
	suite.Equal(expectedEntity.Amount, result.Amount)
	suite.Equal(expectedEntity.EventDate, result.EventDate)
}

func (suite *TransactionMapperTestSuite) TestTransactionToModel_EmptyValues() {
	// given: prepare transaction entity with empty values
	transaction := buildTransactionEntityEmpty()

	// when: map entity to model
	result := mapper.TransactionToModel(transaction)

	// then: assert the expected result
	expectedModel := buildTransactionModelEmpty()

	suite.Equal(expectedModel.ID, result.ID)
	suite.Equal(expectedModel.AccountID, result.AccountID)
	suite.Equal(expectedModel.OperationTypeID, result.OperationTypeID)
	suite.Equal(expectedModel.Amount, result.Amount)
	suite.Equal(expectedModel.EventDate, result.EventDate)
}

func (suite *TransactionMapperTestSuite) TestTransactionToEntity_EmptyValues() {
	// given: prepare transaction model with empty values
	transaction := buildTransactionModelEmpty()

	// when: map model to entity
	result := mapper.TransactionToEntity(transaction)

	// then: assert the expected result
	expectedEntity := buildTransactionEntityEmpty()

	suite.Equal(expectedEntity.ID, result.ID)
	suite.Equal(expectedEntity.AccountID, result.AccountID)
	suite.Equal(expectedEntity.OperationTypeID, result.OperationTypeID)
	suite.Equal(expectedEntity.Amount, result.Amount)
	suite.Equal(expectedEntity.EventDate, result.EventDate)
}

func TestTransactionMapperTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionMapperTestSuite))
}

func buildTransactionEntity() entity.Transaction {
	return entity.Transaction{
		ID:              123,
		AccountID:       456,
		OperationTypeID: 789,
		Amount:          100.50,
		EventDate:       buildEventDate(),
	}
}

func buildTransactionModel() model.Transaction {
	return model.Transaction{
		ID:              123,
		AccountID:       456,
		OperationTypeID: 789,
		Amount:          100.50,
		EventDate:       buildEventDate(),
	}
}

func buildTransactionModelEmpty() model.Transaction {
	return model.Transaction{
		ID:              0,
		AccountID:       0,
		OperationTypeID: 0,
		Amount:          0,
		EventDate:       time.Time{},
	}
}

func buildTransactionEntityEmpty() entity.Transaction {
	return entity.Transaction{
		ID:              0,
		AccountID:       0,
		OperationTypeID: 0,
		Amount:          0,
		EventDate:       time.Time{},
	}
}

func buildEventDate() time.Time {
	return time.Date(2024, time.December, 21, 10, 30, 0, 0, time.UTC)
}
