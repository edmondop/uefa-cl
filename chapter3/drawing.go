package chapter3

import (
	"context"
	"errors"
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
	Groups [8]Group
}

type PairingInput struct {
	Participants Participants
}

func (d *GroupStageDrawingVenue) DrawGroups(ctx context.Context, input GroupStageDrawInput) (GroupStageDrawResult, error) {
	logger := activity.GetLogger(ctx)
	msg := fmt.Sprintf("Drawing group stages in %s with a total of %d teams", d.LocationName, input.Participants.TeamCount())
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

func (d *KnockoutPhaseDrawingVenue) DrawFixtures(ctx context.Context, participants Participants) ([]KnockoutRoundFixture, error) {
	logger := activity.GetLogger(ctx)
	msg := fmt.Sprintf("Creating %d pairings in %s.", participants.TeamCount()/2, d.LocationName)

	logger.Info(msg)
	if participants.TeamCount() == 16 {
		return []KnockoutRoundFixture{
			{
				FirstLeg: KnockoutRoundLeg{
					HomeTeam: Team{Name: "AC Milan"},
					AwayTeam: Team{Name: "Tottenham Hotspur"},
				},
				SecondLeg: KnockoutRoundLeg{
					HomeTeam: Team{Name: "Tottenham Hotspur"},
					AwayTeam: Team{Name: "AC Milan"},
				},
			},
			{
				FirstLeg: KnockoutRoundLeg{
					HomeTeam: Team{Name: "Paris Saint-Germain"},
					AwayTeam: Team{Name: "Bayern München"},
				},
				SecondLeg: KnockoutRoundLeg{
					HomeTeam: Team{Name: "Bayern München"},
					AwayTeam: Team{Name: "Paris Saint-Germain"},
				},
			},
			{
				FirstLeg: KnockoutRoundLeg{
					HomeTeam: Team{Name: "Club Brugge KV"},
					AwayTeam: Team{Name: "SL Benfica"},
				},
				SecondLeg: KnockoutRoundLeg{
					HomeTeam: Team{Name: "SL Benfica"},
					AwayTeam: Team{Name: "Club Brugge KV"},
				},
			},
			{
				FirstLeg: KnockoutRoundLeg{
					HomeTeam: Team{Name: "Borussia Dortmund"},
					AwayTeam: Team{Name: "Chelsea FC"},
				},
				SecondLeg: KnockoutRoundLeg{
					HomeTeam: Team{Name: "Chelsea FC"},
					AwayTeam: Team{Name: "Borussia Dortmund"},
				},
			},
			{
				FirstLeg: KnockoutRoundLeg{
					HomeTeam: Team{Name: "Liverpool FC"},
					AwayTeam: Team{Name: "Real Madrid"},
				},
				SecondLeg: KnockoutRoundLeg{
					HomeTeam: Team{Name: "Real Madrid"},
					AwayTeam: Team{Name: "Liverpool FC"},
				},
			},
			{
				FirstLeg: KnockoutRoundLeg{
					HomeTeam: Team{Name: "Eintracht Frankfurt"},
					AwayTeam: Team{Name: "SSC Napoli"},
				},
				SecondLeg: KnockoutRoundLeg{
					HomeTeam: Team{Name: "SSC Napoli"},
					AwayTeam: Team{Name: "Eintracht Frankfurt"},
				},
			},
			{
				FirstLeg: KnockoutRoundLeg{
					HomeTeam: Team{Name: "RB Leipzig"},
					AwayTeam: Team{Name: "Manchester City"},
				},
				SecondLeg: KnockoutRoundLeg{
					HomeTeam: Team{Name: "Manchester City"},
					AwayTeam: Team{Name: "RB Leipzig"},
				},
			},
			{
				FirstLeg: KnockoutRoundLeg{
					HomeTeam: Team{Name: "FC Internazionale"},
					AwayTeam: Team{Name: "FC Porto"},
				},
				SecondLeg: KnockoutRoundLeg{
					HomeTeam: Team{Name: "FC Porto"},
					AwayTeam: Team{Name: "FC Internazionale"},
				},
			},
		}, nil
	}
	if participants.TeamCount() == 8 {
		return []KnockoutRoundFixture{
			{
				FirstLeg: KnockoutRoundLeg{
					HomeTeam: Team{Name: "SL Benfica"},
					AwayTeam: Team{Name: "FC Internazionale"},
				},
				SecondLeg: KnockoutRoundLeg{
					HomeTeam: Team{Name: "FC Internazionale"},
					AwayTeam: Team{Name: "SL Benfica"},
				},
			},
			{
				FirstLeg: KnockoutRoundLeg{
					HomeTeam: Team{Name: "Manchester City"},
					AwayTeam: Team{Name: "Bayern München"},
				},
				SecondLeg: KnockoutRoundLeg{
					HomeTeam: Team{Name: "Bayern München"},
					AwayTeam: Team{Name: "Manchester City"},
				},
			},
			{
				FirstLeg: KnockoutRoundLeg{
					HomeTeam: Team{Name: "Real Madrid"},
					AwayTeam: Team{Name: "Chelsea FC"},
				},
				SecondLeg: KnockoutRoundLeg{
					HomeTeam: Team{Name: "Chelsea FC"},
					AwayTeam: Team{Name: "Real Madrid"},
				},
			},
			{
				FirstLeg: KnockoutRoundLeg{
					HomeTeam: Team{Name: "AC Milan"},
					AwayTeam: Team{Name: "SSC Napoli"},
				},
				SecondLeg: KnockoutRoundLeg{
					HomeTeam: Team{Name: "SSC Napoli"},
					AwayTeam: Team{Name: "AC Milan"},
				},
			},
		}, nil
	}
	if participants.TeamCount() == 4 {
		return []KnockoutRoundFixture{
			{
				FirstLeg: KnockoutRoundLeg{
					HomeTeam: Team{Name: "Real Madrid"},
					AwayTeam: Team{Name: "Manchester City"},
				},
				SecondLeg: KnockoutRoundLeg{
					HomeTeam: Team{Name: "Manchester City"},
					AwayTeam: Team{Name: "Real Madrid"},
				},
			},
			{
				FirstLeg: KnockoutRoundLeg{
					HomeTeam: Team{Name: "AC Milan"},
					AwayTeam: Team{Name: "FC Internazionale"},
				},
				SecondLeg: KnockoutRoundLeg{
					HomeTeam: Team{Name: "FC Internazionale"},
					AwayTeam: Team{Name: "AC Milan"},
				},
			},
		}, nil
	}

	return []KnockoutRoundFixture{}, errors.New("invalid number of team in participants")
}
