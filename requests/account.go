package requests

type AccountRequest struct {
	Owner          string  `json:"owner"`
	InitialBalance float64 `json:"inital_balance"`
}
