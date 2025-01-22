package utils

type TransactionType int8

const (
	Deposit TransactionType = iota
	Withdrawal
)
