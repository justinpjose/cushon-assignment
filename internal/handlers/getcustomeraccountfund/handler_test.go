//go:build unit

package getcustomeraccountfund_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/justinpjose/cushon-assignment/internal/db/mocks"
	"github.com/justinpjose/cushon-assignment/internal/handlers/getcustomeraccountfund"
	"github.com/justinpjose/cushon-assignment/internal/logging/zerolog"
	"github.com/justinpjose/cushon-assignment/internal/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

const (
	invalidParam = "invalid"
)

var (
	unexpectedErr = errors.New("unexpected error")
)

func TestHandle(t *testing.T) {

	cases := []struct {
		name            string
		idStr           string
		expectedDBCalls func(db *mocks.MockCustomerAccountsFundsDB, id int, rsp models.CustomerAccountsFundDB)
		rspBody         models.CustomerAccountsFundDB
		rspStatusCode   int
	}{
		{
			name:  "success",
			idStr: "1",
			expectedDBCalls: func(db *mocks.MockCustomerAccountsFundsDB, id int, rsp models.CustomerAccountsFundDB) {
				db.EXPECT().GetByID(gomock.Any(), id).Return(rsp, true, nil)
			},
			rspBody: models.CustomerAccountsFundDB{
				ID:                1,
				CustomerAccountNo: 1,
				FundID:            1,
				TotalAmount:       25000,
			},
			rspStatusCode: http.StatusOK,
		},
		{
			name:          "invalid id given",
			idStr:         invalidParam,
			rspStatusCode: http.StatusBadRequest,
		},
		{
			name:  "failed to get customer account fund from db",
			idStr: "1",
			expectedDBCalls: func(db *mocks.MockCustomerAccountsFundsDB, id int, rsp models.CustomerAccountsFundDB) {
				db.EXPECT().GetByID(gomock.Any(), id).Return(rsp, false, unexpectedErr)
			},
			rspStatusCode: http.StatusInternalServerError,
		},
		{
			name:  "when customer account fund for given id does not exist",
			idStr: "1",
			expectedDBCalls: func(db *mocks.MockCustomerAccountsFundsDB, id int, rsp models.CustomerAccountsFundDB) {
				db.EXPECT().GetByID(gomock.Any(), id).Return(rsp, false, nil)
			},
			rspStatusCode: http.StatusNotFound,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			db := mocks.NewMockCustomerAccountsFundsDB(ctrl)
			log := zerolog.NewMockLog()
			handler := getcustomeraccountfund.New(db, log)

			if tc.idStr != invalidParam {
				customerAccNo, err := strconv.Atoi(tc.idStr)
				assert.Nil(t, err)

				if tc.expectedDBCalls != nil {
					tc.expectedDBCalls(db, customerAccNo, tc.rspBody)
				}
			}

			w := httptest.NewRecorder()
			path := fmt.Sprintf("/customer_accounts_funds/%s", tc.idStr)
			fmt.Println(path)

			req := httptest.NewRequest(http.MethodPost, path, nil)
			ctx := req.Context()
			ctx = context.WithValue(ctx, httprouter.ParamsKey, httprouter.Params{
				{
					Key:   "id",
					Value: tc.idStr,
				},
			})
			req = req.WithContext(ctx)

			handler.Handle(w, req)
			rsp := w.Result()

			assert.Equal(t, tc.rspStatusCode, rsp.StatusCode)
			if (tc.rspBody != models.CustomerAccountsFundDB{}) {
				checkRspBody(t, rsp.Body, tc.rspBody)
			}

			rsp.Body.Close()
			ctrl.Finish()
		})
	}
}

func checkRspBody(t *testing.T, actual io.Reader, expected models.CustomerAccountsFundDB) {
	actualRspBody, err := io.ReadAll(actual)
	assert.Nil(t, err)

	expectedRspBody, err := json.Marshal(expected)
	assert.Nil(t, err)

	assert.Equal(t, expectedRspBody, actualRspBody)
}
