//go:build unit

package createtransaction_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/justinpjose/cushon-assignment/internal/db/mocks"
	"github.com/justinpjose/cushon-assignment/internal/handlers/createtransaction"
	"github.com/justinpjose/cushon-assignment/internal/logging/zerolog"
	"github.com/justinpjose/cushon-assignment/internal/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var (
	unexpectedErr = errors.New("unexpected error")
)

func TestHandle(t *testing.T) {

	cases := []struct {
		name            string
		req             models.CreateTransactionReq
		expectedDBCalls func(db *mocks.MockDB, req models.CreateTransactionReq)
		rspStatusCode   int
	}{
		{
			name: "success",
			req: models.CreateTransactionReq{
				CustomerAccountFundID: 1,
				Amount:                25000,
			},
			expectedDBCalls: func(db *mocks.MockDB, req models.CreateTransactionReq) {
				currentTotalAmount := 0
				newTotalAmount := currentTotalAmount + req.Amount

				db.EXPECT().GetTotalAmount(gomock.Any(), req.CustomerAccountFundID).Return(currentTotalAmount, nil)
				db.EXPECT().CreateTransaction(gomock.Any(), req, newTotalAmount).Return(nil)
			},
			rspStatusCode: http.StatusCreated,
		},
		{
			name: "invalid request",
			req: models.CreateTransactionReq{
				CustomerAccountFundID: -100,
				Amount:                25000,
			},
			rspStatusCode: http.StatusBadRequest,
		},
		{
			name: "failed to get total amount",
			req: models.CreateTransactionReq{
				CustomerAccountFundID: 1,
				Amount:                25000,
			},
			expectedDBCalls: func(db *mocks.MockDB, req models.CreateTransactionReq) {
				db.EXPECT().GetTotalAmount(gomock.Any(), req.CustomerAccountFundID).Return(0, unexpectedErr)
			},
			rspStatusCode: http.StatusInternalServerError,
		},
		{
			name: "failed to create transaction",
			req: models.CreateTransactionReq{
				CustomerAccountFundID: 1,
				Amount:                25000,
			},
			expectedDBCalls: func(db *mocks.MockDB, req models.CreateTransactionReq) {
				currentTotalAmount := 0
				newTotalAmount := currentTotalAmount + req.Amount

				db.EXPECT().GetTotalAmount(gomock.Any(), req.CustomerAccountFundID).Return(currentTotalAmount, nil)
				db.EXPECT().CreateTransaction(gomock.Any(), req, newTotalAmount).Return(unexpectedErr)
			},
			rspStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			db := mocks.NewMockDB(ctrl)
			log := zerolog.NewMockLog()
			handler := createtransaction.New(db, log)

			if tc.expectedDBCalls != nil {
				tc.expectedDBCalls(db, tc.req)
			}

			w := httptest.NewRecorder()
			b := new(bytes.Buffer)
			json.NewEncoder(b).Encode(tc.req)
			req := httptest.NewRequest(http.MethodPost, "/transactions", b)

			handler.Handle(w, req)

			rsp := w.Result()
			assert.Equal(t, tc.rspStatusCode, rsp.StatusCode)

			ctrl.Finish()
		})
	}
}
