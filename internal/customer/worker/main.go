package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	wf "tmprldemo/internal/customer/workflows"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, "GREETING_TASK_QUEUE", worker.Options{})
	w.RegisterWorkflow(wf.GreetingWorkflow)
	w.RegisterActivity(wf.ComposeGreeting)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start worker", err)
	}
}