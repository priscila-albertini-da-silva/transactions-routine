package entity

import "github.com/priscila-albertini-da-silva/transactions-routine/internal/errors"

type OperationType int

const (
	NormalPurchase OperationType = iota + 1
	PurchaseWithInstallments
	Withdrawal
	CreditVoucher
)

func (ot OperationType) IsValid() bool {
	switch ot {
	case NormalPurchase, PurchaseWithInstallments, Withdrawal, CreditVoucher:
		return true
	default:
		return false
	}
}

func NewOperationType(operationTypeID uint64) (OperationType, error) {
	ot := OperationType(operationTypeID)
	if !ot.IsValid() {
		return 0, errors.NewValidationError("invalid operation type")
	}

	return ot, nil
}
