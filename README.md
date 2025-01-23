# Bank Account Manager

## Overview

The Bank Account Manager is a simple RESTful API built in Go for managing bank accounts and their transactions. This API allows users to create accounts, perform transactions, and retrieve account details.

## API Operations

### 1. Create a New Account

- **Endpoint:** `POST /accounts`
- **Description:** Create a new bank account with an initial balance.
- **Request Body:**
  ```json
  {
    "owner": "Account Holder Name",
    "initial_balance": 100.0
  }
  ```

### 2. Retrieve Account Details

- **Endpoint:** `GET /accounts/{id}`
- **Description:** Retrieve details of a specific account by ID.

### 3. List All Accounts

- **Endpoint:** `GET /accounts`
- **Description:** Retrieve a list of all bank accounts.

### 4. Create a Transaction

- **Endpoint:** `POST /accounts/{id}/transactions`
- **Description:** Create a deposit or withdrawal transaction for a specific account.
- **Request Body:**
  ```json
  {
    "type": "deposit", // or "withdrawal"
    "amount": 50.0
  }
  ```

### 5. Retrieve Transactions for an Account

- **Endpoint:** `GET /accounts/{id}/transactions`
- **Description:** Retrieve all transactions associated with a specific account.

### 6. Transfer Between Accounts

- **Endpoint:** `POST /transfer`
- **Description:** Transfer funds from one account to another.
- **Request Body:**
  ```json
  {
    "from_account_id": "123",
    "to_account_id": "456",
    "amount": 30.0
  }
  ```

## Requirements

- **HTTP Methods:** Use appropriate HTTP methods (GET, POST).
- **Data Format:** JSON for request and response bodies.
- **Concurrency:** Handle concurrent transactions safely to maintain data integrity.
- **Error Handling:** Gracefully handle errors such as insufficient funds, invalid account IDs, and invalid transaction types.
- **Code Quality:** Write clean, well-structured, and maintainable code.
- **Persistence:** In-memory storage is sufficient; a database is not required.

## Instructions to Run and Test the Application

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/yourusername/bank-account-manager.git
   cd bank-account-manager
   ```

2. **Install Dependencies:**
   Ensure you have Go installed. Run:

   ```bash
   go mod tidy
   ```

3. **Run the Application:**

   ```bash
   go run main.go
   ```

4. **Test the API:**
   After running the application, navigate to `http://localhost:8000` in your web browser. This will redirect you to the Swagger documentation, where you can test the APIs.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
