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

	account, ok := (*service.Storage.Accounts)[parsedAccountUUID]
	if !ok {
		return models.Transaction{}, utils.ErrAccountNotFound
	}

	if parsedType == utils.Withdrawal && account.Balance < request.Amount {
		return models.Transaction{}, utils.ErrInsufficientFunds
	}

	if parsedType == utils.Deposit {
		account.Balance += request.Amount
	} else if parsedType == utils.Withdrawal {
		account.Balance -= request.Amount
	}

	(*service.Storage.Accounts)[parsedAccountUUID] = account

	transaction := models.Transaction{
		ID:        newUUID,
		AccountID: parsedAccountUUID,
		Type:      parsedType,
		Amount:    request.Amount,
		TimeStamp: time.Now(),
	}

	(*service.Storage.Transactions)[newUUID] = transaction

	return transaction, nil
}

func (service *TransactionService) ReadByAccount(accountId string) ([]models.Transaction, error) {
	transactions := []models.Transaction{}

	parsedAccountUUID, err := uuid.Parse(accountId)
	if err != nil {
		return transactions, utils.ErrInvalidUUID
	}

	_, ok := (*service.Storage.Accounts)[parsedAccountUUID]
	if !ok {
		return transactions, utils.ErrAccountNotFound
	}

	for _, transaction := range *service.Storage.Transactions {
		if transaction.AccountID == parsedAccountUUID {
			transactions = append(transactions, transaction)
		}
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
