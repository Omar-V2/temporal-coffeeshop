package main

import (
	"fmt"
	"log"
	"tmprldemo/internal/customer/workflows/verifyphone"

	"github.com/ilyakaznacheev/cleanenv"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

type Config struct {
	TemporalAddress   string `env:"TEMPORAL_ADDRESS" env-default:"temporal:7233"`
	TemporalTaskQueue string `env:"TEMPORAL_TASK_QUEUE" env-default:"TEMPORAL_COFFEE_SHOP_TASK_QUEUE"`
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

	c, err := client.Dial(client.Options{HostPort: cfg.TemporalAddress})
	if err != nil {
		return fmt.Errorf("unable to create Temporal client: %w", err)
	}
	defer c.Close()

	w := worker.New(c, cfg.TemporalTaskQueue, worker.Options{})

	mockSMSSender := &verifyphone.MockSMSSender{}
	// uncomment this line and use it in place of mockSMSSender to simulate an activity failure
	// faultySMSSender := &verifyphone.FaultySMSSender{}
	smsSender := verifyphone.SMSSender{
		Sender: mockSMSSender,
	}
	w.RegisterActivity(&smsSender)
	w.RegisterWorkflow(verifyphone.NewWorkflow)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		return fmt.Errorf("unable to start worker: %w", err)
	}

	return nil
}
