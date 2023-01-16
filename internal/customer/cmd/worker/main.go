package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	verifyphonewf "tmprldemo/internal/customer/workflows/verifyphone"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, "TEMPORAL_COFEE_SHOP_TASK_QUEUE", worker.Options{})

	// Verify Phone Workflow
	smsSender := verifyphonewf.SMSSender{
		Sender: &verifyphonewf.MockSMSSender{},
	}
	w.RegisterActivity(&smsSender)
	w.RegisterWorkflow(verifyphonewf.NewVerifyPhoneWorkflow)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start worker", err)
	}
}
