package responses

import (
	"bank-account-manager/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Transaction struct {
	ID        string  `json:"id"`
	AccountID string  `json:"account_id"`
	Type      string  `json:"type"`
	Amount    float64 `json:"amount"`
	TimeStamp string  `json:"timestamp"`
}

func TransactionResponse(ctx *fiber.Ctx, status int, transaction models.Transaction) error {
	return Response(ctx, status, Transaction{
		ID:        transaction.ID.String(),
		AccountID: transaction.AccountID.String(),
		Type:      transaction.Type.String(),
		Amount:    transaction.Amount,
		TimeStamp: transaction.TimeStamp.Format(time.RFC3339),
	})
}

func TransactionResponses(ctx *fiber.Ctx, status int, transactions []models.Transaction) error {
	transactionResponses := []Transaction{}
	for _, account := range transactions {
		transactionResponses = append(transactionResponses, Transaction{
			ID:        account.ID.String(),
			AccountID: account.AccountID.String(),
			Type:      account.Type.String(),
			Amount:    account.Amount,
			TimeStamp: account.TimeStamp.Format(time.RFC3339),
		})
	}
	return Response(ctx, status, transactionResponses)
}
