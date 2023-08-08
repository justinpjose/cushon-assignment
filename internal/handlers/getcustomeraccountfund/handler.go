package getcustomeraccountfund

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/justinpjose/cushon-assignment/internal/handlers"
)

const (
	apiRspInvalidCustAccFundID = "invalid customer accounts fund id given"
)

func (h handler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	h.log = h.log.CorrelationID()
	h.log.Infof("starting get customer accounts fund handler")

	params := httprouter.ParamsFromContext(r.Context())
	customerAccountsFundIDStr := params.ByName("id")
	customerAccountsFundID, err := strconv.Atoi(customerAccountsFundIDStr)
	if err != nil {
		h.log.Warnf("failed to convert customerAccountsFundID to int: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		handlers.WriteRsp(w, h.log, apiRspInvalidCustAccFundID)
		return
	}

	h.log.Field("customerAccountsFundID", customerAccountsFundID)

	customerAccountsFund, found, err := h.db.GetByID(ctx, customerAccountsFundID)
	if err != nil {
		h.log.Errorf("failed to get available funds from db: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !found {
		h.log.Infof("customer account fund not found with id %d", customerAccountsFundID)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	b, err := json.Marshal(customerAccountsFund)
	if err != nil {
		h.log.Errorf("failed to marshal customerAccountsFund: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	handlers.WriteRsp(w, h.log, b)

	h.log.Infof("finished get customer accounts fund handler")
}
