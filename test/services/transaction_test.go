package test

import (
	"bank-account-manager/requests"
	"bank-account-manager/services"
	"bank-account-manager/storage"
	"bank-account-manager/utils"
	"testing"
)

func TestCreateTransaction(t *testing.T) {
	// Setup
	storage := storage.Create()
	accountService := services.CreateAccountService(storage)
	transactionService := services.CreateTransactionService(storage)

	// Create an account
	accountRequest := requests.AccountRequest{
		Owner:          "Alice",
		InitialBalance: 2000,
	}
	account, _ := accountService.Create(accountRequest)

	// Test data for transaction
	transactionRequest := requests.TransactionRequest{
		Type:   "deposit",
		Amount: 500,
	}

	// Create transaction
	transaction, err := transactionService.Create(account.ID.String(), transactionRequest)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Validate transaction
	if transaction.AccountID != account.ID {
		t.Errorf("Expected account ID %s, got %s", account.ID, transaction.AccountID)
	}
}

func TestTransfer(t *testing.T) {
	// Setup
	storage := storage.Create()
	accountService := services.CreateAccountService(storage)
	transactionService := services.CreateTransactionService(storage)

	// Create accounts
	fromAccount, _ := accountService.Create(requests.AccountRequest{Owner: "Bob", InitialBalance: 1000})
	toAccount, _ := accountService.Create(requests.AccountRequest{Owner: "Charlie", InitialBalance: 500})

	// Test transfer
	transferRequest := requests.TransferRequest{
		FromAccountID: fromAccount.ID.String(),
		ToAccountID:   toAccount.ID.String(),
		Amount:        300,
	}

	err := transactionService.Transfer(transferRequest)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	fromAccount, _ = accountService.ReadOne(fromAccount.ID.String())
	toAccount, _ = accountService.ReadOne(toAccount.ID.String())

	// Validate balances
	if fromAccount.Balance != 700 {
		t.Errorf("Expected from account balance 700, got %f", fromAccount.Balance)
	}
	if toAccount.Balance != 800 {
		t.Errorf("Expected to account balance 800, got %f", toAccount.Balance)
	}
}

func TestReadByAccount(t *testing.T) {
	// Setup
	storage := storage.Create()
	accountService := services.CreateAccountService(storage)
	transactionService := services.CreateTransactionService(storage)

	// Create an account
	accountRequest := requests.AccountRequest{
		Owner:          "Alice",
		InitialBalance: 2000,
	}
	account, _ := accountService.Create(accountRequest)

	// Create transactions
	transactionRequest1 := requests.TransactionRequest{
		Type:   "deposit",
		Amount: 500,
	}
	transactionService.Create(account.ID.String(), transactionRequest1)

	transactionRequest2 := requests.TransactionRequest{
		Type:   "withdrawal",
		Amount: 200,
	}
	transactionService.Create(account.ID.String(), transactionRequest2)

	// Read transactions by account
	transactions, err := transactionService.ReadByAccount(account.ID.String())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Validate transactions count
	if len(transactions) != 2 {
		t.Errorf("Expected 2 transactions, got %d", len(transactions))
	}
}

// Test for error case when account does not exist
func TestReadByAccount_NotFound(t *testing.T) {
	// Setup
	storage := storage.Create()
	transactionService := services.CreateTransactionService(storage)

	// Attempt to read transactions for a non-existent account
	_, err := transactionService.ReadByAccount("non-existent-account-id")
	if err != utils.ErrInvalidUUID {
		t.Fatalf("Expected error %v, got %v", utils.ErrInvalidUUID, err)
	}
}

// Test for error case when creating a transaction with an invalid account ID
func TestCreateTransaction_InvalidAccountID(t *testing.T) {
	// Setup
	storage := storage.Create()
	transactionService := services.CreateTransactionService(storage)

	// Attempt to create a transaction with an invalid account ID
	transactionRequest := requests.TransactionRequest{
		Type:   "deposit",
		Amount: 500,
	}
	_, err := transactionService.Create("invalid-account-id", transactionRequest)
	if err != utils.ErrInvalidUUID {
		t.Fatalf("Expected error %v, got %v", utils.ErrInvalidUUID, err)
	}
}

// Test for error case when creating a withdrawal transaction with insufficient funds
func TestCreateTransaction_InsufficientFunds(t *testing.T) {
	// Setup
	storage := storage.Create()
	accountService := services.CreateAccountService(storage)
	transactionService := services.CreateTransactionService(storage)

	// Create an account with a low balance
	accountRequest := requests.AccountRequest{
		Owner:          "Alice",
		InitialBalance: 100,
	}
	account, _ := accountService.Create(accountRequest)

	// Attempt to create a withdrawal transaction that exceeds the balance
	transactionRequest := requests.TransactionRequest{
		Type:   "withdrawal",
		Amount: 200,
	}
	_, err := transactionService.Create(account.ID.String(), transactionRequest)
	if err != utils.ErrInsufficientFunds {
		t.Fatalf("Expected error %v, got %v", utils.ErrInsufficientFunds, err)
	}
}

// ... additional tests for error cases ...
