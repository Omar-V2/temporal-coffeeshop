package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"tmprldemo/internal/customer/api"
	customerdata "tmprldemo/internal/customer/data/customer"
	customerpb "tmprldemo/internal/customer/pb/customer/v1"

	sq "github.com/Masterminds/squirrel"
	"google.golang.org/grpc"
)

// TODO: Add flags service parameters things like: address, ports etc

const (
	gRPCCustomerServiceAddress = "localhost:8080"
	postgresAddress            = "POSTGRES_URL"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	db, err := sql.Open("pgx", postgresAddress)
	if err != nil {
		return fmt.Errorf("failed to open connection to db: %w", err)
	}

	sqStatementBuilder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	customerDBCreator := customerdata.NewCustomerDBCreator(db, sqStatementBuilder)
	customerServiceServer := api.NewCustomerServiceGRPCServer(customerDBCreator)

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
