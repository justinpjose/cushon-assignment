package createtransaction

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/justinpjose/cushon-assignment/internal/handlers"
	"github.com/justinpjose/cushon-assignment/internal/models"
)

func (h handler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	h.log = h.log.CorrelationID()
	h.log.Infof("starting create transaction handler")

	var req models.CreateTransactionReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.log.Errorf("failed to decode req body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = validateRequest(req)
	if err != nil {
		h.log.Warnf("failed to validate request: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		handlers.WriteRsp(w, h.log, "invalid request body given")
		return
	}

	h.log.Field("customerAccountFundID", req.CustomerAccountFundID)

	totalAmount, err := h.db.GetTotalAmount(ctx, req.CustomerAccountFundID)
	if err != nil {
		h.log.Errorf("failed to get total amount for customer account fund: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	newTotalAmount := totalAmount + req.Amount

	err = h.db.CreateTransaction(ctx, req, newTotalAmount)
	if err != nil {
		h.log.Errorf("failed to create transaction in db: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	h.log.Infof("finished create transaction handler")
}

func validateRequest(req models.CreateTransactionReq) error {
	if req.CustomerAccountFundID <= 0 {
		return fmt.Errorf("invalid CustomerAccountFundID - provided %d", req.CustomerAccountFundID)
	}

	return nil
}
