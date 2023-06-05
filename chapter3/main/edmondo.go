package main

import (
	"context"
	"github.com/edmondop/uefa-cl/chapter1"
	"github.com/edmondop/uefa-cl/chapter2"
	"github.com/edmondop/uefa-cl/chapter3"
	"go.temporal.io/sdk/client"
	"log"
)

func main() {
	// The client is a heavyweight object that should be created once per process.
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        "2023/2024",
		TaskQueue: "champions-league",
	}
	participants := chapter3.GetParticipants()
	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, chapter2.ChampionsLeague, participants)
	if err != nil {
		log.Fatalln("Unable to execute Champions League", err)
	}
	// Synchronously wait for the workflow completion.
	var result chapter1.Result
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("The UEFA CL 2023/2024 winner is:", result.Winner)
}
