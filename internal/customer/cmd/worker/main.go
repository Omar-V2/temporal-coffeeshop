package main

import (
	"log"
	"tmprldemo/internal/customer/workflows/verifyphone"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{HostPort: "temporal-server:7233"})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, "TEMPORAL_COFFEE_SHOP_TASK_QUEUE", worker.Options{})

	// Verify Phone Workflow
	smsSender := verifyphone.SMSSender{
		Sender: &verifyphone.MockSMSSender{},
	}
	w.RegisterActivity(&smsSender)
	w.RegisterWorkflow(verifyphone.NewWorkflow)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start worker", err)
	}
}
