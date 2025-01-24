// Package handlers contains HTTP request handlers for the bank account manager
package handlers

// Import necessary packages for handling transactions, responses, and services
import (
	"bank-account-manager/requests"
	"bank-account-manager/responses"
	"bank-account-manager/server"
	"bank-account-manager/services"
	"bank-account-manager/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// TransactionHandler struct holds the transaction service instance for handling transaction operations
type TransactionHandler struct {
	TransactionService *services.TransactionService
}

// CreateTransactionHandler initializes a new TransactionHandler with the provided server's storage
func CreateTransactionHandler(server *server.Server) *TransactionHandler {
	return &TransactionHandler{
		TransactionService: services.CreateTransactionService(server.Storage),
	}
}

// Create godoc
// @Summary Create a new transaction
// @Description Creates a new transaction for the specified bank account
// @Tags Transactions
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Param transaction body requests.TransactionRequest true "Transaction details"
// @Success 201 {object} responses.Transaction
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /accounts/{id}/transactions [post]
func (handler *TransactionHandler) Create(context *fiber.Ctx) error {
	// Extract account ID from request parameters
	accountID := context.Params("id")
	if accountID == "" {
		return responses.ErrorResponse(context, http.StatusBadRequest, utils.MsgIDCannotBeEmpty)
	}

	// Initialize and parse transaction request from request body
	request := requests.TransactionRequest{}
	if err := context.BodyParser(&request); err != nil {
		return responses.ErrorResponse(context, http.StatusBadRequest, utils.MsgInvalidRequestBody)
	}

	// Validate the transaction request data
	if err := request.Validate(); err != nil {
		return responses.ErrorResponse(context, http.StatusBadRequest, utils.MsgValidationFailed)
	}

	// Attempt to create transaction using service layer
	transaction, err := handler.TransactionService.Create(accountID, request)
	if err != nil {
		// Handle various transaction-specific errors with appropriate status codes
		switch err {
		case utils.ErrInvalidTxType:
			return responses.ErrorResponse(context, http.StatusBadRequest, utils.MsgInvalidTxType)
		case utils.ErrInsufficientFunds:
			return responses.ErrorResponse(context, http.StatusBadRequest, utils.MsgInsufficientFunds)
		case utils.ErrAccountNotFound:
			return responses.ErrorResponse(context, http.StatusNotFound, utils.MsgAccountNotFound)
		case utils.ErrInvalidUUID:
			return responses.ErrorResponse(context, http.StatusBadRequest, utils.MsgInvalidUUID)
		default:
			return responses.ErrorResponse(context, http.StatusInternalServerError, utils.MsgFailedCreateTx)
		}
	}

	// Return successful response with created transaction details
	return responses.TransactionResponse(context, http.StatusCreated, transaction)
}

// ReadByAccount godoc
// @Summary Get account transactions
// @Description Retrieves all transactions for the specified bank account
// @Tags Transactions
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {object} []responses.Transaction
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /accounts/{id}/transactions [get]
func (handler *TransactionHandler) ReadByAccount(context *fiber.Ctx) error {
	// Extract account ID from request parameters
	accountID := context.Params("id")
	if accountID == "" {
		return responses.ErrorResponse(context, http.StatusBadRequest, utils.MsgIDCannotBeEmpty)
	}

	// Attempt to retrieve transactions using service layer
	transactions, err := handler.TransactionService.ReadByAccount(accountID)
	if err != nil {
		// Handle various error cases with appropriate status codes
		switch err {
		case utils.ErrAccountNotFound:
			return responses.ErrorResponse(context, http.StatusNotFound, utils.MsgAccountNotFound)
		case utils.ErrInvalidUUID:
			return responses.ErrorResponse(context, http.StatusBadRequest, utils.MsgInvalidUUID)
		default:
			return responses.ErrorResponse(context, http.StatusInternalServerError, utils.MsgFailedRetrieveTx)
		}
	}

	// Return successful response with all transactions for the account
	return responses.TransactionResponses(context, http.StatusOK, transactions)
}

// Transfer godoc
// @Summary Transfer funds between accounts
// @Description Transfer funds from one account to another
// @Tags Transactions
// @Accept json
// @Produce json
// @Param transaction body requests.TransferRequest true "Transfer details"
// @Success 201 {object} responses.Message
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /transfer [post]
func (handler *TransactionHandler) Transfer(context *fiber.Ctx) error {
	// Initialize and parse transfer request from request body
	request := requests.TransferRequest{}
	if err := context.BodyParser(&request); err != nil {
		return responses.ErrorResponse(context, http.StatusBadRequest, utils.MsgInvalidRequestBody)
	}

	// Validate the transfer request data
	if err := request.Validate(); err != nil {
		return responses.ErrorResponse(context, http.StatusBadRequest, utils.MsgValidationFailed)
	}

	// Check if source and destination accounts are different
	if request.FromAccountID == request.ToAccountID {
		return responses.ErrorResponse(context, http.StatusBadRequest, utils.MsgSameAccountTransfer)
	}

	// Attempt to process transfer using service layer
	err := handler.TransactionService.Transfer(request)
	if err != nil {
		// Handle various transfer-specific errors with appropriate status codes
		switch err {
		case utils.ErrInvalidTxType:
			return responses.ErrorResponse(context, http.StatusBadRequest, utils.MsgInvalidTxType)
		case utils.ErrInsufficientFunds:
			return responses.ErrorResponse(context, http.StatusBadRequest, utils.MsgInsufficientFunds)
		case utils.ErrAccountNotFound:
			return responses.ErrorResponse(context, http.StatusNotFound, utils.MsgAccountNotFound)
		case utils.ErrInvalidUUID:
			return responses.ErrorResponse(context, http.StatusBadRequest, utils.MsgInvalidUUID)
		default:
			return responses.ErrorResponse(context, http.StatusInternalServerError, utils.MsgFailedCreateTx)
		}
	}

	// Return successful response with transfer confirmation message
	return responses.MessageResponse(context, http.StatusCreated, utils.MsgTransferSuccess)
}
