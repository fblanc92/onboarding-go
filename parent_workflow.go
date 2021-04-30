package app

import (
	"fmt"

	"go.temporal.io/sdk/workflow"
)

// @@@SNIPSTART onboarding-go-workflow

// SampleParentWorkflow workflow definition
func SampleParentWorkflow(ctx workflow.Context) (string, error) {
	logger := workflow.GetLogger(ctx)

	cwo := workflow.ChildWorkflowOptions{}
	ctx = workflow.WithChildOptions(ctx, cwo)

	gsuite := true
	msft := true

	// var result string
	var result1 string
	var result2 string
	var result3 string
	if gsuite {
		err := workflow.ExecuteChildWorkflow(ctx, GsuiteWorkflow, "gsuite").Get(ctx, &result1)
		if err != nil {
			logger.Error("Parent execution received child execution failure.", "Error", err)
			return result1, err
		}
	}
	if msft {
		err := workflow.ExecuteChildWorkflow(ctx, MsftWorkflow, "msft").Get(ctx, &result2)
		if err != nil {
			logger.Error("Parent execution received child execution failure.", "Error", err)
			return result2, err
		}
	}
	err := workflow.ExecuteActivity(ctx, FinalActivity, "final").Get(ctx, &result3)
	if err != nil {
		return result3, err
	}
	logger.Info("Parent execution completed.", "Result", result1+" "+result2)
	return result1 + " " + result2, nil
}

func SampleActivity(input string) (string, error) {
	name := "sampleActivity"
	fmt.Printf("Run %s with input %v \n", name, input)
	return "Result_" + input, nil
}

func MsftActivity(input string) (string, error) {
	name := "msftActivity"
	fmt.Printf("Run %s with input %v \n", name, input)
	return "Result_" + input, nil
}

func GsuiteActivity(input string) (string, error) {
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
