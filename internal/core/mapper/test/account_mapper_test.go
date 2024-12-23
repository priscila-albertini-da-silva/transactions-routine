package mapper_test

import (
	"testing"

	"github.com/priscila-albertini-da-silva/transactions-routine/internal/adapter/repository/mapper"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/adapter/repository/model"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/entity"
	"github.com/stretchr/testify/suite"
)

type AccountMapperTestSuite struct {
	suite.Suite
}

func (suite *AccountMapperTestSuite) TestAccountToModel() {
	// given: prepare entity.Account
	account := buildAccountEntity()

	// when: map entity to model
	result := mapper.AccountToModel(account)

	// then: assert the expected result
	suite.Equal(account.ID, result.ID)
	suite.Equal(account.DocumentNumber, result.DocumentNumber)
}

func (suite *AccountMapperTestSuite) TestAccountToEntity() {
	// given: prepare model.Account
	account := buildAccountModel()

	// when: map model to entity
	result := mapper.AccountToEntity(account)

	// then: assert the expected result
	suite.Equal(account.ID, result.ID)
	suite.Equal(account.DocumentNumber, result.DocumentNumber)
}

func (suite *AccountMapperTestSuite) TestAccountToModel_EmptyValues() {
	// given: prepare entity.Account with empty or zero values
	account := buildAccountEntityWithEmptyValues()

	// when: map entity to model
	result := mapper.AccountToModel(account)

	// then: assert the expected result
	suite.Equal(account.ID, result.ID)
	suite.Equal(account.DocumentNumber, result.DocumentNumber)
}

func (suite *AccountMapperTestSuite) TestAccountToEntity_EmptyValues() {
	// given: prepare entity.Account with empty or zero values
	account := buildAccountModelWithEmptyValues()

	// when: map entity to model
	result := mapper.AccountToEntity(account)

	// then: assert the expected result
	suite.Equal(account.ID, result.ID)
	suite.Equal(account.DocumentNumber, result.DocumentNumber)
}

func TestAccountMapper(t *testing.T) {
	suite.Run(t, new(AccountMapperTestSuite))
}

func buildAccountEntity() entity.Account {
	return entity.Account{
		ID:             1,
		DocumentNumber: "123456789",
	}
}

func buildAccountModel() model.Account {
	return model.Account{
		ID:             1,
		DocumentNumber: "123456789",
	}
}

func buildAccountEntityWithEmptyValues() entity.Account {
	return entity.Account{
		ID:             0,
		DocumentNumber: "",
	}
}

func buildAccountModelWithEmptyValues() model.Account {
	return model.Account{
		ID:             0,
		DocumentNumber: "",
	}
}
