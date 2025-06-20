package converter

import (
	"context"
	"log"

	pb "github.com/erickrodriguez/currencygrpc/currencygrpc/currencygrpc"
	"github.com/erickrodriguez/currencygrpc/internal/repository"
	"github.com/erickrodriguez/currencygrpc/internal/service"
)


type Server struct {
	pb.UnimplementedCurrencyServer
}

func (s *Server) Converter(ctx context.Context, req *pb.ConverterRequest) (*pb.ConverterResponse, error){
	log.Printf("Received: %v", req)

	r, err := service.NewCurrencyService(repository.NewCurrencyRepo()).GetCurrencyRate(*req.CountryCode)
	if err != nil {
		return nil, err
	}

	log.Printf("value : %s", *req.CountryCode)

	return &pb.ConverterResponse{
		CurrencyName: &r.CurrencyName,
		Base: &r.Base, 
		Amount: &r.Exchange,
		CountryCode: &r.CountryCode,
	}, nil
}