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

// SampleParentWorkflow workflow definition
func SampleParentWorkflow(ctx workflow.Context) (string, error) {
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
	// childs
	cwo := workflow.ChildWorkflowOptions{}
	ctx = workflow.WithChildOptions(ctx, cwo)
	// triggers
	gsuite := true
	msft := true

	// child workflows results
	var result1 string
	var result2 string
	// parent workflow activity result
	var result3 string
	if gsuite {
		err := workflow.ExecuteChildWorkflow(ctx, GsuiteWorkflow, "gsuite").Get(ctx, &result1)
		result1 += " " + time.Now().String()
		if err != nil {
			logger.Error("Parent execution received child execution failure.", "Error", err)
			return result1, err
		}
	}
	if msft {
		err := workflow.ExecuteChildWorkflow(ctx, MsftWorkflow, "msft").Get(ctx, &result2)
		result2 += " " + time.Now().String()
		if err != nil {
			logger.Error("Parent execution received child execution failure.", "Error", err)
			return result2, err
		}
	}

	ctx = workflow.WithActivityOptions(ctx, options)
	err := workflow.ExecuteActivity(ctx, FinalActivity, "final").Get(ctx, &result3)
	result3 += " " + time.Now().String()
	if err != nil {
		return result3, err
	}
	logger.Info("Parent execution completed.", "Result", result1+" "+result2)

	return "\n.....Result......\n" + "Child Workflow GSUITE:\n" + result1 + "\nChild Workflow MSFT:\t\t" + result2 + "\nParent Workflow Activity:\t" + result3 + "\n", nil
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

func FinalActivity(input string) (string, error) {
	name := "finalActivity"
	fmt.Printf("Run %s with input %v \n", name, input)
	return "Result_" + input, nil
}

// @@@SNIPEND
