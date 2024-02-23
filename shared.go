package app

// @@@SNIPSTART github.com/Raumo0/temporal-money-transfer-template-shared-task-queue
const MoneyTransferTaskQueueName = "TRANSFER_MONEY_TASK_QUEUE"

// @@@SNIPEND

// @@@SNIPSTART github.com/Raumo0/temporal-money-transfer-template-transferdetails
type PaymentDetails struct {
	SourceAccount string
	TargetAccount string
	Amount        int
	ReferenceID   string
}

// @@@SNIPEND
