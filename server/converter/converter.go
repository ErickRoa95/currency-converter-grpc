package converter

import (
	"context"
	"log"

	pb "github.com/erickrodriguez/currencygrpc/currencygrpc/currencygrpc"
	"github.com/erickrodriguez/currencygrpc/internal/repository"
	"github.com/erickrodriguez/currencygrpc/internal/service"
)

var base string = "USD"

type Server struct {
	pb.UnimplementedCurrencyServer
}

func (s *Server) Converter(ctx context.Context, req *pb.ConverterRequest) (*pb.ConverterResponse, error) {
	log.Printf("Received: %v", req)

	r, err := service.NewCurrencyService(repository.NewCurrencyRepo()).GetCurrencyRate(*req.CountryCode)
	if err != nil {
		return nil, err
	}

	log.Printf("value : %s", *req.CountryCode)

	return &pb.ConverterResponse{
		CurrencyName: &r.CurrencyName,
		Base:         &r.Base,
		Amount:       &r.Exchange,
		CountryCode:  &r.CountryCode,
	}, nil
}

func (s *Server) Exchange(ct context.Context, req *pb.ExchangeRequest) (*pb.ExchangeResponse, error) {
	log.Printf("=== Exchange RPC call !")
	log.Printf("Received : %v", req)

	amount := *req.Amount

	cs := service.NewCurrencyService(repository.NewCurrencyRepo())
	r, err := cs.GetExchange(*req.CountryCode, amount)
	if err != nil {
		return nil, err
	}

	return &pb.ExchangeResponse{
		CountryCode: req.CountryCode,
		Amount:      &amount,
		Base:        &base,
		Exchange:    &r,
	}, nil

}
