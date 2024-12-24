package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/entity"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/usecase"
)

type AccountController interface {
	CreateAccount(c *gin.Context)
	GetAccount(c *gin.Context)
}

type AccountControllerImpl struct {
	accountUseCase usecase.AccountUseCase
}

func NewAccountController(accountUseCase usecase.AccountUseCase) AccountController {
	return &AccountControllerImpl{accountUseCase: accountUseCase}
}

func (ac *AccountControllerImpl) CreateAccount(c *gin.Context) {
	var input struct {
		DocumentNumber string `json:"document_number" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	account := entity.NewAccount(input.DocumentNumber)

	res, err := ac.accountUseCase.CreateAccount(*account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (ac *AccountControllerImpl) GetAccount(c *gin.Context) {
	accountIdStr := c.Param("accountId")
	accountId, err := strconv.ParseUint(accountIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid accountId"})
		return
	}

	account, err := ac.accountUseCase.GetAccountByID(accountId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	c.JSON(http.StatusOK, account)
}
