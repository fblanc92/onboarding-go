package app

import (
	"context"
	"fmt"
)

// TRANSFER MONEY
// func Withdraw(ctx context.Context, transferDetails TransferDetails) error {
// 	fmt.Printf(
// 		"\nWithdrawing $%f from account %s. ReferenceId: %s\n",
// 		transferDetails.Amount,
// 		transferDetails.FromAccount,
// 		transferDetails.ReferenceID,
// 	)
// 	return nil
// }

// @@@SNIPSTART onboarding-go-activity
// func Deposit(ctx context.Context, transferDetails TransferDetails) error {
// 	fmt.Printf(
// 		"\nDepositing $%f into account %s. ReferenceId: %s\n",
// 		transferDetails.Amount,
// 		transferDetails.ToAccount,
// 		transferDetails.ReferenceID,
// 	)
// 	// Switch out comments on the return statements to simulate an error
// 	//return fmt.Errorf("deposit did not occur due to an issue")
// 	return nil
// }

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
