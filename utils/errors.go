package utils

import "fmt"

var (
	ErrInvalidUUID        = fmt.Errorf("invalid UUID")
	ErrAccountNotFound    = fmt.Errorf("account not found")
	ErrInvalidTxType      = fmt.Errorf("invalid transaction type")
	ErrInsufficientFunds  = fmt.Errorf("insufficient funds")
	ErrInvalidRequestBody = fmt.Errorf("invalid request body")
)
