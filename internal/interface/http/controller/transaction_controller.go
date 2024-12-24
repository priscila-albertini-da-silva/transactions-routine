package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/entity"
	"github.com/priscila-albertini-da-silva/transactions-routine/internal/core/usecase"
)

type TransactionController interface {
	CreateTransaction(c *gin.Context)
}

type TransactionControllerImpl struct {
	transactionUseCase usecase.TransactionUseCase
}

func NewTransactionController(transactionUseCase usecase.TransactionUseCase) TransactionController {
	return &TransactionControllerImpl{transactionUseCase: transactionUseCase}
}

func (tc *TransactionControllerImpl) CreateTransaction(c *gin.Context) {
	var input struct {
		AccountID       uint64  `json:"account_id" binding:"required"`
		OperationTypeID uint64  `json:"operation_type_id" binding:"required"`
		Amount          float64 `json:"amount" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	transaction, err := entity.NewTransaction(input.AccountID, input.OperationTypeID, input.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res, err := tc.transactionUseCase.Create(*transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}
