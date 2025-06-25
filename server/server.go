package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/erickrodriguez/currencygrpc/currencygrpc/currencygrpc"
	"github.com/erickrodriguez/currencygrpc/server/converter"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	addr := fmt.Sprintf(":%d", *port)

	// Initied listening port for grpc server.
	list, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	// Create & register server dependencies.
	s := grpc.NewServer()
	pb.RegisterCurrencyServer(s, &converter.Server{})

	// Serve Grpc Server at specific port.
	log.Printf("GRPC Server listening at %v !", list.Addr())
	if err := s.Serve(list); err != nil {
		log.Fatalf("Failed to serve : %v", err)
	}
}
