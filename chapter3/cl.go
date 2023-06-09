package chapter3

import (
	"go.temporal.io/sdk/workflow"
	"time"
)

func ChampionsLeague(ctx workflow.Context, participants Participants) (Result, error) {
	groupStageDrawResult, err := drawGroupStageGroups(ctx, participants)
	if err != nil {
		return Result{}, err
	}
	var groupStageResult GroupStageResult
	err = workflow.ExecuteChildWorkflow(ctx, GroupStage, groupStageDrawResult).Get(ctx, &groupStageResult)
	if err != nil {
		return Result{}, err
	}
	var finalist Finalists
	err = workflow.ExecuteChildWorkflow(ctx, KnockoutStage, groupStageResult).Get(ctx, &finalist)
	if err != nil {
		return Result{}, err
	}
	return Result{
		Winner: Team{Name: "FC Internazionale"},
	}, nil
}

func drawGroupStageGroups(ctx workflow.Context, participants Participants) (GroupStageDrawResult, error) {
	ctx = workflow.WithActivityOptions(
		ctx,
		workflow.ActivityOptions{
			StartToCloseTimeout: 10 * time.Hour,
		},
	)
	var groupStageDrawingVenue *GroupStageDrawingVenue
	groupStageDraws := workflow.ExecuteActivity(ctx, groupStageDrawingVenue.DrawGroups, GroupStageDrawInput{Participants: participants})
	var groupStageDrawResult GroupStageDrawResult
	err := groupStageDraws.Get(ctx, &groupStageDrawResult)
	return groupStageDrawResult, err
}
