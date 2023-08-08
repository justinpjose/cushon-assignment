package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func getRouter(h handlers, version string) *httprouter.Router {
	router := httprouter.New()

	if version == "v0" {
		getRoutesV0(router, h)
	}

	return router
}

func getRoutesV0(router *httprouter.Router, h handlers) {
	router.HandlerFunc(
		http.MethodGet,
		v0("/customer_accounts/:accountNo/available_funds"),
		h.getAvailableFundsHandler.Handle,
	)

	router.HandlerFunc(
		http.MethodGet,
		v0("/customer_accounts_funds/:id"),
		h.getCustomerAccountFundHandler.Handle,
	)

	router.HandlerFunc(
		http.MethodPost,
		v0("/transactions"),
		h.createTransactionHandler.Handle,
	)
}

func v0(path string) string {
	return fmt.Sprintf("/v0%s", path)
}
