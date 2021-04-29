package app

// @@@SNIPSTART onboarding-go-shared-task-queue
const TransferMoneyTaskQueue = "TRANSFER_MONEY_TASK_QUEUE"
const OnboardingTaskQueue = "ONBOARDING_TASK_QUEUE"

// @@@SNIPEND

type TransferDetails struct {
	Amount      float32
	FromAccount string
	ToAccount   string
	ReferenceID string
}

type OnBoardingTask struct {
	TaskID       int // fix for each onboarding person
	FromDivision string
	ToDivision   string
	TrackingID   int // supposed to increment when passing from a division to another "creating its accounts"
}
