package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"onboarding-go/app"
)

// @@@SNIPSTART onboarding-go-worker
func main() {
	// Create the client object just once per process
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	// This worker hosts both Worker and Activity functions

	wT := worker.New(c, app.TransferMoneyTaskQueue, worker.Options{})
	wT.RegisterWorkflow(app.TransferMoney)
	wT.RegisterActivity(app.Withdraw)
	wT.RegisterActivity(app.Deposit)

	wO := worker.New(c, app.OnboardingTaskQueue, worker.Options{})
	wO.RegisterWorkflow(app.RequestOnboarding)
	wO.RegisterActivity(app.GotoNextDivision)

	// Start listening to the Task Queue
	run_option := "Onboarding"
	if run_option == "Transfer" {
		err = wT.Run(worker.InterruptCh())
		if err != nil {
			log.Fatalln("unable to start TRANSFER Worker", err)
		}
	} else if run_option == "Onboarding" {
		err = wO.Run(worker.InterruptCh())
		if err != nil {
			log.Fatalln("unable to start ONBOARDING Worker", err)
		}
	}
}

// @@@SNIPEND
