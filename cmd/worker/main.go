// cmd/worker/main.go
package main

import (
	"log"

	"BAPE/internal/workflow"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, workflow.BapeTaskQueue, worker.Options{})

	w.RegisterWorkflow(workflow.BapeWorkflow)
	w.RegisterActivity(workflow.BapeActivity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
