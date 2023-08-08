package models

// CustomerAccountsFundDB is the model of each record stored in the customer_accounts_funds table in the db
type CustomerAccountsFundDB struct {
	ID                int `db:"id" json:"id"`
	CustomerAccountNo int `db:"customer_account_no" json:"customer_account_no"`
	FundID            int `db:"fund_id" json:"fund_id"`
	TotalAmount       int `db:"total_amount" json:"total_amount"`
}
