package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type TransactionRequest struct {
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

func (request TransactionRequest) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Type, validation.Required),
		validation.Field(&request.Amount, validation.Required),
	)
}
