package app

import (
	"go.temporal.io/sdk/workflow"
)

// @@@SNIPSTART onboarding-go-workflow

// SampleParentWorkflow workflow definition
func SampleParentWorkflow(ctx workflow.Context) (string, error) {
	logger := workflow.GetLogger(ctx)

	cwo := workflow.ChildWorkflowOptions{}
	ctx = workflow.WithChildOptions(ctx, cwo)

	// var result string
	var results []string
	err := workflow.ExecuteChildWorkflow(ctx, SampleChildWorkflow, "World").Get(ctx, &results)
	if err != nil {
		logger.Error("Parent execution received child execution failure.", "Error", err)
		return "", err
	}
	// logger.Info("Parent execution completed.", "Result", result)
	// return result, nil
	return "", nil
}

// @@@SNIPEND
