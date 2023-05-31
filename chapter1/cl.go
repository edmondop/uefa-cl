package chapter1

import "go.temporal.io/sdk/workflow"

func ChampionsLeague(ctx workflow.Context, participants Participants) (Result, error) {
	return Result{
		Winner: Team{Name: "FC Internazionale"},
	}, nil
}
