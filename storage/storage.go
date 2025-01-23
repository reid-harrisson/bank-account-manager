package storage

import (
	"bank-account-manager/models"
	"sync"

	"github.com/google/uuid"
)

type Storage struct {
	Accounts       []models.Account
	Transactions   map[uuid.UUID][]models.Transaction
	AccountIndices map[uuid.UUID]int
	Mutex          *sync.Mutex
}

func Create() *Storage {
	accounts := []models.Account{}
	transactions := map[uuid.UUID][]models.Transaction{}
	accountIndices := map[uuid.UUID]int{}
	lock := sync.Mutex{}

	return &Storage{
		Accounts:       accounts,
		AccountIndices: accountIndices,
		Transactions:   transactions,
		Mutex:          &lock,
	}
}
