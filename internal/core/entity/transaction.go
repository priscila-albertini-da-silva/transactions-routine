package entity

import (
	"errors"
	"time"
)

type Transaction struct {
	ID              uint64
	AccountID       uint64
	OperationTypeID uint64
	Amount          float64
	EventDate       time.Time
}

func NewTransaction(accountID uint64, operationTypeID uint64, amount float64) (*Transaction, error) {
	operationType, err := NewOperationType(operationTypeID)
	if err != nil {
		return nil, err
	}

	adjustedAmount, err := AdjustAmountForOperation(amount, operationType)
	if err != nil {
		return nil, err
	}

	transaction := &Transaction{
		AccountID:       accountID,
		OperationTypeID: operationTypeID,
		Amount:          adjustedAmount,
		EventDate:       time.Now(),
	}

	return transaction, nil
}

func AdjustAmountForOperation(amount float64, operationType OperationType) (float64, error) {
	switch operationType {
	case NormalPurchase, PurchaseWithInstallments, Withdrawal:
		if amount > 0 {
			return -amount, nil
		}
	case CreditVoucher:
		if amount < 0 {
			return -amount, nil
		}
	default:
		return 0, errors.New("invalid operation type for adjustment")
	}
	return amount, nil
}
