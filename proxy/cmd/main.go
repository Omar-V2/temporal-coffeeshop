package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	customerpb "tmprldemo/internal/pb/customer/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/ilyakaznacheev/cleanenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	CustomerServiceAddress string `env:"CUSTOMER_SERVICE_ADDRESS" env-default:"customer-service:8080"`
	GatewayAddress         string `env:"GATEWAY_ADDRESS" env-default:"0.0.0.0:8081"`
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return fmt.Errorf("failed to read config from environment variables: %w", err)
	}
	mux := runtime.NewServeMux()

	err := customerpb.RegisterCustomerServiceHandlerFromEndpoint(
		context.Background(),
		mux,
		cfg.CustomerServiceAddress,
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	)
	if err != nil {
		return fmt.Errorf("failed to serve gRPC gateway: %w", err)
	}

	log.Printf("gRPC gateway server listening on %s", cfg.GatewayAddress)
	http.ListenAndServe(cfg.GatewayAddress, mux)

	return nil
}
