package app

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

// SampleChildWorkflow workflow definition
func MsftWorkflow(ctx workflow.Context, name string) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	logger.Info("MSFT workflow execution")
	var result string
	err := workflow.ExecuteActivity(ctx, MsftActivity, "msft").Get(ctx, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
