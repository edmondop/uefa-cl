package chapter2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadTeamsSuccessfully(t *testing.T) {
	participants, err := ReadParticipants("../static/group_stages/pot%d.txt")
	assert.NoError(t, err)

	assert.Equal(t, 4, len(participants.Pots))
	for i, pot := range participants.Pots {
		assert.Equal(t, 8, len(pot.Teams), "Pot %d should contain 8 teams", i+1)
	}
}

// TODO: tests for error cases
