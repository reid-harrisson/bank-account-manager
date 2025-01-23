package utils

import "fmt"

var (
	ErrInvalidUUID            = fmt.Errorf("invalid UUID")
	ErrAccountNotFound        = fmt.Errorf("account not found")
	ErrInvalidTransactionType = fmt.Errorf("invalid transaction type")
	ErrInsufficientFunds      = fmt.Errorf("insufficient funds")
)
