package getavailablefunds

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/justinpjose/cushon-assignment/internal/handlers"
)

const (
	apiRspInvalidCustAccNo = "invalid customer account number given"
)

func (h handler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	h.log = h.log.CorrelationID()
	h.log.Infof("starting get available funds handler")

	params := httprouter.ParamsFromContext(r.Context())
	customerAccountNoStr := params.ByName("accountNo")
	customerAccountNo, err := strconv.Atoi(customerAccountNoStr)
	if err != nil {
		h.log.Warnf("failed to convert customerAccountNo to int: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		handlers.WriteRsp(w, h.log, apiRspInvalidCustAccNo)
		return
	}

	h.log.Field("customerAccountNo", customerAccountNo)

	availableFunds, err := h.db.GetAvailableFunds(ctx, customerAccountNo)
	if err != nil {
		h.log.Errorf("failed to get available funds from db: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(availableFunds)
	if err != nil {
		h.log.Errorf("failed to marshal available funds: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	handlers.WriteRsp(w, h.log, b)

	h.log.Infof("finished get available funds handler")
}
