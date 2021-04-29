package app

import (
	"context"
	"fmt"
)

//ONBOARDING
func GotoNextDivision(ctx context.Context, onboardingDetails OnBoardingTask) error {
	fmt.Printf(
		"\nOnboarding number: %d from division %s to division %s. TRACKID: %d\n",
		onboardingDetails.TaskID,
		onboardingDetails.FromDivision,
		onboardingDetails.ToDivision,
		onboardingDetails.TrackingID,
	)
	// return fmt.Errorf("Goto Next division ERROR")
	return nil
}

// @@@SNIPEND"
