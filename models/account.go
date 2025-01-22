package models

import "github.com/google/uuid"

type Account struct {
	ID      uuid.UUID
	Owner   string
	Balance float64
}
