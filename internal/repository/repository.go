//go:generate mockgen -source=repository.go -destination=mock/repository.go

package repository

import "github.com/erickrodriguez/currencygrpc/internal/model"

type ICurrencyRepo interface {
	Search(string) (model.Currency, bool)
}

func NewCurrencyRepo() ICurrencyRepo {
	return NewDummyCurrencyRepo()
}
