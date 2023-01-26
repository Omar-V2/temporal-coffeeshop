package main

import (
	"context"
	"log"
	"time"
	"tmprldemo/internal/customer/workflows/verifyphone"

	"go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/client"
)

func main() {
	// Create the client object just once per process
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:                       "customer_id1",
		TaskQueue:                "TEMPORAL_COFFEE_SHOP_TASK_QUEUE",
		WorkflowExecutionTimeout: time.Minute * 10,
		WorkflowIDReusePolicy:    enums.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY,
	}

	verifyPhoneParams := verifyphone.WorkflowParams{
		PhoneNumber:          "+447500140",
		MaximumAttempts:      3,
		CodeValidityDuration: time.Minute * 15,
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, verifyphone.NewWorkflow, verifyPhoneParams)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}
	err = we.Get(context.Background(), nil)
}
