package responses

import (
	"bank-account-manager/models"

	"github.com/gofiber/fiber/v2"
)

type Account struct {
	ID      string  `json:"id"`
	Owner   string  `json:"owner"`
	Balance float64 `json:"balance"`
}

func AccountResponse(ctx *fiber.Ctx, status int, account models.Account) error {
	return Response(ctx, status, Account{
		ID:      account.ID.String(),
		Owner:   account.Owner,
		Balance: account.Balance,
	})
}
