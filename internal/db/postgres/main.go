package postgres

import (
	"database/sql"
	"fmt"

	"github.com/justinpjose/cushon-assignment/configs"
	internalDB "github.com/justinpjose/cushon-assignment/internal/db"
	"github.com/justinpjose/cushon-assignment/internal/logging"

	_ "github.com/lib/pq" // nolint - required for postgres
)

const (
	postgresDriverName = "postgres"
)

type postgresDB struct {
	db  *sql.DB
	log logging.Logger
}

// New creates a new postgres db connection based on the configurations given
func New(dbCfg configs.DBConfig, log logging.Logger) (internalDB.DB, error) {
	db, err := getDB(dbCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %v", err)
	}

	return postgresDB{
		db:  db,
		log: log,
	}, nil
}

// getDB - opens a PostgreSQL database and returns SQL db
func getDB(dbCfg configs.DBConfig) (*sql.DB, error) {
	dbInfo := fmt.Sprintf(
		"sslmode=disable host=%s port=%d user=%s password=%s dbname=%s",
		dbCfg.Host, dbCfg.Port, dbCfg.Username, dbCfg.Password, dbCfg.Name)

	db, err := sql.Open(postgresDriverName, dbInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping to db: %w", err)
	}

	return db, nil
}
