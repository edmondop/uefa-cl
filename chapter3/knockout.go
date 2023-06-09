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
	Team1 Team
	Team2 Team
}

type KnockoutStageMatchResult struct {
	HomeGoals int
	AwayGoals int
}

func PlayKnockoutRoundLeg(_ context.Context, match KnockoutRoundLeg) (KnockoutStageMatchResult, error) {
	return getKnockoutRoundResult(match)
}

func PlayKnockoutFixture(ctx workflow.Context, fixture KnockoutRoundFixture) (Team, error) {
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

// Returning a Pot array doesn't probably make sense but I didn't have enough time
func playKnockoutRound(ctx workflow.Context, fixtures []KnockoutRoundFixture) ([]Team, error) {
	var winners []Team
	for _, fixture := range fixtures {
		var winner Team
		err := workflow.ExecuteChildWorkflow(ctx, PlayKnockoutFixture, fixture).Get(ctx, &winner)
		if err != nil {
			return winners, nil
		}
		winners = append(winners, winner)
	}
	return winners, nil
}

func distributeInPots(teams []Team, potCount int) []Pot {
	var pots []Pot
	for i := 0; i < potCount; i++ {
		pots = append(pots, Pot{})
	}
	for index, team := range teams {
		pots[index%potCount].Teams = append(pots[index%potCount].Teams, team)
	}
	return pots
}

func KnockoutStage(ctx workflow.Context, result GroupStageResult) (Finalists, error) {
	roundOf16Partecipants := Participants(result)
	activityOptions := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, activityOptions)
	var activity *KnockoutPhaseDrawingVenue
	var err error
	var roundOf16Fixtures []KnockoutRoundFixture
	err = workflow.ExecuteActivity(ctx, activity.DrawFixtures, roundOf16Partecipants).Get(ctx, &roundOf16Fixtures)
	if err != nil {
		return Finalists{}, err
	}
	round16Winners, err := playKnockoutRound(ctx, roundOf16Fixtures)
	if err != nil {
		return Finalists{}, err
	}
	roundOf8Participants := distributeInPots(round16Winners, 4)
	var roundOf8Fixtures []KnockoutRoundFixture
	err = workflow.ExecuteActivity(ctx, activity.DrawFixtures, Participants{
		Pots: roundOf8Participants,
	}).Get(ctx, &roundOf8Fixtures)
	if err != nil {
		return Finalists{}, err
	}
	round8Winners, err := playKnockoutRound(ctx, roundOf8Fixtures)
	if err != nil {
		return Finalists{}, err
	}
	roundOf4Participants := distributeInPots(round8Winners, 2)
	var roundOf4Fixtures []KnockoutRoundFixture
	err = workflow.ExecuteActivity(ctx, activity.DrawFixtures,
		Participants{
			Pots: roundOf4Participants,
		},
	).Get(ctx, &roundOf4Fixtures)
	if err != nil {
		return Finalists{}, err
	}
	roundOf4Winners, err := playKnockoutRound(ctx, roundOf4Fixtures)
	if err != nil {
		return Finalists{}, err
	}
	return Finalists{Team1: roundOf4Winners[0], Team2: roundOf4Winners[1]}, nil
}
