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
	host        = flag.String("port", "localhost:50051", "Port of the server is running.")
	countryCode = flag.String("code", "MXN", "Country code to receive currency's converiton.")
	rpc         = flag.String("rpc", "Converter", "Call specific RPC endpoint.")
	amount      = flag.Float64("amount", 10.00, "Amount to convert.")
)

func main() {
	flag.Parse()
	log.Printf("=== Client call -> %s\n", *host)

	switch *rpc {
	case "Converter":
		ClientCallConverter(*countryCode)
	case "Exchange":
		a := float32(*amount)
		ClientCallExchange(*countryCode, a)
	default:
		log.Println(" == No instructions were provided == ")
	}

	log.Println("=== Client Call completed ===")
}

// Call  GRPC service Converter.
func ClientCallConverter(countryCode string) {
	log.Println("== CALL to CONVERTER RPC endpoint. ==")

	conn, err := grpc.NewClient(*host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("FAILED| Could not create GRPC Client: %v", err)
	}

	// Create Client for Currency GRPC server.
	defer conn.Close()
	c := pb.NewCurrencyClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Converter(ctx, &pb.ConverterRequest{CountryCode: &countryCode})
	if err != nil {
		log.Fatalf("FAILED| Couldn't complete RPC call: %v", err)
	}

	log.Printf("== Call Response: %v", r)
}

func ClientCallExchange(countryCode string, amount float32) {
	log.Println("== CALL to EXCHANGE RPC endpoint. ==")

	conn, err := grpc.NewClient(*host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("FAILED: Could not create GRPC Client: %v", err)
	}

	// Create Client for Currency GRPC server.
	defer conn.Close()
	c := pb.NewCurrencyClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Exchange(ctx, &pb.ExchangeRequest{CountryCode: &countryCode, Amount: &amount})
	if err != nil {
		log.Fatalf("FAILED| Couldn't complete RPC call : %v", err)
	}

	log.Printf(" == Call Response: %v", r)
}
