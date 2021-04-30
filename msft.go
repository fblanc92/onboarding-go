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
	invocations := [9]string{"MSFT_sign in office.com", "MSFT_Select License", "MSFT_Add User Data", "MSFT_Generate Password", "MSFT_Configure New Password Policies", "MSFT_Finish User Creation", "MSFT_Add User To Groups", "MSFT_Add User To Calendars", "MSFT_Configure Calendars"}
	var result, result_msft string
	for i := 0; i < len(invocations); i++ {
		err := workflow.ExecuteActivity(ctx, MsftActivity, invocations[i]).Get(ctx, &result)
		result_msft += result + time.Now().String() + "\n"
		if err != nil {
			return result, err
		}
	}
	return result_msft, nil
}
