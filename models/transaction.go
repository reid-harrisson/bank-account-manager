package models

import (
	"bank-account-manager/utils"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID        uuid.UUID
	AccountID uuid.UUID
	Type      utils.TransactionType
	Amount    float64
	TimeStamp time.Time
}
