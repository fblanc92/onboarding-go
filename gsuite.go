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
	invocations := [10]string{"acc_creation", "1", "2", "1", "2", "1", "2", "1", "2", "fin"}
	var result, result_gsuite string
	for i := 0; i < 10; i++ {
		err := workflow.ExecuteActivity(ctx, GsuiteActivity, invocations[i]).Get(ctx, &result)
		result_gsuite += result + time.Now().String() + "\n"
		if err != nil {
			return result, err
		}
	}
	return result_gsuite, nil
}
