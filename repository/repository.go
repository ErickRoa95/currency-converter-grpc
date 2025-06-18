 package repository
 
import "github.com/erickrodriguez/currencygrpc/model"

 type ICurrencyRepo interface{
	Found(string) model.Currency
}