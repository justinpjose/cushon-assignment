package postgres

import (
	"context"
	"fmt"

	"github.com/justinpjose/cushon-assignment/internal/models"
)

func (p postgresDB) GetAvailableFunds(ctx context.Context, customerAccountNo int) ([]models.AvailableFund, error) {
	rows, err := p.db.QueryContext(ctx, getAvailableFunds, customerAccountNo)
	if err != nil {
		return nil, fmt.Errorf("failed to query getAvailableFunds in db: %v", err)
	}

	var availableFunds []models.AvailableFund
	for rows.Next() {
		var fund models.AvailableFund

		err = rows.Scan(&fund.ID, &fund.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row in db: %v", err)
		}

		availableFunds = append(availableFunds, fund)
	}

	err = rows.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close rows in db: %v", err)
	}

	return availableFunds, nil
}
