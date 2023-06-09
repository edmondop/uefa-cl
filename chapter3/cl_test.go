package chapter3

import (
	"context"
	"github.com/stretchr/testify/mock"
	"go.temporal.io/sdk/workflow"
	"testing"

	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func verifyFcInternazionaleIsWinner(t *testing.T, env *testsuite.TestWorkflowEnvironment) {
	participants := GetParticipants()
	env.ExecuteWorkflow(ChampionsLeague, participants)

	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())
	var result Result
	require.NoError(t, env.GetWorkflowResult(&result))
	require.Equal(t, Team{Name: "FC Internazionale"}, result.Winner)
}

func TestChampionsLeagueIsolation(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()
	var activity *GroupStageDrawingVenue
	// Mock activity implementation
	env.OnActivity(activity.DrawGroups, mock.Anything, mock.Anything).Return(
		func(ctx context.Context, input GroupStageDrawInput) (GroupStageDrawResult, error) {
			require.Equal(t, input.Participants.TeamCount(), 32)
			return GroupStageDrawResult{}, nil
		})
	env.RegisterWorkflow(GroupStage)
	env.RegisterWorkflow(KnockoutStage)
	env.OnWorkflow(GroupStage, mock.Anything, mock.Anything).Return(
		func(ctx workflow.Context, result GroupStageDrawResult) (GroupStageResult, error) {
			return GroupStageResult{}, nil
		},
	)
	env.OnWorkflow(KnockoutStage, mock.Anything, mock.Anything).Return(
		func(ctx workflow.Context, result GroupStageResult) (Finalists, error) {
			return Finalists{Team1: Team{Name: "Manchester City"}, Team2: Team{Name: "FC Internazionale"}}, nil
		})
	verifyFcInternazionaleIsWinner(t, env)
}

func TestChampionsLeagueIntegration(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()
	var groupDrawingVenue GroupStageDrawingVenue = GroupStageDrawingVenue{
		LocationName: "Monte Carlo",
	}
	// Mock activity implementation
	env.RegisterActivity(groupDrawingVenue.DrawGroups)
	var knockoutDrawingVenue KnockoutPhaseDrawingVenue = KnockoutPhaseDrawingVenue{
		LocationName: "Nyon",
	}
	// Mock activity implementation
	env.RegisterActivity(knockoutDrawingVenue.DrawFixtures)
	env.RegisterActivity(PlayKnockoutRoundLeg)
	env.RegisterActivity(PlayGroupStageMatch)
	env.RegisterWorkflow(GroupStage)
	env.RegisterWorkflow(PlayGroupMatches)
	env.RegisterWorkflow(KnockoutStage)
	env.RegisterWorkflow(PlayKnockoutFixture)

	verifyFcInternazionaleIsWinner(t, env)
}
