package main

import (
	"context"
	"fmt"
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

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, chapter4.MilanInterFirstLeg, time.Second)
	if err != nil {
		log.Fatalln(err)
	}
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				response, err := c.QueryWorkflow(context.Background(), we.GetID(), "", "current_result")
				if err != nil {
					done <- true
				}
				var result chapter4.Result
				err = response.Get(&result)
				if err != nil {
					done <- true
				} else {
					fmt.Println("At ", t, " the result was ", result)
				}

			}
		}
	}()
	// Synchronously wait for the workflow completion.
	var result chapter4.Result
	err = we.Get(context.Background(), &result)
	done <- true
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("The result is ", result)
}
