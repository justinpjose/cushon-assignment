package main

import (
	internalDB "github.com/justinpjose/cushon-assignment/internal/db"
	internalHandler "github.com/justinpjose/cushon-assignment/internal/handlers"
	"github.com/justinpjose/cushon-assignment/internal/handlers/createtransaction"
	"github.com/justinpjose/cushon-assignment/internal/handlers/getavailablefunds"
	"github.com/justinpjose/cushon-assignment/internal/handlers/getcustomeraccountfund"
	"github.com/justinpjose/cushon-assignment/internal/logging"
)

type handlers struct {
	createTransactionHandler      internalHandler.Handler
	getAvailableFundsHandler      internalHandler.Handler
	getCustomerAccountFundHandler internalHandler.Handler
}

func getHandlers(db internalDB.DB, log logging.Logger) handlers {
	return handlers{
		createTransactionHandler:      createtransaction.New(db, log),
		getAvailableFundsHandler:      getavailablefunds.New(db, log),
		getCustomerAccountFundHandler: getcustomeraccountfund.New(db, log),
	}
}
