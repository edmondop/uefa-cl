# The end of the adventure (maybe)

I wish I would be able to write other code using Temporal in this example, but there are just too many useful features!
So I only have time for two left: satisfying an impatient supporter and bribing the referee.

## Satisfying an impatient supporter

Yes, you can wait the end of a workflow to wait its outcome, like the winner of a match. But can you really wait so
long without knowing what's going on? Personally, I can't. I am the kind of supporter that is either watching the match
or checking his phone to see if there are updates.

Meet **Queries**, which allow applications to observe the workflow state from the outside. In order to support Queries,
the first thing we need to do is register a **query handler** on the workflow. Note that the Query Handler can also take
argument, but we won't 

```go
func MilanInterFirstLeg(ctx workflow.Context, d time.Duration) (Result, error) {

    currentResult := Result{
        Team1: Team{"AC Milan"} ,
        Team2: Team{"FC Internazionale"},
    }
    queryType := "current_result"
    err := workflow.SetQueryHandler(ctx, queryType, func () (Result, error) {
        return currentResult, nil
    })
	// Other code
}
```

This is not magic, or maybe it is. Temporal can serialize that struct to clients, and if hold in that struct the state
of your workflow, effectively you can expose it to the outside.

Our entire workflow is below. I am using `workflow.Sleep` to simulate the time between the goals, and updating the state
when Edin Džeko and Henrikh Mkhitaryan scores. 

```go
func MilanInterFirstLeg(ctx workflow.Context, d time.Duration) (Result, error) {

	currentResult := Result{
		Team1: Team{"AC Milan"},
		Team2: Team{"FC Internazionale"},
	}
	queryType := "current_result"
	err := workflow.SetQueryHandler(ctx, queryType, func() (Result, error) {
		return currentResult, nil
	})
	if err != nil {
		return Result{}, err
	}
	workflow.Sleep(ctx, 8*d)
	currentResult.Team2Goals += 1
	workflow.Sleep(ctx, 3*d)
	currentResult.Team2Goals += 1
	workflow.Sleep(ctx, (90-8-3)*d)
	return currentResult, nil

}

```

Now, if our `main.go` uses the Temporal `Client` to issues periodic queries to the running workflows, and end up producing
the following output. As you see, the result on my smartphone is updated and I 

```
2023/06/09 15:41:36 INFO  No logger configured for temporal client. Created default one.
At  2023-06-09 15:41:37.05814 -0700 PDT m=+1.012341376  the result was  {{AC Milan} {FC Internazionale} 0 0}
At  2023-06-09 15:41:38.057457 -0700 PDT m=+2.011628917  the result was  {{AC Milan} {FC Internazionale} 0 0}
At  2023-06-09 15:41:39.058338 -0700 PDT m=+3.012479751  the result was  {{AC Milan} {FC Internazionale} 0 0}
At  2023-06-09 15:41:40.058191 -0700 PDT m=+4.012302501  the result was  {{AC Milan} {FC Internazionale} 0 0}
At  2023-06-09 15:41:41.057582 -0700 PDT m=+5.011663501  the result was  {{AC Milan} {FC Internazionale} 0 0}
At  2023-06-09 15:41:42.057472 -0700 PDT m=+6.011523917  the result was  {{AC Milan} {FC Internazionale} 0 0}
At  2023-06-09 15:41:43.057835 -0700 PDT m=+7.011856584  the result was  {{AC Milan} {FC Internazionale} 0 0}
At  2023-06-09 15:41:44.058346 -0700 PDT m=+8.012337876  the result was  {{AC Milan} {FC Internazionale} 0 0}
At  2023-06-09 15:41:45.05839 -0700 PDT m=+9.012351584  the result was  {{AC Milan} {FC Internazionale} 0 1}
At  2023-06-09 15:41:46.058167 -0700 PDT m=+10.012098709  the result was  {{AC Milan} {FC Internazionale} 0 1}
At  2023-06-09 15:41:47.058437 -0700 PDT m=+11.012338417  the result was  {{AC Milan} {FC Internazionale} 0 1}
At  2023-06-09 15:41:48.057805 -0700 PDT m=+12.011677292  the result was  {{AC Milan} {FC Internazionale} 0 2}
At  2023-06-09 15:41:49.058498 -0700 PDT m=+13.012339751  the result was  {{AC Milan} {FC Internazionale} 0 2}
```

## Bribing the referee

`Query` help us retrieving the state from the outside, what about signaling from the outside that something has changed?
For example, how would you inform a referee that he has just received a wire transfer on his bank account, and he should
keep into account for the rest of the match?

Meet `Signals`: the way you can have workflows listen to outside state. This other workflow should be clear: you can
listen to a channel, and according to what you receive on the channel, update the state, which you can use to base 
decision later

```go
func MilanInterSecondLeg(ctx workflow.Context, d time.Duration) (Result, error) {
	currentResult := Result{
		Team1: Team{"FC Internazionale"},
		Team2: Team{"AC Milan"},
	}
	bribed := false
	log := workflow.GetLogger(ctx)
	selector := workflow.NewSelector(ctx)
	selector.AddReceive(workflow.GetSignalChannel(ctx, "bribing-channel"), func(c workflow.ReceiveChannel, more bool) {
		c.Receive(ctx, nil)
		bribed = true
		log.Info("Bribe received")
	})
	workflow.Sleep(ctx, 74*d)
	currentResult.Team1Goals += 1
	if bribed {
		log.Info("Sorry, it was clearly offside!")
		currentResult.Team1Goals -= 1
	}
	workflow.Sleep(ctx, (90-74)*d)
	return currentResult, nil

}
```

