package chapter3

import (
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
	"testing"
)

func TestInterPorto(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()
	env.RegisterActivity(PlayKnockoutRoundLeg)
	env.ExecuteWorkflow(KnockoutRound, KnockoutRoundFixture{
		FirstLeg: KnockoutRoundLeg{
			HomeTeam: Team{"FC Internazionale"},
			AwayTeam: Team{"FC Porto"},
		},
		SecondLeg: KnockoutRoundLeg{
			HomeTeam: Team{"FC Porto"},
			AwayTeam: Team{"FC Internazionale"},
		},
	})
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())
	var qualifiedTeam Team
	require.NoError(t, env.GetWorkflowResult(&qualifiedTeam))
	require.Equal(t, Team{Name: "FC Internazionale"}, qualifiedTeam)
}

func TestInterAcMilan(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()
	env.RegisterActivity(PlayKnockoutRoundLeg)
	env.ExecuteWorkflow(KnockoutRound, KnockoutRoundFixture{
		FirstLeg: KnockoutRoundLeg{
			HomeTeam: Team{"AC Milan"},
			AwayTeam: Team{"FC Internazionale"},
		},
		SecondLeg: KnockoutRoundLeg{
			HomeTeam: Team{"FC Internazionale"},
			AwayTeam: Team{"AC Milan"},
		},
	})
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())
	var qualifiedTeam Team
	require.NoError(t, env.GetWorkflowResult(&qualifiedTeam))
	require.Equal(t, Team{Name: "FC Internazionale"}, qualifiedTeam)
}
