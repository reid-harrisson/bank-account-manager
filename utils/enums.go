package utils

import "fmt"

type TransactionType int8

const (
	Deposit TransactionType = iota
	Withdrawal
	Invalid
)

func (value TransactionType) String() string {
	if value == Deposit {
		return "deposit"
	} else if value == Withdrawal {
		return "withdrawal"
	}
	return "invalid"
}

func ParseTransactionType(value string) (TransactionType, error) {
	switch value {
	case "deposit":
		return Deposit, nil
	case "withdrawal":
		return Withdrawal, nil
	}
	return Invalid, fmt.Errorf("invalid transaction type: %s", value)
}
