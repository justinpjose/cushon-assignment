package main

import (
	"fmt"
	"net/http"

	"github.com/justinpjose/cushon-assignment/internal/router"
)

func getRoutes(r router.Router, h handlers, version string) {
	if version == "v0" {
		getRoutesV0(r, h)
	}
}

func getRoutesV0(r router.Router, h handlers) {
	r.HandlerFunc(
		http.MethodGet,
		v0("/customer_accounts/:accountNo/available_funds"),
		h.getAvailableFundsHandler,
	)

	r.HandlerFunc(
		http.MethodGet,
		v0("/customer_accounts_funds/:id"),
		h.getCustomerAccountFundHandler,
	)

	r.HandlerFunc(
		http.MethodPost,
		v0("/transactions"),
		h.createTransactionHandler,
	)
}

func v0(path string) string {
	return fmt.Sprintf("/v0%s", path)
}
