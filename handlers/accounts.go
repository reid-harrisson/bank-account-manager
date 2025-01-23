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
// @Success 201 {object} responses.Message
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

	err := handler.AccountService.Create(request)
	if err != nil {
		return responses.ErrorResponse(context, http.StatusInternalServerError, "Failed to create account", err)
	}

	return responses.MessageResponse(context, http.StatusCreated, "Account successfully created")
}
