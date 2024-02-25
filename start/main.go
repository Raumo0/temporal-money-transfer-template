package main

import (
	"context"
	"log"
	"time"

	"go.temporal.io/sdk/client"

	"github.com/Raumo0/temporal-money-transfer-template"
)

// @@@SNIPSTART github.com/Raumo0/temporal-money-transfer-template-start-workflow
func main() {
	// Create the client object just once per process
	c, err := client.Dial(client.Options{})

	if err != nil {
		log.Fatalln("Unable to create Temporal client:", err)
	}

	defer c.Close()

	depositTime := time.Now().UTC().Add(time.Minute * 2)
	input := app.PaymentDetails{
		SourceAccount: "85-150",
		TargetAccount: "43-812",
		Amount:        250,
		ReferenceID:   "12345",
		DepositTime:   depositTime,
	}

	options := client.StartWorkflowOptions{
		ID:         "pay-invoice-701",
		TaskQueue:  app.MoneyTransferTaskQueueName,
		StartDelay: time.Minute,
	}

	log.Printf("Starting transfer from account %s to account %s for %d", input.SourceAccount, input.TargetAccount, input.Amount)

	we, err := c.ExecuteWorkflow(context.Background(), options, app.MoneyTransfer, input)
	if err != nil {
		log.Fatalln("Unable to start the Workflow:", err)
	}

	log.Printf("WorkflowID: %s RunID: %s\n", we.GetID(), we.GetRunID())

	var result string

	err = we.Get(context.Background(), &result)

	if err != nil {
		log.Fatalln("Unable to get Workflow result:", err)
	}

	log.Println(result)
}

// @@@SNIPEND
