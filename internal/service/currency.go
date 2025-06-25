//go:generate mockgen -source=currency.go -destination=mock/currency.go

package service

import (
	"github.com/erickrodriguez/currencygrpc/internal/model"
	"github.com/erickrodriguez/currencygrpc/internal/repository"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ICurrencyService interface {
	GetCurrencyRate(string) (model.Currency, error)
	GetExchange(string, float32) (float32, error)
}

type CurrencyService struct {
	cr repository.ICurrencyRepo
}

func (cs *CurrencyService) GetCurrencyRate(code string) (model.Currency, error) {
	result, found := cs.cr.Search(code)
	if !found {
		return model.Currency{}, status.Errorf(codes.NotFound, "CountryCode not Supported.")
	}

	return result, nil
}

func (cs *CurrencyService) GetExchange(code string, amount float32) (float32, error) {
	r, err := cs.GetCurrencyRate(code)
	if err != nil {
		return 0.00, err
	}

	exchange := amount / r.Exchange

	return exchange, nil
}

func NewCurrencyService(irc repository.ICurrencyRepo) ICurrencyService {
	return &CurrencyService{irc}
}
