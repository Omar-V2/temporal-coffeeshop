package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	customerpb "tmprldemo/internal/customer/pb/customer/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	gRPCCustomerServiceAddress = "localhost:8080"
	gRPCGatewayAddress         = "localhost:8081"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	mux := runtime.NewServeMux()

	err := customerpb.RegisterCustomerServiceHandlerFromEndpoint(
		context.Background(),
		mux,
		gRPCCustomerServiceAddress,
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	)
	if err != nil {
		return fmt.Errorf("failed to serve gRPC gateway: %w", err)
	}

	log.Printf("gRPC gateway server listening on %s", gRPCGatewayAddress)
	http.ListenAndServe(gRPCGatewayAddress, mux)

	return nil
}
