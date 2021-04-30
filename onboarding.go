package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// @@@SNIPSTART onboarding-go-workflow

// OnboardingWorkflow workflow definition
func OnboardingWorkflow(ctx workflow.Context) (string, error) {
	// workflow activity config
	retrypolicy := &temporal.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2.0,
		MaximumInterval:    time.Minute,
		MaximumAttempts:    500,
	}
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
		RetryPolicy:         retrypolicy,
	}

	//logger
	logger := workflow.GetLogger(ctx)
	// triggers
	gsuite := true
	msft := true

	// child workflows results
	var result_gsuite string
	var result_msft string
	// parent workflow activity result
	var result_mail string
	// gsuite
	cwo := workflow.ChildWorkflowOptions{WorkflowID: "gsuite_process"}
	ctx = workflow.WithChildOptions(ctx, cwo)
	if gsuite {
		err := workflow.ExecuteChildWorkflow(ctx, GsuiteWorkflow, "gsuite").Get(ctx, &result_gsuite)
		// result_gsuite += " " + time.Now().String()
		if err != nil {
			logger.Error("Parent execution received child execution failure in gsuite.", "Error", err)
			return result_gsuite, err
		}
	}
	// gsuite
	cwo = workflow.ChildWorkflowOptions{WorkflowID: "msft_process"}
	ctx = workflow.WithChildOptions(ctx, cwo)
	if msft {
		err := workflow.ExecuteChildWorkflow(ctx, MsftWorkflow, "msft").Get(ctx, &result_msft)
		// result_msft += " " + time.Now().String()
		if err != nil {
			logger.Error("Parent execution received child execution failure in msft.", "Error", err)
			return result_msft, err
		}
	}

	ctx = workflow.WithActivityOptions(ctx, options)
	err := workflow.ExecuteActivity(ctx, SendMailActivity, "Onboarding process finished. A Welcome Email has been sent.").Get(ctx, &result_mail)
	result_mail += " - " + time.Now().Format("2006-01-02 15:04:05")
	if err != nil {
		return result_mail, err
	}
	logger.Info("\nParent execution completed.", "Result", result_gsuite+" "+result_msft)

	return "\n\n.....Result......" + "\n\nChild Workflow GSUITE:" + result_gsuite + "\n\nChild Workflow MSFT:" + result_msft + "\n\nSend Mail Activity:\n" + result_mail + "\n", nil
}

func MsftActivity(input string) (string, error) {
	name := "msftActivity"
	fmt.Printf("Run %s with input %v \n", name, input)
	return "Result_" + input, nil
}

func GsuiteActivity(input string) (string, error) {
	if input == "Gsuite_Create New Password" {
		content, err := ioutil.ReadFile("GSUITE_PASS")
		if err != nil {
			log.Fatal(err)
		}
		if string(content) != "True" {
			return os.Getenv("GSUITE_PASS"), fmt.Errorf("\n\nERROR -> Gsuite_Create New Password\n\n")
		}
	}
	name := "gsuiteActivity"
	fmt.Printf("Run %s with input %v \n", name, input)
	return "Result_" + input, nil
}

func SendMailActivity(input string) (string, error) {
	name := "sendMailActivity"
	fmt.Printf("Run %s with input %v \n", name, input)
	return "Result_" + input, nil
}

// @@@SNIPEND
