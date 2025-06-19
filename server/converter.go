package main

import (
	"context"
	"log"

	pb "github.com/erickrodriguez/currencygrpc/currencygrpc/currencygrpc"
)


type server struct {
	pb.UnimplementedCurrencyServer
}

func (s *server) Converter(ctx context.Context, req *pb.ConverterRequest) (*pb.ConverterResponse, error){
	log.Printf("Received: %v", req)
	currency := "Mexican pesos"
	dollar := int32(1)
	amount:= int32(19)
	return &pb.ConverterResponse{
		CurrencyName: &currency,
		Dollar: &dollar, 
		Amount: &amount,
	}, nil
}