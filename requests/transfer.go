package requests

import validation "github.com/go-ozzo/ozzo-validation"

type TransferRequest struct {
	FromAccountID string  `json:"from_acount_id"`
	ToAccountID   string  `json:"to_account_id"`
	Amount        float64 `json:"amount"`
}

func (request TransferRequest) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.FromAccountID, validation.Required),
		validation.Field(&request.ToAccountID, validation.Required),
		validation.Field(&request.Amount, validation.Required),
	)
}
