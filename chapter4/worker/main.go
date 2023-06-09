package main

import (
	"github.com/edmondop/uefa-cl/chapter4"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// The client and management_team are heavyweight objects that should be created once per process.
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "my-queue", worker.Options{})

	w.RegisterWorkflow(chapter4.MilanInterFirstLeg)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln(err)
	}
}
