//go:generate mockgen -source=currency.go -destination=mock/currency.go

package service

import (
	"github.com/erickrodriguez/currencygrpc/internal/model"
	"github.com/erickrodriguez/currencygrpc/internal/repository"

  "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ICurrencyService interface{
  GetCurrencyRate(string) (model.Currency,error)
}

type CurrencyService struct {
  cr repository.ICurrencyRepo
}

func (cs *CurrencyService) GetCurrencyRate  (code string) (model.Currency, error){
  result, found := cs.cr.Search(code)
  if !found {
    return model.Currency{}, status.Errorf(codes.NotFound, "CountryCode not Supported.")
  }

  return result, nil
}

func NewCurrencyService (irc repository.ICurrencyRepo) ICurrencyService{
  return &CurrencyService{irc}
}
