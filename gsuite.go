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
	invocations := [10]string{"Gsuite_Sign in admin.google.com", "Gsuite_Add user data", "Gsuite_Update Org Units", "Gsuite_Fill Wellcome Template", "Gsuite_Create New Password", "Gsuite_Finish User Creation", "Gsuite_Close Ticket", "Gsuite_Finish Welcome Template", "Gsuite_Add User To Groups", "Gsuite_Add User To Calendars With Correct Permissions"}
	var result, result_gsuite string
	for i := 0; i < len(invocations); i++ {
		err := workflow.ExecuteActivity(ctx, GsuiteActivity, invocations[i]).Get(ctx, &result)
		// result_gsuite += result + time.Now().String() + "\n"
		current_time := time.Now()
		result_gsuite += "\n" + result + " - " + current_time.Format("2006-01-02 15:04:05")
		if err != nil {
			return result, err
		}
	}
	return result_gsuite, nil
}
