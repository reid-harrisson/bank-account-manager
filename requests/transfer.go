package requests

type TransferRequest struct {
	FromAcountID string  `json:"from_acount_id"`
	ToAccountID  string  `json:"to_account_id"`
	Amount       float64 `json:"amount"`
}
