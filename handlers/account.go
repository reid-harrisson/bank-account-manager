// Package handlers contains HTTP request handlers for the bank account manager
package handlers

// Import necessary packages for handling HTTP requests, responses, and services
import (
	"bank-account-manager/requests"
	"bank-account-manager/responses"
	"bank-account-manager/server"
	"bank-account-manager/services"
	"net/http"

	"bank-account-manager/utils"

	"github.com/gofiber/fiber/v2"
)

type AccountHandler struct {
	AccountService *services.AccountService
}

func CreateAccountHandler(server *server.Server) *AccountHandler {
	return &AccountHandler{
		AccountService: services.CreateAccountService(server.Storage),
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
// @Router /accounts [post]
func (handler *AccountHandler) Create(context *fiber.Ctx) error {

	// Parse request body into AccountRequest struct
	request := requests.AccountRequest{}
	if err := context.BodyParser(&request); err != nil {
		return responses.ErrorResponse(context, http.StatusBadRequest, utils.MsgInvalidRequestBody)
	}

	// Validate the request data
	if err := request.Validate(); err != nil {
		return responses.ErrorResponse(context, http.StatusBadRequest, utils.MsgValidationFailed)
	}

	// Attempt to create account using service layer
	account, err := handler.AccountService.Create(request)
	if err != nil {
		return responses.ErrorResponse(context, http.StatusInternalServerError, utils.MsgFailedCreateAccount)
	}

	// Return successful response with created account details
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
// @Router /accounts/{id} [get]
func (handler *AccountHandler) ReadOne(context *fiber.Ctx) error {
	// Extract account ID from request parameters
	id := context.Params("id")
	if id == "" {
		return responses.ErrorResponse(context, http.StatusBadRequest, utils.MsgIDCannotBeEmpty)
	}

	// Attempt to retrieve account using service layer
	account, err := handler.AccountService.ReadOne(id)
	if err != nil {
		// Handle different error cases with appropriate HTTP status codes
		switch err {
		case utils.ErrInvalidUUID:
			return responses.ErrorResponse(context, http.StatusBadRequest, utils.MsgInvalidUUID)
		case utils.ErrAccountNotFound:
			return responses.ErrorResponse(context, http.StatusNotFound, utils.MsgAccountNotFound)
		}
		return responses.ErrorResponse(context, http.StatusInternalServerError, utils.MsgFailedRetrieveAccount)
	}

	// Return successful response with account details
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
// @Router /accounts [get]
func (handler *AccountHandler) ReadAll(context *fiber.Ctx) error {
	// Attempt to retrieve all accounts using service layer
	accounts, err := handler.AccountService.ReadAll()
	if err != nil {
		return responses.ErrorResponse(context, http.StatusInternalServerError, utils.MsgFailedRetrieveAccounts)
	}

	// Return successful response with all accounts
	return responses.AccountResponses(context, http.StatusOK, accounts)
}
