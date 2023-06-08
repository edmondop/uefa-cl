package chapter3

import (
	"context"
	"errors"
	"go.temporal.io/sdk/workflow"
	"time"
)

type KnockoutRoundParticipants struct {
	Pots []Pot
}

type Finalists struct {
	team1 Team
	team2 Team
}

type KnockoutStageMatchResult struct {
	HomeGoals int
	AwayGoals int
}

func PlayKnockoutRoundLeg(_ context.Context, match KnockoutRoundLeg) (KnockoutStageMatchResult, error) {
	return getKnockoutRoundResult(match)
}

func KnockoutRound(ctx workflow.Context, fixture KnockoutRoundFixture) (Team, error) {
	var firstLegResult KnockoutStageMatchResult
	activityOptions := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, activityOptions)
	err := workflow.ExecuteActivity(ctx, PlayKnockoutRoundLeg, fixture.FirstLeg).Get(ctx, &firstLegResult)
	if err != nil {
		return Team{}, err
	}
	var secondLegResult KnockoutStageMatchResult
	err = workflow.ExecuteActivity(ctx, PlayKnockoutRoundLeg, fixture.SecondLeg).Get(ctx, &secondLegResult)
	if err != nil {
		return Team{}, err
	}
	firstTeamGoal := firstLegResult.HomeGoals + secondLegResult.AwayGoals
	secondTeamGoals := firstLegResult.AwayGoals + secondLegResult.HomeGoals
	if firstTeamGoal > secondTeamGoals {
		return fixture.FirstLeg.HomeTeam, nil
	} else {
		return fixture.SecondLeg.HomeTeam, nil // If equal, the results are wrong, it's not possible
	}
	return Team{}, errors.New("since 2021 UEFA has abolished away goals tiebreaker rule, this situation is not possible")
}

func KnockoutStage(ctx workflow.Context, result GroupStageResult) (Finalists, error) {
	roundOf16Partecipants := Participants{
		Pots: result.Pots,
	}
	var activity *KnockoutPhaseDrawingVenue
	var err error
	var nextRoundFixtures []KnockoutRoundFixture
	err = workflow.ExecuteActivity(ctx, activity.DrawFixtures, roundOf16Partecipants).Get(ctx, &nextRoundFixtures)
	if err != nil {
		return Finalists{}, err
	}
	return Finalists{}, nil
}
