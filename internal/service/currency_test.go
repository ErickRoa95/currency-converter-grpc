package service

import (
	"testing"

	"github.com/erickrodriguez/currencygrpc/internal/model"
	mock_repository "github.com/erickrodriguez/currencygrpc/internal/repository/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCurrencyService_GetCurrencyRate(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	icr := mock_repository.NewMockICurrencyRepo(ctrl)
	s := NewCurrencyService(icr)
	
	testCases := []struct {
		name	string
		mockFunc 	func()
		want model.Currency
		wantError error	
		arg string
	}{
		{
			name: "Success",
			mockFunc: func() {
				icr.EXPECT().Search("MXN").Return(model.Currency{ Base:"USD", CountryCode:"MXN", CurrencyName: "Mexican Peso",  Exchange: float32(19.00649) }, true).Times(1)
			},
			want: model.Currency{ Base:"USD", CountryCode:"MXN", CurrencyName: "Mexican Peso",  Exchange: float32(19.00649) },
			wantError: nil,
			arg: "MXN",
		},
		{
			name: "CountryCode not supported",
			mockFunc: func() {
				icr.EXPECT().Search("MX").Return(model.Currency{},false).Times(1)
			},
			want: model.Currency{},
			wantError: status.Errorf(codes.NotFound, "CountryCode not Supported."),
			arg: "MX",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T){
			tc.mockFunc()
			got,err := s.GetCurrencyRate(tc.arg)
			assert.Equal(t, tc.want,got)
			assert.Equal(t, tc.wantError, err)
		})
	}
}
