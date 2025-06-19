package service

import (
	"log"

	"github.com/erickrodriguez/currencygrpc/internal/model"
	"github.com/erickrodriguez/currencygrpc/internal/repository"
)

type CurrencyService struct {}

func (*CurrencyService) GetCurrencyRate  (code string) (model.Currency){
  r := repository.NewCurrencyRepo()
  result, found := r.Search(code)
  if !found {
    log.Fatalf("ContryCode is not supported.")
  }

  return result
}

func (cs *CurrencyService) ExchangeCurrency(){}

func NewCurrencyService () *CurrencyService{
  return new(CurrencyService)
}
