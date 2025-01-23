package requests

import validation "github.com/go-ozzo/ozzo-validation"

type AccountRequest struct {
	Owner          string  `json:"owner"`
	InitialBalance float64 `json:"inital_balance"`
}

func (request AccountRequest) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Owner, validation.Required),
		validation.Field(&request.InitialBalance, validation.Required),
	)
}
