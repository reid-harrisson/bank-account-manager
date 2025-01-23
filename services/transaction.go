package services

import (
	"bank-account-manager/models"
	"bank-account-manager/requests"
	"bank-account-manager/server"
	"bank-account-manager/storage"
	"bank-account-manager/utils"
	"time"

	"github.com/google/uuid"
)

type TransactionService struct {
	Storage *storage.Storage
}

func CreateTransactionService(server *server.Server) *TransactionService {
	return &TransactionService{
		Storage: server.Storage,
	}
}

func (service *TransactionService) Create(accountId string, request requests.TransactionRequest) (models.Transaction, error) {
	newUUID := uuid.New()

	parsedAccountUUID, err := uuid.Parse(accountId)
	if err != nil {
		return models.Transaction{}, utils.ErrInvalidUUID
	}

	parsedType, err := utils.ParseTransactionType(request.Type)
	if err != nil {
		return models.Transaction{}, utils.ErrInvalidTransactionType
	}

	accountIndex, ok := service.Storage.AccountIndices[parsedAccountUUID]
	if !ok {
		return models.Transaction{}, utils.ErrAccountNotFound
	}

	account := service.Storage.Accounts[accountIndex]

	if parsedType == utils.Withdrawal && account.Balance < request.Amount {
		return models.Transaction{}, utils.ErrInsufficientFunds
	}

	if parsedType == utils.Deposit {
		account.Balance += request.Amount
	} else if parsedType == utils.Withdrawal {
		account.Balance -= request.Amount
	}

	service.Storage.Accounts[accountIndex] = account

	transaction := models.Transaction{
		ID:        newUUID,
		AccountID: parsedAccountUUID,
		Type:      parsedType,
		Amount:    request.Amount,
		TimeStamp: time.Now(),
	}

	_, ok = service.Storage.Transactions[parsedAccountUUID]
	if !ok {
		service.Storage.Transactions[parsedAccountUUID] = []models.Transaction{}
	}

	service.Storage.Transactions[parsedAccountUUID] = append(service.Storage.Transactions[parsedAccountUUID], transaction)

	return transaction, nil
}

func (service *TransactionService) ReadByAccount(accountId string) ([]models.Transaction, error) {
	parsedAccountUUID, err := uuid.Parse(accountId)
	if err != nil {
		return []models.Transaction{}, utils.ErrInvalidUUID
	}

	_, ok := service.Storage.AccountIndices[parsedAccountUUID]
	if !ok {
		return []models.Transaction{}, utils.ErrAccountNotFound
	}

	transactions, ok := service.Storage.Transactions[parsedAccountUUID]
	if !ok {
		return []models.Transaction{}, nil
	}

	return transactions, nil
}

func (services *TransactionService) Transfer(request requests.TransferRequest) error {
	_, err := services.Create(request.FromAcountID, requests.TransactionRequest{
		Type:   utils.Withdrawal.String(),
		Amount: request.Amount,
	})

	if err != nil {
		return err
	}

	_, err = services.Create(request.ToAccountID, requests.TransactionRequest{
		Type:   utils.Deposit.String(),
		Amount: request.Amount,
	})

	if err != nil {
		return err
	}

	return nil
}
