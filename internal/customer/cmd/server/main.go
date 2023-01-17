package main

import (
	"fmt"
	"log"
	"net"
	customerpb "tmprldemo/internal/customer/pb/customer/v1"

	"google.golang.org/grpc"
)

// TODO: Add flags service parameters things like: address, ports etc

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	server := grpc.NewServer()
	customerpb.RegisterCustomerServiceServer(server, &customerServiceServer{})

	address := "127.0.0.1:8000"
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", address, err)
	}

	log.Printf("gRPC server listening on %s", address)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}

type customerServiceServer struct {
	customerpb.UnimplementedCustomerServiceServer
}
