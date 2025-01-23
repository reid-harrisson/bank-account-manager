package storage

import (
	"bank-account-manager/models"

	"github.com/google/uuid"
)

type Storage struct {
	Accounts       []models.Account
	Transactions   map[uuid.UUID][]models.Transaction
	AccountIndices map[uuid.UUID]int
}

func Create() *Storage {
	accounts := []models.Account{}
	transactions := map[uuid.UUID][]models.Transaction{}
	accountIndices := map[uuid.UUID]int{}

	return &Storage{
		Accounts:       accounts,
		AccountIndices: accountIndices,
		Transactions:   transactions,
	}
}
