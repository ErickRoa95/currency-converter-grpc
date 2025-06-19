package repository

import (
	"github.com/erickrodriguez/currencygrpc/internal/model"
	"github.com/erickrodriguez/currencygrpc/internal/data"
)

type DummyCurrencyRepo struct{}

func (*DummyCurrencyRepo) Search(country_code string) (model.Currency, bool) {
	for i:=0 ; i<len(data.COUNTRY_CURRENCIES) ; i++{ 
		got := data.COUNTRY_CURRENCIES[i]
		if  got.CountryCode == country_code{ 
			return got, true
		}
	}

	return model.Currency{}, false
}

func NewDummyCurrencyRepo () *DummyCurrencyRepo{
	return &DummyCurrencyRepo{}
}