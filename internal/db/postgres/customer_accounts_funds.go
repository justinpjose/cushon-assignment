package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/justinpjose/cushon-assignment/internal/models"
)

func (p postgresDB) GetByID(ctx context.Context, id int) (models.CustomerAccountsFundDB, bool, error) {
	row := p.db.QueryRowContext(ctx, getCustomerAccountsFundByID, id)

	var customerAccountFund models.CustomerAccountsFundDB
	err := row.Scan(
		&customerAccountFund.ID,
		&customerAccountFund.CustomerAccountNo,
		&customerAccountFund.FundID,
		&customerAccountFund.TotalAmount,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return customerAccountFund, false, nil
		}

		return customerAccountFund, false, fmt.Errorf("failed to scan customerAccountFund from db: %v", err)
	}

	return customerAccountFund, true, nil
}

func (p postgresDB) GetTotalAmount(ctx context.Context, customerAccountFundID int) (int, error) {
	row := p.db.QueryRowContext(ctx, getTotalAmount, customerAccountFundID)

	var totalAmount int
	err := row.Scan(&totalAmount)
	if err != nil {
		return totalAmount, fmt.Errorf("failed to scan totalAmount from db: %v", err)
	}

	return totalAmount, nil
}
