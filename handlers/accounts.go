package handlers

import (
	"bank-account-manager/requests"
	"bank-account-manager/responses"
	"bank-account-manager/server"
	"bank-account-manager/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AccountHandler struct {
	AccountService *services.AccountService
}

func CreateAccountHandler(server *server.Server) *AccountHandler {
	return &AccountHandler{
		AccountService: services.CreateAccountService(server),
	}
}

// CreateAccount godoc
// @Summary Create a new bank account
// @Description Creates a new bank account with the provided details
// @Tags Accounts
// @Accept json
// @Produce json
// @Param account body requests.AccountRequest true "Account details"
// @Success 201 {object} responses.Account
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/v1/accounts [post]
func (handler *AccountHandler) Create(context *fiber.Ctx) error {
	request := requests.AccountRequest{}
	if err := context.BodyParser(&request); err != nil {
		return responses.ErrorResponse(context, http.StatusBadRequest, "Invalid request body", err)
	}

	if err := request.Validate(); err != nil {
		return responses.ErrorResponse(context, http.StatusBadRequest, "Validation failed", err)
	}

	account, err := handler.AccountService.Create(request)
	if err != nil {
		return responses.ErrorResponse(context, http.StatusInternalServerError, "Failed to create account", err)
	}

	return responses.AccountResponse(context, http.StatusCreated, account)
}

// ReadAccount godoc
// @Summary Get a bank account by ID
// @Description Retrieves a bank account's details by its ID
// @Tags Accounts
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {object} responses.Account
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/v1/accounts/{id} [get]
func (handler *AccountHandler) ReadOne(context *fiber.Ctx) error {
	id := context.Params("id")
	if id == "" {
		return responses.ErrorResponse(context, http.StatusBadRequest, "Invalid account ID: ID cannot be empty", nil)
	}

	account, err := handler.AccountService.ReadOne(id)
	if err != nil {
		if err == services.ErrAccountNotFound {
			return responses.ErrorResponse(context, http.StatusNotFound, "Account not found", err)
		}
		return responses.ErrorResponse(context, http.StatusInternalServerError, "Failed to retrieve account", err)
	}

	return responses.AccountResponse(context, http.StatusOK, account)
}

// ReadAccounts godoc
// @Summary Get all bank accounts
// @Description Retrieves all bank accounts' details
// @Tags Accounts
// @Accept json
// @Produce json
// @Success 200 {array} responses.Account
// @Failure 500 {object} responses.Error
// @Router /api/v1/accounts [get]
func (handler *AccountHandler) ReadAll(context *fiber.Ctx) error {
	accounts, err := handler.AccountService.ReadAll()
	if err != nil {
		return responses.ErrorResponse(context, http.StatusInternalServerError, "Failed to retrieve accounts", err)
	}

	return responses.AccountResponses(context, http.StatusOK, accounts)
}
