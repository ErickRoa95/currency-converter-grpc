package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/erickrodriguez/currencygrpc/currencygrpc/currencygrpc"
	"google.golang.org/grpc"
)

var(
	port = flag.Int("port", 50051, "The server port")
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

func main (){
	flag.Parse()
	addr:= fmt.Sprintf(":%d", *port)
	list, err:= net.Listen("tcp", addr)
	if err!=nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCurrencyServer(s, &server{})
	log.Printf("server listening at %v", list.Addr())
	if err:= s.Serve(list); err != nil {
		log.Fatalf("Failed to serve : %v", err)
	}
}