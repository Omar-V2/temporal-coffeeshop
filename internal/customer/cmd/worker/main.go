package main

import (
	"database/sql"
	"fmt"
	"log"

	customerdata "tmprldemo/internal/customer/data/customer"
	"tmprldemo/internal/customer/workflows/verifyphone"

	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

type Config struct {
	TemporalAddress   string `env:"TEMPORAL_ADDRESS" env-default:"temporal:7233"`
	TemporalTaskQueue string `env:"TEMPORAL_TASK_QUEUE" env-default:"TEMPORAL_COFFEE_SHOP_TASK_QUEUE"`
	PostgresPort      string `env:"POSTGRES_PORT" env-default:"5432"`
	PostgresHost      string `env:"POSTGRES_HOST" env-default:"postgres"`
	PostgresUser      string `env:"POSTGRES_USER" env-default:"postgres"`
	PostgresPassword  string `env:"POSTGRES_PASSWORD" env-default:"root"`
	PostgresDB        string `env:"POSTGRES_DB" env-default:"customer"`
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return fmt.Errorf("failed to read config from environment variables: %w", err)
	}

	temporalClient, err := client.Dial(client.Options{HostPort: cfg.TemporalAddress})
	if err != nil {
		return fmt.Errorf("unable to create Temporal client: %w", err)
	}
	defer temporalClient.Close()

	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDB,
	)
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return fmt.Errorf("failed to open connection to db: %w", err)
	}

	customerVerifier := customerdata.NewCustomerDBVerifier(db)

	//uncomment this line in place of the line below it to enable random code generation
	// codeGenerator := verifyphone.RandomCodeGenerator{}
	codeGenerator := verifyphone.StaticCodeGenerator{}

	// uncomment this line and use it in place of mockSMSSender to simulate an activity failure
	// faultySMSSender := &verifyphone.FaultySMSSender{}
	mockSMSSender := &verifyphone.MockSMSSender{}
	activities := verifyphone.NewActivities(mockSMSSender, customerVerifier, codeGenerator)

	w := worker.New(temporalClient, cfg.TemporalTaskQueue, worker.Options{DisableRegistrationAliasing: true})

	w.RegisterActivity(activities)
	w.RegisterWorkflowWithOptions(verifyphone.NewWorkflow, workflow.RegisterOptions{
		Name: "VerifyPhoneWorkflow",
	})

	err = w.Run(worker.InterruptCh())
	if err != nil {
		return fmt.Errorf("unable to start worker: %w", err)
	}

	return nil
}
