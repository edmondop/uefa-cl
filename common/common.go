package common

type Team struct {
	Name string
}

type Partecipants struct {
	SeededTeams []Team
	OtherTeams  []Team
}

type MatchResult struct {
	Team1 Team
	Team2 Team
	Goal1 int8
	Goal2 int8
}

func (m *MatchResult) Winner() *Team {
	if m.Goal1 > m.Goal2 {
		return &m.Team1
	}
	if m.Goal1 < m.Goal2 {
		return &m.Team2
	}
	return nil
}

type KnockoutPhaseResult struct {
	FinalMatch       MatchResult
	SemiFinalMatches []MatchResult
	QuarterFinalists []MatchResult
	RoundOf16        []MatchResult
}

type Result struct {
	Winner Team
}
