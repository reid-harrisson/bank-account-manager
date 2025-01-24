package utils

const (
	// Common messages
	MsgInvalidRequestBody = "Invalid request body"
	MsgValidationFailed   = "Validation failed"
	MsgIDCannotBeEmpty    = "ID cannot be empty"
	MsgInvalidUUID        = "Invalid Account UUID"
	MsgAccountNotFound    = "Account not found"

	// Transaction specific messages
	MsgFailedCreateTx      = "Failed to create transaction"
	MsgInvalidTxType       = "Invalid transaction type"
	MsgInsufficientFunds   = "Insufficient funds"
	MsgFailedRetrieveTx    = "Failed to retrieve transactions"
	MsgTransferSuccess     = "Successfully transferred"
	MsgSameAccountTransfer = "From and To account IDs cannot be the same"

	// Account specific messages
	MsgFailedCreateAccount    = "Failed to create account"
	MsgFailedRetrieveAccount  = "Failed to retrieve account"
	MsgFailedRetrieveAccounts = "Failed to retrieve accounts"
)
