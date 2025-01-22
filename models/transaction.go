package models

import (
	"bank-account-manager/utils"

	"github.com/google/uuid"
)

type Transaction struct {
	ID        uuid.UUID
	AccountID uuid.UUID
	Type      utils.TransactionType
	Owner     string
	Balance   float64
}
