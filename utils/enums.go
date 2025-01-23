package utils

type TransactionType int8

const (
	Deposit TransactionType = iota
	Withdrawal
	Invalid
)

func (value TransactionType) String() string {
	switch value {
	case Deposit:
		return "deposit"
	case Withdrawal:
		return "withdrawal"
	}
	return ""
}

func Parse(value string) TransactionType {
	switch value {
	case "deposit":
		return Deposit
	case "withdrawal":
		return Withdrawal
	}
	return Invalid
}
