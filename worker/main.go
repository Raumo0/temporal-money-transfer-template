package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/Raumo0/temporal-money-transfer-template"
)

// @@@SNIPSTART github.com/Raumo0/temporal-money-transfer-template-worker
func main() {

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client.", err)
	}
	defer c.Close()

	options := worker.Options{
		MaxConcurrentWorkflowTaskPollers:        2,
		MaxConcurrentActivityTaskPollers:        2,
		MaxConcurrentLocalActivityExecutionSize: 0,
		MaxConcurrentActivityExecutionSize:      1, // single thread activity execution
		MaxConcurrentWorkflowTaskExecutionSize:  2,
		MaxConcurrentSessionExecutionSize:       0,
		MaxConcurrentEagerActivityExecutionSize: 0,
		DisableWorkflowWorker:                   false,
		DisableEagerActivities:                  true,
	}
	w := worker.New(c, app.MoneyTransferTaskQueueName, options)

	// This worker hosts both Workflow and Activity functions.
	w.RegisterWorkflow(app.MoneyTransfer)
	w.RegisterActivity(app.Withdraw)
	w.RegisterActivity(app.Deposit)
	w.RegisterActivity(app.Refund)

	// Start listening to the Task Queue.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}

// @@@SNIPEND
