package chapter3

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
	"testing"
	"time"
)

func assertGroupHasRightResult(t *testing.T, group Group, winnerTeam string, runnerUp string) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()
	env.SetTestTimeout(60 * time.Hour)
	env.OnActivity(PlayMatch, mock.Anything, mock.Anything).Return(
		func(ctx context.Context, match GroupStageMatch) (GroupStageMatchResult, error) {
			result, error := getResult(match)
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
	assertGroupHasRightResult(t, GroupC(), "Bayern München", "Inter")
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
