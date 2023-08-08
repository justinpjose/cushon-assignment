//go:build unit

package getavailablefunds_test

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
	"github.com/justinpjose/cushon-assignment/internal/handlers/getavailablefunds"
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
		name             string
		customerAccNoStr string
		expectedDBCalls  func(db *mocks.MockCustomerAccountsDB, customerAccNo int, rsp []models.AvailableFund)
		rspBody          []models.AvailableFund
		rspStatusCode    int
	}{
		{
			name:             "success",
			customerAccNoStr: "1",
			expectedDBCalls: func(db *mocks.MockCustomerAccountsDB, customerAccNo int, rsp []models.AvailableFund) {
				db.EXPECT().GetAvailableFunds(gomock.Any(), customerAccNo).Return(rsp, nil)
			},
			rspBody: []models.AvailableFund{
				{
					ID:   1,
					Name: "Cushon Equities Fund",
				},
			},
			rspStatusCode: http.StatusOK,
		},
		{
			name:             "invalid account number given",
			customerAccNoStr: invalidParam,
			rspStatusCode:    http.StatusBadRequest,
		},
		{
			name:             "failed to get available funds from db",
			customerAccNoStr: "1",
			expectedDBCalls: func(db *mocks.MockCustomerAccountsDB, customerAccNo int, rsp []models.AvailableFund) {
				db.EXPECT().GetAvailableFunds(gomock.Any(), customerAccNo).Return(nil, unexpectedErr)
			},
			rspStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			db := mocks.NewMockCustomerAccountsDB(ctrl)
			log := zerolog.NewMockLog()
			handler := getavailablefunds.New(db, log)

			if tc.customerAccNoStr != invalidParam {
				customerAccNo, err := strconv.Atoi(tc.customerAccNoStr)
				assert.Nil(t, err)

				if tc.expectedDBCalls != nil {
					tc.expectedDBCalls(db, customerAccNo, tc.rspBody)
				}
			}

			w := httptest.NewRecorder()
			path := fmt.Sprintf("/customer_accounts/%s/available_funds", tc.customerAccNoStr)
			fmt.Println(path)

			req := httptest.NewRequest(http.MethodPost, path, nil)
			ctx := req.Context()
			ctx = context.WithValue(ctx, httprouter.ParamsKey, httprouter.Params{
				{
					Key:   "accountNo",
					Value: tc.customerAccNoStr,
				},
			})
			req = req.WithContext(ctx)

			handler.Handle(w, req)
			rsp := w.Result()

			assert.Equal(t, tc.rspStatusCode, rsp.StatusCode)
			if tc.rspBody != nil {
				checkRspBody(t, rsp.Body, tc.rspBody)
			}

			rsp.Body.Close()
			ctrl.Finish()
		})
	}
}

func checkRspBody(t *testing.T, actual io.Reader, expected []models.AvailableFund) {
	actualRspBody, err := io.ReadAll(actual)
	assert.Nil(t, err)

	expectedRspBody, err := json.Marshal(expected)
	assert.Nil(t, err)

	assert.Equal(t, expectedRspBody, actualRspBody)
}
