package chapter1

import "go.temporal.io/sdk/workflow"

func ChampionsLeague(ctx workflow.Context, participants Participants) (error, Result) {
	return nil, Result{
		Winner: Team{Name: "FC Internazionale"},
	}
}
