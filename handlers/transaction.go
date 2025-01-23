package handlers

import (
	"bank-account-manager/requests"
	"bank-account-manager/responses"
	"bank-account-manager/server"
	"bank-account-manager/services"
	"net/http"

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
