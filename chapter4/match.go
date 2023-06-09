package chapter4

import (
	"go.temporal.io/sdk/workflow"
	"time"
)

type Team struct {
	Name string
}

type Result struct {
	Team1      Team
	Team2      Team
	Team1Goals int
	Team2Goals int
}

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
