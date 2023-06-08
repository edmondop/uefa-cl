package main

import (
	"github.com/edmondop/uefa-cl/chapter2"
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

	w := worker.New(c, "champions-league", worker.Options{
		Identity: "Monte Carlo",
	})

	activities := &chapter2.GroupStageDrawingVenue{
		LocationName: "Monte Carlo",
	}
	w.RegisterActivity(activities)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start Monte Carlo", err)
	}
}
