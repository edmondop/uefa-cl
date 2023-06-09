package main

import (
	"github.com/edmondop/uefa-cl/chapter3"
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
		Identity: "Champions League Team",
	})
	w.RegisterWorkflow(chapter3.GroupStage)
	w.RegisterWorkflow(chapter3.PlayGroupMatches)
	w.RegisterWorkflow(chapter3.KnockoutStage)
	w.RegisterWorkflow(chapter3.PlayKnockoutFixture)
	w.RegisterWorkflow(chapter3.ChampionsLeague)
	w.RegisterActivity(chapter3.PlayKnockoutRoundLeg)
	w.RegisterActivity(chapter3.PlayGroupStageMatch)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start management_team", err)
	}
}
