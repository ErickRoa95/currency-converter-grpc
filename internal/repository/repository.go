 package repository
 
import "github.com/erickrodriguez/currencygrpc/internal/model"

 type ICurrencyRepo interface{
	Search(string) (model.Currency, bool)
}

func NewCurrencyRepo () ICurrencyRepo{
	return NewDummyCurrencyRepo()
}