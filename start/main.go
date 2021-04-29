package main

import (
	"context"
	"log"
	"onboarding-go/app"

	"go.temporal.io/sdk/client"
)

// @@@SNIPSTART onboarding-go-start-workflow
func main() {
	// Create the client object just once per process
	// onboardingDetails := app.OnBoardingTask{
	// 	TaskID:       1,
	// 	FromDivision: "Turbine",
	// 	ToDivision:   "Gmail",
	// 	TrackingID:   1,
	// }

	// weO, errO := c.ExecuteWorkflow(context.Background(), optionsO, app.RequestOnboarding, onboardingDetails)

	// if errO != nil {
	// 	log.Fatalln("error starting RequestOnboarding workflow", errO)
	// }
	// log.Println("Started workflow",
	// 	"WorkflowID", weO.GetID(), "RunID", weO.GetRunID())
	// printOnboardingResults(onboardingDetails, weO.GetID(), weO.GetRunID())

	// The client is a heavyweight object that should be created once per process.
	c, err := client.NewClient(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// This workflow ID can be user business logic identifier as well.
	workflowID := "parent-workflow_"
	workflowOptions := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: "child-workflow",
	}

	workflowRun, err := c.ExecuteWorkflow(context.Background(), workflowOptions, app.SampleParentWorkflow)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow",
		"WorkflowID", workflowRun.GetID(), "RunID", workflowRun.GetRunID())

	// Synchronously wait for the workflow completion. Behind the scenes the SDK performs a long poll operation.
	// If you need to wait for the workflow completion from another process use
	// Client.GetWorkflow API to get an instance of a WorkflowRun.
	var result string
	err = workflowRun.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Failure getting workflow result", err)
	}
	log.Println("Workflow result: %v", "result", result)
}

// @@@SNIPEND

// func printOnboardingResults(onboardingDetails app.OnBoardingTask, workflowID, runID string) {
// 	log.Printf(
// 		"\n\n\nOnboarding number: %d. Status: \nFrom: %s to %s.\nTracking ID: %d",
// 		onboardingDetails.TaskID,
// 		onboardingDetails.FromDivision,
// 		onboardingDetails.ToDivision,
// 		onboardingDetails.TrackingID,
// 	)
// 	log.Printf(
// 		"\nWorkflowID: %s RunID: %s\n",
// 		workflowID,
// 		runID,
// 	)

// }
