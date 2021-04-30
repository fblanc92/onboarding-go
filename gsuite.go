package app

import (
	"log"
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
	invocations := [10]string{"SIGN IN admin.google.com", "Add user data", "Update Org Units", "Fill Wellcome Template", "Create New Password", "Finish User Creation", "Close Ticket", "Finish Welcome Template", "Add User To Groups", "Add User To Calendars With Correct Permissions"}
	var result, result_gsuite string
	for i := 0; i < 10; i++ {
		if invocations[i] == "Create New Password" {
			err := "UNABLE TO CONTINUE in ---" + invocations[i]
			log.Fatalln(err)
			return err, nil
		}
		err := workflow.ExecuteActivity(ctx, GsuiteActivity, invocations[i]).Get(ctx, &result)
		result_gsuite += result + time.Now().String() + "\n"
		if err != nil {
			return result, err
		}
	}
	return result_gsuite, nil
}
