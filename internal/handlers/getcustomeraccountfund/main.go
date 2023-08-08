package getcustomeraccountfund

import (
	"github.com/justinpjose/cushon-assignment/internal/db"
	"github.com/justinpjose/cushon-assignment/internal/handlers"
	"github.com/justinpjose/cushon-assignment/internal/logging"
)

type handler struct {
	db  db.CustomerAccountsFundsDB
	log logging.Logger
}

// New creates a 'get customer account fund' handler
func New(d db.CustomerAccountsFundsDB, l logging.Logger) handlers.Handler {
	return handler{
		db:  d,
		log: l,
	}
}
