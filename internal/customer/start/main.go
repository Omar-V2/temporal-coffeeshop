package main

import (
	"context"
	"log"
	"time"

	"go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/client"

	verifyphonewf "tmprldemo/internal/customer/workflows/verifyphone"
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
		TaskQueue:                "TEMPORAL_COFEE_SHOP_TASK_QUEUE",
		WorkflowExecutionTimeout: time.Minute * 10,
		WorkflowIDReusePolicy:    enums.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY,
	}

	verifyPhoneParams := verifyphonewf.VerifyPhoneWorkflowParams{
		PhoneNumber:          "+447500140",
		MaximumAttempts:      3,
		CodeValidityDuration: time.Minute * 15,
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, verifyphonewf.NewVerifyPhoneWorkflow, verifyPhoneParams)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}
	err = we.Get(context.Background(), nil)
}
