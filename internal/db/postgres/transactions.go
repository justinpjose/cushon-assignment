package postgres

import (
	"context"
	"fmt"

	"github.com/justinpjose/cushon-assignment/internal/models"
)

func (p postgresDB) CreateTransaction(ctx context.Context, req models.CreateTransactionReq, totalAmount int) error {
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin db transaction: %v", err)
	}

	// create transaction
	_, err = tx.ExecContext(ctx, createTransaction,
		req.CustomerAccountFundID,
		req.Amount,
	)
	if err != nil {
		p.rollback(tx)
		return fmt.Errorf("failed to exec createTransaction query: %v", err)
	}

	// update customer_accounts_fund
	_, err = tx.ExecContext(ctx, updateTotalAmountInFund,
		totalAmount,
		req.CustomerAccountFundID,
	)
	if err != nil {
		p.rollback(tx)
		return fmt.Errorf("failed to exec updateTotalAmountInFund query: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		p.rollback(tx)
		return fmt.Errorf("failed to commit db transaction: %v", err)
	}

	return nil
}
