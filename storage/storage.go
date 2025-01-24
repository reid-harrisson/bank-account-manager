package storage

import (
	"bank-account-manager/models"
	"bank-account-manager/utils"
	"sync"

	"github.com/google/uuid"
)

// Storage represents an in-memory data store for accounts and transactions
// with thread-safe operations through mutex locking
type Storage struct {
	Accounts     []models.Account     // Slice containing all bank accounts
	Transactions []models.Transaction // Slice containing all transactions
	Mutex        *sync.Mutex          // Mutex for thread-safe operations
}

// Create initializes and returns a new Storage instance with empty
// accounts and transactions slices and a mutex lock
func Create() *Storage {
	accounts := []models.Account{}
	transactions := []models.Transaction{}
	lock := sync.Mutex{}

	return &Storage{
		Accounts:     accounts,
		Transactions: transactions,
		Mutex:        &lock,
	}
}

// FindAccount searches for an account by its UUID and returns its index
// in the Accounts slice. Returns -1 and ErrAccountNotFound if not found
func (storage *Storage) FindAccount(id uuid.UUID) (int, error) {
	for index, account := range storage.Accounts {
		if account.ID == id {
			return index, nil
		}
	}
	return -1, utils.ErrAccountNotFound
}
