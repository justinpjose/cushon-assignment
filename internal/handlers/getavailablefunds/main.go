package getavailablefunds

import (
	"github.com/justinpjose/cushon-assignment/internal/db"
	"github.com/justinpjose/cushon-assignment/internal/handlers"
	"github.com/justinpjose/cushon-assignment/internal/logging"
)

type handler struct {
	db  db.CustomerAccountsDB
	log logging.Logger
}

// New creates a 'get available funds' handler
func New(d db.CustomerAccountsDB, l logging.Logger) handlers.Handler {
	return handler{
		db:  d,
		log: l,
	}
}
