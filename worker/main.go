package main

import (
	"log"
	"onboarding-go/app"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

// @@@SNIPSTART onboarding-go-worker
func main() {
	// The client and worker are heavyweight objects that should be created once per process.
	c, err := client.NewClient(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "child-workflow", worker.Options{})

	w.RegisterWorkflow(app.SampleParentWorkflow)
	w.RegisterWorkflow(app.SampleChildWorkflow)
	w.RegisterActivity(app.SampleActivity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}

// @@@SNIPEND
