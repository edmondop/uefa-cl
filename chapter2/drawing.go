package chapter2

import (
	"context"
	"fmt"
	"go.temporal.io/sdk/activity"
)

type GroupStageDrawingVenue struct {
	LocationName string
}

type KnockoutPhaseDrawingVenue struct {
	LocationName string
}

type GroupStageDrawInput struct {
	Participants Participants
}

type GroupStageDrawResult struct {
	Group [8]Group
}

func (d *GroupStageDrawingVenue) DrawGroups(ctx context.Context, input GroupStageDrawInput) (GroupStageDrawResult, error) {
	logger := activity.GetLogger(ctx)
	msg := fmt.Sprintf("Drawing group stages in %s with a total of %d teams", d.LocationName, input.Participants.TeamCount())
	logger.Info(msg)
	return GroupStageDrawResult{}, nil
}

func (d *KnockoutPhaseDrawingVenue) DrawFixtures(ctx context.Context, participants Participants) ([]Fixture, error) {
	logger := activity.GetLogger(ctx)
	msg := fmt.Sprintf("Creating %d pairings in %s.", participants.TeamCount()/2, d.LocationName)
	logger.Info(msg)
	return []Fixture{}, nil
}
