package storage

import (
	"bank-account-manager/models"

	"github.com/google/uuid"
)

type Storage struct {
	Accounts     *map[uuid.UUID]models.Account
	Transactions *map[uuid.UUID]models.Transaction
}

func Create() *Storage {
	accounts := map[uuid.UUID]models.Account{}
	transactions := map[uuid.UUID]models.Transaction{}

	return &Storage{
		Accounts:     &accounts,
		Transactions: &transactions,
	}
}
