package chapter3

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
	"sort"
	"testing"
	"time"
)

func assertGroupHasRightResult(t *testing.T, group Group, winnerTeam string, runnerUp string) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()
	env.SetTestTimeout(60 * time.Hour)
	env.OnActivity(PlayGroupStageMatch, mock.Anything, mock.Anything).Return(
		func(ctx context.Context, match GroupStageMatch) (GroupStageMatchResult, error) {
			result, error := getStageMatchResult(match)
			require.NoError(t, error)
			return result, error
		})
	env.ExecuteWorkflow(GroupWorkflow, group)

	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())
	var ranking GroupRanking
	require.NoError(t, env.GetWorkflowResult(&ranking))
	require.Equal(t, Team{Name: winnerTeam}, ranking.GroupRankingEntries[0].Team)
	require.Equal(t, Team{Name: runnerUp}, ranking.GroupRankingEntries[1].Team)
}
func TestGroupA(t *testing.T) {
	assertGroupHasRightResult(t, GroupA(), "SSC Napoli", "Liverpool FC")
}

func TestGroupB(t *testing.T) {
	assertGroupHasRightResult(t, GroupB(), "FC Porto", "Club Brugge KV")
}

func TestGroupC(t *testing.T) {
	assertGroupHasRightResult(t, GroupC(), "Bayern München", "FC Internazionale")
}

func TestGroupD(t *testing.T) {
	assertGroupHasRightResult(t, GroupD(), "Tottenham Hotspur", "Eintracht Frankfurt")
}

func TestGroupE(t *testing.T) {
	assertGroupHasRightResult(t, GroupE(), "Chelsea FC", "AC Milan")
}

func TestGroupF(t *testing.T) {
	assertGroupHasRightResult(t, GroupF(), "Real Madrid", "RB Leipzig")
}

func TestGroupG(t *testing.T) {
	assertGroupHasRightResult(t, GroupG(), "Manchester City", "Borussia Dortmund")
}

func TestGroupH(t *testing.T) {
	assertGroupHasRightResult(t, GroupH(), "SL Benfica", "Paris Saint-Germain")
}

func sortTeamsInPot(pot *Pot) {
	sort.Slice(pot.Teams, func(i, j int) bool {
		return pot.Teams[i].Name < pot.Teams[j].Name
	})
}

func TestGroupStage(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()
	env.SetTestTimeout(60 * time.Hour)
	env.OnActivity(PlayGroupStageMatch, mock.Anything, mock.Anything).Return(
		func(ctx context.Context, match GroupStageMatch) (GroupStageMatchResult, error) {
			result, error := getStageMatchResult(match)
			require.NoError(t, error)
			return result, error
		})
	groupStageDrawResult := GroupStageDrawResult{
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
	}
	env.RegisterWorkflow(GroupWorkflow)
	env.ExecuteWorkflow(GroupStage, groupStageDrawResult)

	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())
	var groupStageResult GroupStageResult
	require.NoError(t, env.GetWorkflowResult(&groupStageResult))
	winnerPot := Pot{
		Teams: []Team{
			{Name: "SSC Napoli"},
			{Name: "FC Porto"},
			{Name: "Bayern München"},
			{Name: "Tottenham Hotspur"},
			{Name: "Chelsea FC"},
			{Name: "Real Madrid"},
			{Name: "Manchester City"},
			{Name: "SL Benfica"},
		},
	}
	runnerUpPot := Pot{
		Teams: []Team{
			{Name: "Liverpool FC"},
			{Name: "Club Brugge KV"},
			{Name: "FC Internazionale"},
			{Name: "Eintracht Frankfurt"},
			{Name: "AC Milan"},
			{Name: "RB Leipzig"},
			{Name: "Borussia Dortmund"},
			{Name: "Paris Saint-Germain"},
		},
	}
	// Handling concurrency avoiding problems
	sortTeamsInPot(&winnerPot)
	sortTeamsInPot(&runnerUpPot)
	sortTeamsInPot(&groupStageResult.Pots[0])
	sortTeamsInPot(&groupStageResult.Pots[1])
	require.Equal(t, winnerPot, groupStageResult.Pots[0])
	require.Equal(t, runnerUpPot, groupStageResult.Pots[1])

}
