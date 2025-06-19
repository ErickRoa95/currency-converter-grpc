package service

import (
	"github.com/erickrodriguez/currencygrpc/internal/model"
	"github.com/erickrodriguez/currencygrpc/internal/repository"

  "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CurrencyService struct {}

func (*CurrencyService) GetCurrencyRate  (code string) (model.Currency, error){
  r := repository.NewCurrencyRepo()
  result, found := r.Search(code)
  if !found {
    return model.Currency{}, status.Errorf(codes.NotFound, "CountryCode not Supported.")
  }

  return result, nil
}

func NewCurrencyService () *CurrencyService{
  return new(CurrencyService)
}
