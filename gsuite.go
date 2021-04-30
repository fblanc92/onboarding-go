package app

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

// SampleChildWorkflow workflow definition
func GsuiteWorkflow(ctx workflow.Context, name string) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	logger.Info("Gsuite workflow execution")
	var result string
	err := workflow.ExecuteActivity(ctx, GsuiteActivity, "gsuite").Get(ctx, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
