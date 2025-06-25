package service

import (
	"testing"

	"github.com/erickrodriguez/currencygrpc/internal/model"
	"github.com/erickrodriguez/currencygrpc/internal/repository"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCurrencyServiceIntegration_GetCurrencyRate(t *testing.T) {

	icr := repository.NewDummyCurrencyRepo()
	s := NewCurrencyService(icr)

	testCases := []struct {
		name      string
		want      model.Currency
		wantError error
		arg       string
	}{
		{
			name:      "Success",
			want:      model.Currency{Base: "USD", CountryCode: "MXN", CurrencyName: "Mexican Peso", Exchange: float32(19.00649)},
			wantError: nil,
			arg:       "MXN",
		},
		{
			name:      "CountryCode not supported",
			want:      model.Currency{},
			wantError: status.Errorf(codes.NotFound, "CountryCode not Supported."),
			arg:       "MX",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := s.GetCurrencyRate(tc.arg)
			assert.Equal(t, tc.want, got)
			assert.Equal(t, tc.wantError, err)
		})
	}
}
