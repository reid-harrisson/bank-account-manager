package handlers

import (
	"bank-account-manager/requests"
	"bank-account-manager/responses"
	"bank-account-manager/server"
	"bank-account-manager/services"
	"bank-account-manager/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	TransactionService *services.TransactionService
}

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
	accountID := context.Params("id")
	if accountID == "" {
		return responses.ErrorResponse(context, http.StatusBadRequest, "ID cannot be empty")
	}

	request := requests.TransactionRequest{}
	if err := context.BodyParser(&request); err != nil {
		return responses.ErrorResponse(context, http.StatusBadRequest, "Invalid request body")
	}

	if err := request.Validate(); err != nil {
		return responses.ErrorResponse(context, http.StatusBadRequest, "Validation failed")
	}

	transaction, err := handler.TransactionService.Create(accountID, request)
	if err != nil {
		if err == utils.ErrInvalidTransactionType {
			return responses.ErrorResponse(context, http.StatusBadRequest, "Invalid transaction type")
		} else if err == utils.ErrInsufficientFunds {
			return responses.ErrorResponse(context, http.StatusBadRequest, "Insufficient funds")
		} else if err == utils.ErrAccountNotFound {
			return responses.ErrorResponse(context, http.StatusNotFound, "Account not found")
		} else if err == utils.ErrInvalidUUID {
			return responses.ErrorResponse(context, http.StatusBadRequest, "Invalid Account UUID")
		}
		return responses.ErrorResponse(context, http.StatusInternalServerError, "Failed to create transaction")
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
// @Router /accounts/{id}/transactions [get]
func (handler *TransactionHandler) ReadByAccount(context *fiber.Ctx) error {
	accountID := context.Params("id")
	if accountID == "" {
		return responses.ErrorResponse(context, http.StatusBadRequest, "ID cannot be empty")
	}

	transactions, err := handler.TransactionService.ReadByAccount(accountID)
	if err != nil {
		if err == utils.ErrAccountNotFound {
			return responses.ErrorResponse(context, http.StatusNotFound, "Account not found")
		} else if err == utils.ErrInvalidUUID {
			return responses.ErrorResponse(context, http.StatusBadRequest, "Invalid Account UUID")
		}
		return responses.ErrorResponse(context, http.StatusInternalServerError, "Failed to retrieve transactions")
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
// @Router /transfer [post]
func (handler *TransactionHandler) Transfer(context *fiber.Ctx) error {
	request := requests.TransferRequest{}
	if err := context.BodyParser(&request); err != nil {
		return responses.ErrorResponse(context, http.StatusBadRequest, "Invalid request body")
	}

	if err := request.Validate(); err != nil {
		return responses.ErrorResponse(context, http.StatusBadRequest, "Validation failed")
	}

	if request.FromAcountID == request.ToAccountID {
		return responses.ErrorResponse(context, http.StatusBadRequest, "From and To account IDs cannot be the same")
	}

	err := handler.TransactionService.Transfer(request)
	if err != nil {
		if err == utils.ErrInvalidTransactionType {
			return responses.ErrorResponse(context, http.StatusBadRequest, "Invalid transaction type")
		} else if err == utils.ErrInsufficientFunds {
			return responses.ErrorResponse(context, http.StatusBadRequest, "Insufficient funds")
		} else if err == utils.ErrAccountNotFound {
			return responses.ErrorResponse(context, http.StatusNotFound, "Account not found")
		} else if err == utils.ErrInvalidUUID {
			return responses.ErrorResponse(context, http.StatusBadRequest, "Invalid Account UUID")
		}
		return responses.ErrorResponse(context, http.StatusInternalServerError, "Failed to create transaction")
	}

	return responses.MessageResponse(context, http.StatusCreated, "Successfully transferred")
}
