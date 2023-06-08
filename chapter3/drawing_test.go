package chapter3

import (
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
	"testing"
)

func TestGroupStageDraw(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestActivityEnvironment()
	var drawing = &GroupStageDrawingVenue{LocationName: "Monte Carlo"}
	env.RegisterActivity(drawing.DrawGroups)
	groups, err := env.ExecuteActivity(drawing.DrawGroups, GroupStageDrawInput{
		Participants: GetParticipants(),
	})
	require.NoError(t, err)
	var res *GroupStageDrawResult
	require.NoError(t, groups.Get(&res))
	require.Equal(t, 8, len(res.Groups))

}
func assertKnockoutPhaseDrawFixturesWorkCorrectly(t *testing.T, participants Participants) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestActivityEnvironment()
	var drawing = &KnockoutPhaseDrawingVenue{LocationName: "Nyon"}
	env.RegisterActivity(drawing.DrawFixtures)
	groups, err := env.ExecuteActivity(drawing.DrawFixtures, participants)
	require.NoError(t, err)
	var res *GroupStageDrawResult
	require.NoError(t, groups.Get(&res))
	require.Equal(t, 8, len(res.Groups))
}

func TestKnockoutPhaseDrawFixturesForRoundOf16(t *testing.T) {
	assertKnockoutPhaseDrawFixturesWorkCorrectly(t, Participants{
		Pots: []Pot{Pot1(), Pot2()},
	})
}

func TestKnockoutPhaseDrawFixturesForRoundOf8(t *testing.T) {
	assertKnockoutPhaseDrawFixturesWorkCorrectly(t, Participants{
		Pots: []Pot{Pot1()}, // Cheating
	})
}
