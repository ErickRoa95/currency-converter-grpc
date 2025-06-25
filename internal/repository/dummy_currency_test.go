package repository

import (
	"testing"

	"github.com/erickrodriguez/currencygrpc/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestCurrentSearch(t *testing.T) {
	r := NewCurrencyRepo()
	testCases := []struct {
		name      string
		want      model.Currency
		wantError bool
		arg       string
	}{
		{
			name:      "OK Found CountryCurrency",
			want:      model.Currency{Base: "USD", CountryCode: "MXN", CurrencyName: "Mexican Peso", Exchange: float32(19.00649)},
			wantError: true,
			arg:       "MXN",
		},
		{
			name:      "ERROR Country not Supported",
			want:      model.Currency{},
			wantError: false,
			arg:       "MXS",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := r.Search(tc.arg)
			assert.Equal(t, tc.want, got)
			assert.Equal(t, tc.wantError, err)
		})
	}

}
