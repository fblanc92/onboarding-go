package app

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

// SampleChildWorkflow workflow definition
func SampleChildWorkflow(ctx workflow.Context, name string) ([]string, error) {
	logger := workflow.GetLogger(ctx)
	defer logger.Info("Workflow completed.")

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	future1, settable1 := workflow.NewFuture(ctx)
	workflow.Go(ctx, func(ctx workflow.Context) {
		defer logger.Info("First goroutine completed.")

		var results []string
		var result string
		err := workflow.ExecuteActivity(ctx, SampleActivity, "branch1.1").Get(ctx, &result)
		if err != nil {
			settable1.SetError(err)
			return
		}
		results = append(results, result)
		err = workflow.ExecuteActivity(ctx, SampleActivity, "branch1.2").Get(ctx, &result)
		if err != nil {
			settable1.SetError(err)
			return
		}
		results = append(results, result)
		settable1.SetValue(results)
	})

	future2, settable2 := workflow.NewFuture(ctx)
	workflow.Go(ctx, func(ctx workflow.Context) {
		defer logger.Info("Second goroutine completed.")

		var result string
		err := workflow.ExecuteActivity(ctx, SampleActivity, "branch2").Get(ctx, &result)
		settable2.Set(result, err)
	})

	var results []string
	// Future.Get returns error from Settable.SetError
	// Note that the first goroutine puts a slice into the settable while the second a string value
	err := future1.Get(ctx, &results)
	if err != nil {
		return nil, err
	}
	var result string
	err = future2.Get(ctx, &result)
	if err != nil {
		return nil, err
	}
	results = append(results, result)

	return results, nil
}
