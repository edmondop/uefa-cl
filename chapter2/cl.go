package chapter2

import (
	"go.temporal.io/sdk/workflow"
	"time"
)

func ChampionsLeague(ctx workflow.Context, participants Participants) (Result, error) {
	workflow.Sleep(ctx, time.Second*60)
	return Result{
		Winner: Team{Name: "FC Internazionale"},
	}, nil
}
