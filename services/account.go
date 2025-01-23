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

func CreateAccountService(storage *storage.Storage) *AccountService {
	return &AccountService{
		Storage: storage,
	}
}

func (service *AccountService) Create(request requests.AccountRequest) (models.Account, error) {
	newUUID := uuid.New()
	account := models.Account{
		ID:      newUUID,
		Owner:   request.Owner,
		Balance: request.InitialBalance,
	}

	service.Storage.Mutex.Lock()
	defer service.Storage.Mutex.Unlock()

	index := len(service.Storage.Accounts)
	service.Storage.Accounts = append(service.Storage.Accounts, account)
	service.Storage.AccountIndices[newUUID] = index
	return account, nil
}

func (service *AccountService) ReadOne(id string) (models.Account, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return models.Account{}, utils.ErrInvalidUUID
	}

	service.Storage.Mutex.Lock()
	defer service.Storage.Mutex.Unlock()

	index, ok := service.Storage.AccountIndices[parsedUUID]
	if !ok {
		return models.Account{}, utils.ErrAccountNotFound
	}

	return service.Storage.Accounts[index], nil
}

func (service *AccountService) ReadAll() ([]models.Account, error) {

	service.Storage.Mutex.Lock()
	defer service.Storage.Mutex.Unlock()

	return service.Storage.Accounts, nil
}
