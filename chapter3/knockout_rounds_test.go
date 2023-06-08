package chapter3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterWonAgainstACMilan(t *testing.T) {
	firstLeg := KnockoutRoundLeg{
		HomeTeam: Team{
			Name: "FC Internazionale",
		},
		AwayTeam: Team{
			Name: "AC Milan",
		},
	}
	result, err := getKnockoutRoundResult(firstLeg)
	assert.NoError(t, err)
	assert.Equal(t, KnockoutStageMatchResult{HomeGoals: 2, AwayGoals: 0}, result)
}
