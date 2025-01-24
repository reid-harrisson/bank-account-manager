package services

import (
	"bank-account-manager/models"
	"bank-account-manager/requests"
	"bank-account-manager/storage"
	"bank-account-manager/utils"

	"github.com/google/uuid"
)

type AccountService struct {
	Storage *storage.Storage
}

// CreateAccountService initializes a new AccountService with the provided storage
func CreateAccountService(storage *storage.Storage) *AccountService {
	return &AccountService{
		Storage: storage,
	}
}

// Create generates a new account based on the provided request
func (service *AccountService) Create(request requests.AccountRequest) (models.Account, error) {
	// Create a new Account instance with the request data
	newUUID := uuid.New()
	account := models.Account{
		ID:      newUUID,
		Owner:   request.Owner,
		Balance: request.InitialBalance,
	}

	// Lock mutex to ensure thread-safe operations
	service.Storage.Mutex.Lock()
	defer service.Storage.Mutex.Unlock()

	// Add the new account to storage
	service.Storage.Accounts = append(service.Storage.Accounts, account)
	return account, nil
}

// ReadOne retrieves a single account by its ID
func (service *AccountService) ReadOne(id string) (models.Account, error) {
	// Convert string ID to UUID type
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return models.Account{}, utils.ErrInvalidUUID
	}

	// Lock mutex to ensure thread-safe operations
	service.Storage.Mutex.Lock()
	defer service.Storage.Mutex.Unlock()

	// Find the account index in storage
	index, err := service.Storage.FindAccount(parsedUUID)
	if err != nil {
		return models.Account{}, err
	}

	// Return the found account
	return service.Storage.Accounts[index], nil
}

// ReadAll retrieves all accounts from storage
func (service *AccountService) ReadAll() ([]models.Account, error) {
	// Lock mutex to ensure thread-safe operations
	service.Storage.Mutex.Lock()
	defer service.Storage.Mutex.Unlock()

	// Return all accounts from storage
	return service.Storage.Accounts, nil
}
