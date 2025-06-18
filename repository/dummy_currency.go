package repository

import (
	"github.com/erickrodriguez/currencygrpc/model"
	"github.com/erickrodriguez/currencygrpc/data"
)

type DummyCurrencyRepo struct{}

func (*DummyCurrencyRepo) Search(country_code string) (model.Currency, bool) {
	for i:=0 ; i<len(data.country_codes); i++{ 
		got := data.country_codes[i]
		if  got.CountryCode == country_code{ 
			got, true
		}
	}

	return new(model.Currency), false
}