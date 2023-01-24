package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"tmprldemo/internal/customer/api"
	customerdata "tmprldemo/internal/customer/data/customer"
	customerpb "tmprldemo/internal/pb/customer/v1"

	_ "github.com/jackc/pgx/v4/stdlib"

	"google.golang.org/grpc"
)

// TODO: Add flags service parameters things like: address, ports etc

const (
	gRPCCustomerServiceAddress = "localhost:8080"
	postgresAddress            = "postgres"
	postgresPort               = "5432"
	postgresUser               = "postgres"
	postgresPass               = "root"
	postgresDB                 = "customer"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	connectionString := fmt.Sprintf(
		"postgres://%s:%s@localhost:%s/%s",
		postgresUser, postgresPass, postgresPort, postgresDB,
	)
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return fmt.Errorf("failed to open connection to db: %w", err)
	}

	customerDBCreator := customerdata.NewCustomerDBCreator(db)
	customerDBGetter := customerdata.NewCustomerDBGetter(db)
	customerServiceServer := api.NewCustomerServiceGRPCServer(customerDBCreator, customerDBGetter)

	server := grpc.NewServer()
	customerpb.RegisterCustomerServiceServer(server, customerServiceServer)

	listener, err := net.Listen("tcp", gRPCCustomerServiceAddress)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", gRPCCustomerServiceAddress, err)
	}

	log.Printf("gRPC server listening on %s", gRPCCustomerServiceAddress)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}
