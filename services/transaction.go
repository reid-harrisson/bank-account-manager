package services

import (
	"bank-account-manager/models"
	"bank-account-manager/requests"
	"bank-account-manager/storage"
	"bank-account-manager/utils"
	"time"

	"github.com/google/uuid"
)

type TransactionService struct {
	Storage *storage.Storage
}

func CreateTransactionService(storage *storage.Storage) *TransactionService {
	return &TransactionService{
		Storage: storage,
	}
}

// Create handles creation of a new transaction for an account
func (service *TransactionService) Create(accountId string, request requests.TransactionRequest) (models.Transaction, error) {
	// Validate and parse the account UUID
	parsedAccountUUID, err := uuid.Parse(accountId)
	if err != nil {
		return models.Transaction{}, utils.ErrInvalidUUID
	}

	// Validate and parse the transaction type
	parsedType, err := utils.ParseTransactionType(request.Type)
	if err != nil {
		return models.Transaction{}, utils.ErrInvalidTxType
	}

	// Lock storage for thread safety
	service.Storage.Mutex.Lock()
	defer service.Storage.Mutex.Unlock()

	// Find the account in storage
	accountIndex, err := service.Storage.FindAccount(parsedAccountUUID)
	if err != nil {
		return models.Transaction{}, err
	}

	// Get account and check balance for withdrawals
	account := service.Storage.Accounts[accountIndex]
	if parsedType == utils.Withdrawal && account.Balance < request.Amount {
		return models.Transaction{}, utils.ErrInsufficientFunds
	}

	// Update account balance based on transaction type
	if parsedType == utils.Deposit {
		account.Balance += request.Amount
	} else if parsedType == utils.Withdrawal {
		account.Balance -= request.Amount
	}

	// Update storage with new account balance
	service.Storage.Accounts[accountIndex] = account

	// Create new transaction with unique ID
	newUUID := uuid.New()
	transaction := models.Transaction{
		ID:        newUUID,
		AccountID: parsedAccountUUID,
		Type:      parsedType,
		Amount:    request.Amount,
		TimeStamp: time.Now(),
	}

	// Add transaction to storage
	service.Storage.Transactions = append(service.Storage.Transactions, transaction)

	return transaction, nil
}

// ReadByAccount retrieves all transactions for a specific account
func (service *TransactionService) ReadByAccount(accountId string) ([]models.Transaction, error) {
	// Validate and parse the account UUID
	parsedAccountUUID, err := uuid.Parse(accountId)
	if err != nil {
		return []models.Transaction{}, utils.ErrInvalidUUID
	}

	// Lock storage for thread safety
	service.Storage.Mutex.Lock()
	defer service.Storage.Mutex.Unlock()

	// Filter transactions for the specified account
	transactions := []models.Transaction{}
	for _, transaction := range service.Storage.Transactions {
		if transaction.AccountID == parsedAccountUUID {
			transactions = append(transactions, transaction)
		}
	}

	return transactions, nil
}

// Transfer handles money transfer between two accounts
func (services *TransactionService) Transfer(request requests.TransferRequest) error {
	// Create withdrawal transaction from source account
	_, err := services.Create(request.FromAccountID, requests.TransactionRequest{
		Type:   utils.Withdrawal.String(),
		Amount: request.Amount,
	})

	if err != nil {
		return err
	}

	// Create deposit transaction to destination account
	_, err = services.Create(request.ToAccountID, requests.TransactionRequest{
		Type:   utils.Deposit.String(),
		Amount: request.Amount,
	})

	// If deposit fails, rollback the withdrawal by removing the last transaction
	if err != nil {
		services.Storage.Mutex.Lock()
		services.Storage.Transactions = services.Storage.Transactions[:len(services.Storage.Transactions)-1]
		services.Storage.Mutex.Unlock()
		return err
	}
	return nil
}
