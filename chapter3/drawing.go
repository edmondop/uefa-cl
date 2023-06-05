package chapter3

import (
	"context"
	"fmt"
	"go.temporal.io/sdk/activity"
)

type GroupStageDrawing struct {
	Name string
}

type KnockoutPhasePairing struct {
	Name string
}

type GroupStageDrawInput struct {
	Participants Participants
}

type GroupStageDrawResult struct {
	Groups [8]Group
}

type PairingInput struct {
	Participants Participants
}

type PairingResult struct {
	Pairings []Pairing
}

func (d *GroupStageDrawing) DrawGroups(ctx context.Context, input GroupStageDrawInput) (GroupStageDrawResult, error) {
	logger := activity.GetLogger(ctx)
	msg := fmt.Sprintf("Drawing group stages in %s with a total of %d teams", d.Name, input.Participants.TeamCount())
	logger.Info(msg)
	return GroupStageDrawResult{
		Groups: [8]Group{
			GroupA(),
			GroupB(),
			GroupC(),
			GroupD(),
			GroupE(),
			GroupF(),
			GroupG(),
			GroupH(),
		},
	}, nil
}

func (d *KnockoutPhasePairing) DrawPairings(ctx context.Context, pairingInput PairingInput) (PairingResult, error) {
	logger := activity.GetLogger(ctx)
	msg := fmt.Sprintf("Creating %d pairings in %s.", pairingInput.Participants.TeamCount()/2, d.Name)
	logger.Info(msg)
	return PairingResult{}, nil
}
