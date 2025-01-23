package test

import (
	"bank-account-manager/requests"
	"bank-account-manager/services"
	"bank-account-manager/storage"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	// Setup
	storage := storage.Create() // Assuming a function to initialize storage
	service := services.CreateAccountService(storage)

	// Test data
	request := requests.AccountRequest{
		Owner:          "John Doe",
		InitialBalance: 1000,
	}

	// Create account
	account, err := service.Create(request)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Validate account
	if account.Owner != request.Owner {
		t.Errorf("Expected owner %s, got %s", request.Owner, account.Owner)
	}
	if account.Balance != request.InitialBalance {
		t.Errorf("Expected balance %f, got %f", request.InitialBalance, account.Balance)
	}
}

func TestReadAccount(t *testing.T) {
	// Setup
	storage := storage.Create()
	service := services.CreateAccountService(storage)

	// Create an account
	request := requests.AccountRequest{
		Owner:          "Jane Doe",
		InitialBalance: 500,
	}
	account, _ := service.Create(request)

	// Read the account
	readAccount, err := service.ReadOne(account.ID.String())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Validate read account
	if readAccount.ID != account.ID {
		t.Errorf("Expected account ID %s, got %s", account.ID, readAccount.ID)
	}
}

func TestReadAllAccounts(t *testing.T) {
	// Setup
	storage := storage.Create()
	service := services.CreateAccountService(storage)

	// Create test accounts
	accounts := []requests.AccountRequest{
		{Owner: "Alice", InitialBalance: 1000},
		{Owner: "Bob", InitialBalance: 2000},
	}

	for _, req := range accounts {
		_, err := service.Create(req)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
	}

	// Read all accounts
	readAccounts, err := service.ReadAll()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Validate the number of accounts read
	if len(readAccounts) != len(accounts) {
		t.Errorf("Expected %d accounts, got %d", len(accounts), len(readAccounts))
	}
}

func TestReadAccountError(t *testing.T) {
	// Setup
	storage := storage.Create()
	service := services.CreateAccountService(storage)

	// Attempt to read a non-existent account
	_, err := service.ReadOne("non-existent-id")
	if err == nil {
		t.Fatalf("Expected an error for non-existent account, got none")
	}
}

// ... additional tests for ReadAll and error cases ...
