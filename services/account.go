package services

import (
	"bank-account-manager/models"
	"bank-account-manager/requests"
	"bank-account-manager/server"
	"bank-account-manager/storage"
	"bank-account-manager/utils"

	"github.com/google/uuid"
)

type AccountService struct {
	Storage *storage.Storage
}

func CreateAccountService(server *server.Server) *AccountService {
	return &AccountService{
		Storage: server.Storage,
	}
}

func (service *AccountService) Create(request requests.AccountRequest) (models.Account, error) {
	newUUID := uuid.New()
	account := models.Account{
		ID:      newUUID,
		Owner:   request.Owner,
		Balance: request.InitialBalance,
	}

	index := len(service.Storage.Accounts)
	service.Storage.Mutex.Lock()
	service.Storage.Accounts = append(service.Storage.Accounts, account)
	service.Storage.AccountIndices[newUUID] = index
	service.Storage.Mutex.Unlock()
	return account, nil
}

func (service *AccountService) ReadOne(id string) (models.Account, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return models.Account{}, utils.ErrInvalidUUID
	}

	index, ok := service.Storage.AccountIndices[parsedUUID]
	if !ok {
		return models.Account{}, utils.ErrAccountNotFound
	}

	return service.Storage.Accounts[index], nil
}

func (service *AccountService) ReadAll() ([]models.Account, error) {
	return service.Storage.Accounts, nil
}
