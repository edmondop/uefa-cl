package main

import (
	"context"
	"github.com/edmondop/uefa-cl/chapter4"
	"go.temporal.io/sdk/client"
	"log"
	"time"
)

func main() {
	// The client is a heavyweight object that should be created once per process.
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		TaskQueue: "my-queue",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, chapter4.MilanInterSecondLeg, time.Second)
	if err != nil {
		log.Fatalln(err)
	}
	err = c.SignalWorkflow(context.Background(), we.GetID(), we.GetRunID(), "bribing-channel", "")
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	// Synchronously wait for the workflow completion.
	var result chapter4.Result
	err = we.Get(context.Background(), &result)

	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("The result is ", result)
}
