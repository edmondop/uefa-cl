package chapter3

import (
	"go.temporal.io/sdk/workflow"
	"time"
)

func ChampionsLeague(ctx workflow.Context, participants Participants) (Result, error) {
	var groupStageDrawingVenue *GroupStageDrawingVenue
	activityOptions := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Hour,
	}
	ctx = workflow.WithActivityOptions(ctx, activityOptions)
	groupStageDraws := workflow.ExecuteActivity(ctx, groupStageDrawingVenue.DrawGroups, GroupStageDrawInput{Participants: participants})
	var groupStageDrawResult GroupStageDrawResult
	err := groupStageDraws.Get(ctx, &groupStageDrawResult)
	if err != nil {
		return Result{}, err
	}
	var groupStageResult GroupStageResult
	err = workflow.ExecuteChildWorkflow(ctx, GroupStage, groupStageDrawResult).Get(ctx, &groupStageResult)
	if err != nil {
		return Result{}, err
	}
	// Playing group stage
	workflow.Sleep(ctx, time.Second*60)

	// Pairing for knockout phase

	return Result{
		Winner: Team{Name: "FC Internazionale"},
	}, nil
}

//func ChampionsLeague(ctx workflow.Context, participants Participants) (Result, error) {
