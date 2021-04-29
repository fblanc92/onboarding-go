package main

import (
	"context"
	"log"

	"github.com/google/uuid"
	"go.temporal.io/sdk/client"

	"onboarding-go/app"
)

// @@@SNIPSTART onboarding-go-start-workflow
func main() {
	// Create the client object just once per process
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	optionsT := client.StartWorkflowOptions{
		ID:        "transfer-money-workflow",
		TaskQueue: app.TransferMoneyTaskQueue,
	}
	optionsO := client.StartWorkflowOptions{
		ID:        "onboarding-workflow",
		TaskQueue: app.OnboardingTaskQueue,
	}
	transferDetails := app.TransferDetails{
		Amount:      54.99,
		FromAccount: "001-001",
		ToAccount:   "002-002",
		ReferenceID: uuid.New().String(),
	}

	onboardingDetails := app.OnBoardingTask{
		TaskID:       1,
		FromDivision: "Turbine",
		ToDivision:   "Gmail",
		TrackingID:   1,
	}

	weT, errT := c.ExecuteWorkflow(context.Background(), optionsT, app.TransferMoney, transferDetails)
	weO, errO := c.ExecuteWorkflow(context.Background(), optionsO, app.RequestOnboarding, onboardingDetails)
	if errT != nil {
		log.Fatalln("Error starting TransferMoney workflow", errT)
	}
	if errO != nil {
		log.Fatalln("error starting RequestOnboarding workflow", errO)
	}
	printResults(transferDetails, weT.GetID(), weT.GetRunID())
	printOnboardingResults(onboardingDetails, weO.GetID(), weO.GetRunID())
}

// @@@SNIPEND

func printResults(transferDetails app.TransferDetails, workflowID, runID string) {
	log.Printf(
		"\nTransfer of $%f from account %s to account %s is processing. ReferenceID: %s\n",
		transferDetails.Amount,
		transferDetails.FromAccount,
		transferDetails.ToAccount,
		transferDetails.ReferenceID,
	)
	log.Printf(
		"\nWorkflowID: %s RunID: %s\n",
		workflowID,
		runID,
	)
}

func printOnboardingResults(onboardingDetails app.OnBoardingTask, workflowID, runID string) {
	log.Printf(
		"\n\n\nOnboarding number: %d. Status: \nFrom: %s to %s.\nTracking ID: %d",
		onboardingDetails.TaskID,
		onboardingDetails.FromDivision,
		onboardingDetails.ToDivision,
		onboardingDetails.TrackingID,
	)
	log.Printf(
		"\nWorkflowID: %s RunID: %s\n",
		workflowID,
		runID,
	)

}
