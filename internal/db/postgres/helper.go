package postgres

import (
	"database/sql"
)

func (p postgresDB) rollback(tx *sql.Tx) {
	err := tx.Rollback()
	if err != nil {
		p.log.Errorf("failed to rollback db transaction: %v", err)
	}
}
