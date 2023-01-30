package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"tmprldemo/internal/customer/api"
	customerdata "tmprldemo/internal/customer/data/customer"
	customerpb "tmprldemo/internal/pb/customer/v1"

	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go.temporal.io/sdk/client"

	"google.golang.org/grpc"
)

type Config struct {
	CustomerServiceAddress string `env:"CUSTOMER_SERVICE_ADDRESS" env-default:"customer-service:8080"`
	PostgresPort           string `env:"POSTGRES_PORT" env-default:"5432"`
	PostgresAddress        string `env:"POSTGRES_ADDRESS" env-default:"postgres"`
	PostgresUser           string `env:"POSTGRES_USER" env-default:"postgres"`
	PostgresPassword       string `env:"POSTGRES_PASSWORD" env-default:"root"`
	PostgresDB             string `env:"POSTGRES_DB" env-default:"customer"`
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return fmt.Errorf("failed to read environment variables: %w", err)
	}

	connectionString := fmt.Sprintf(
		"postgres://%s:%s@postgres:%s/%s",
		cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresPort, cfg.PostgresDB,
	)
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return fmt.Errorf("failed to open connection to db: %w", err)
	}

	temporalClient, err := client.Dial(client.Options{HostPort: "temporal-server:7233"})
	if err != nil {
		return fmt.Errorf("failed to instantiate temporal client: %w", err)
	}

	customerDBCreator := customerdata.NewCustomerDBCreator(db)
	customerDBGetter := customerdata.NewCustomerDBGetter(db)
	customerServiceServer := api.NewCustomerServiceGRPCServer(
		customerDBCreator, customerDBGetter, temporalClient,
	)

	server := grpc.NewServer()
	customerpb.RegisterCustomerServiceServer(server, customerServiceServer)

	listener, err := net.Listen("tcp", cfg.CustomerServiceAddress)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", cfg.CustomerServiceAddress, err)
	}

	log.Printf("gRPC server listening on %s", cfg.CustomerServiceAddress)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}
