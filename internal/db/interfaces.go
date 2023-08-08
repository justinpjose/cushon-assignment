//go:generate mockgen -source=interfaces.go -destination=mocks/db.go -package=mocks

package db

import (
	"context"

	"github.com/justinpjose/cushon-assignment/internal/models"
)

type DB interface {
	TransactionsDB
	CustomerAccountsDB
	CustomerAccountsFundsDB
}

type TransactionsDB interface {
	// CreateTransaction creates a transaction record in the transaction table
	CreateTransaction(ctx context.Context, req models.CreateTransactionReq, totalAmount int) error
}

type CustomerAccountsDB interface {
	// GetAvailableFunds gets the available funds which the customer can invest into from their customer account
	GetAvailableFunds(ctx context.Context, customerAccountNo int) ([]models.AvailableFund, error)
}

type CustomerAccountsFundsDB interface {
	// GetByID gets information about the specific fund which the customer has invested in via their customer account
	GetByID(ctx context.Context, id int) (models.CustomerAccountsFundDB, bool, error)

	// GetAvailableFunds gets total amount invested in the specific fund in their customer account
	GetTotalAmount(ctx context.Context, customerAccountFundID int) (int, error)
}
