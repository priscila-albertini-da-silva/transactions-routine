package model

type OperationType struct {
	ID          uint64
	Description string
}

func NewOperationType(description string) *OperationType {
	return &OperationType{Description: description}
}
