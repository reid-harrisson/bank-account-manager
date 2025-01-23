package handlers

import (
	"bank-account-manager/requests"
	"bank-account-manager/responses"
	"bank-account-manager/server"
	"bank-account-manager/services"
	"net/http"
	"regexp"

	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	TransactionService *services.TransactionService
}

func CreateTransactionHandler(server *server.Server) *TransactionHandler {
	return &TransactionHandler{
		TransactionService: services.CreateTransactionService(server),
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
// @Router /api/v1/accounts/{id}/transactions [post]
func (handler *TransactionHandler) Create(context *fiber.Ctx) error {
	accountID := context.Params("id")
	if accountID == "" {
		return responses.ErrorResponse(context, http.StatusBadRequest, "Invalid account ID: ID cannot be empty", nil)
	}

	if !isValidUUID(accountID) {
		return responses.ErrorResponse(context, http.StatusBadRequest, "Invalid account ID format", nil)
	}

	request := requests.TransactionRequest{}
	if err := context.BodyParser(&request); err != nil {
		return responses.ErrorResponse(context, http.StatusBadRequest, "Invalid request body", err)
	}

	if err := request.Validate(); err != nil {
		return responses.ErrorResponse(context, http.StatusBadRequest, "Validation failed", err)
	}

	transaction, err := handler.TransactionService.Create(accountID, request)
	if err != nil {
		return responses.ErrorResponse(context, http.StatusInternalServerError, "Failed to create transaction", err)
	}

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
// @Router /api/v1/accounts/{id}/transactions [get]
func (handler *TransactionHandler) ReadByAccount(context *fiber.Ctx) error {
	accountID := context.Params("id")
	if accountID == "" {
		return responses.ErrorResponse(context, http.StatusBadRequest, "Invalid account ID: ID cannot be empty", nil)
	}

	if !isValidUUID(accountID) {
		return responses.ErrorResponse(context, http.StatusBadRequest, "Invalid account ID format", nil)
	}

	transactions, err := handler.TransactionService.ReadByAccount(accountID)
	if err != nil {
		return responses.ErrorResponse(context, http.StatusInternalServerError, "Failed to retrieve transactions", err)
	}

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
// @Router /api/v1/transfer [post]
func (handler *TransactionHandler) Transfer(context *fiber.Ctx) error {
	request := requests.TransferRequest{}
	if err := context.BodyParser(&request); err != nil {
		return responses.ErrorResponse(context, http.StatusBadRequest, "Invalid request body", err)
	}

	if err := request.Validate(); err != nil {
		return responses.ErrorResponse(context, http.StatusBadRequest, "Validation failed", err)
	}

	err := handler.TransactionService.Transfer(request)
	if err != nil {
		return responses.ErrorResponse(context, http.StatusInternalServerError, "Failed to create transaction", err)
	}

	return responses.MessageResponse(context, http.StatusCreated, "Successfully transferred")
}

// Helper function to validate UUID format
func isValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}
