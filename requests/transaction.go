package requests

import "bank-account-manager/utils"

type TransactionRequest struct {
	Type   utils.TransactionType `json:"type"`
	Amount float64               `json:"amount"`
}
