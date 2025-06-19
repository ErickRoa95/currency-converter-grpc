package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/erickrodriguez/currencygrpc/currencygrpc/currencygrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	host = flag.String("port", "localhost:50051", "Port of the server is running.")
	countryCode = flag.String("code", "MX", "Country code to receive currency's converiton.")
)

func main(){
	flag.Parse()
	log.Printf("%s", *host)

	r, err := initClient()
	if err != nil {
		log.Fatalf("Could't create GRPC client: %v", err)
	}

	log.Printf("Converter Response:  %v", r)
	log.Printf("==== Client Call completed ===")
}

func initClient()(*pb.ConverterResponse,error){
	conn,err := grpc.NewClient(*host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("FAILED: Could not create GRPC Client: %v", err)
	}

	defer conn.Close()
	c := pb.NewCurrencyClient(conn) 

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	
	r, err := CallConverter(c, ctx)
	if err != nil {
		log.Fatalf("FAILED: Could not receive response: %v", err)
	}

	return r,nil
}

func CallConverter (c pb.CurrencyClient, ctx context.Context) (*pb.ConverterResponse, error){
	r, err := c.Converter(ctx, &pb.ConverterRequest{CountryCode: countryCode})
	if err != nil {
		return nil, err
	}

	return r, nil
}