package models

// CreateTransactionReq is the model for the request body when making a create transaction request
type CreateTransactionReq struct {
	CustomerAccountFundID int `json:"customer_accounts_funds_id"`
	Amount                int `json:"amount"`
}
