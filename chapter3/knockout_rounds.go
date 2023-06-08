package chapter3

import "errors"

func addRoundOf8Results() {
	knockoutStageResults["SL Benfica-FC Internazionale"] = KnockoutStageMatchResult{HomeGoals: 0, AwayGoals: 2}
	knockoutStageResults["Manchester City-Bayern München"] = KnockoutStageMatchResult{HomeGoals: 3, AwayGoals: 0}
	knockoutStageResults["Real Madrid-Chelsea FC"] = KnockoutStageMatchResult{HomeGoals: 2, AwayGoals: 0}
	knockoutStageResults["AC Milan-SSC Napoli"] = KnockoutStageMatchResult{HomeGoals: 1, AwayGoals: 0}
	knockoutStageResults["Chelsea FC-Real Madrid"] = KnockoutStageMatchResult{HomeGoals: 0, AwayGoals: 2}
	knockoutStageResults["SSC Napoli-AC Milan"] = KnockoutStageMatchResult{HomeGoals: 1, AwayGoals: 1}
	knockoutStageResults["FC Internazionale-SL Benfica"] = KnockoutStageMatchResult{HomeGoals: 3, AwayGoals: 3}
	knockoutStageResults["Bayern München-Manchester City"] = KnockoutStageMatchResult{HomeGoals: 1, AwayGoals: 1}
}

func addRoundOf4Results() {
	knockoutStageResults["Real Madrid-Manchester City"] = KnockoutStageMatchResult{HomeGoals: 1, AwayGoals: 1}
	knockoutStageResults["AC Milan-FC Internazionale"] = KnockoutStageMatchResult{HomeGoals: 0, AwayGoals: 2}
	knockoutStageResults["FC Internazionale-AC Milan"] = KnockoutStageMatchResult{HomeGoals: 1, AwayGoals: 0}
	knockoutStageResults["Manchester City-Real Madrid"] = KnockoutStageMatchResult{HomeGoals: 4, AwayGoals: 0}
}

var knockoutStageResults map[string]KnockoutStageMatchResult

func addRoundOf16Results() {
	knockoutStageResults["AC Milan-Tottenham Hotspur"] = KnockoutStageMatchResult{HomeGoals: 1, AwayGoals: 0}
	knockoutStageResults["Paris Saint-Germain-Bayern München"] = KnockoutStageMatchResult{HomeGoals: 0, AwayGoals: 1}
	knockoutStageResults["Club Brugge KV-SL Benfica"] = KnockoutStageMatchResult{HomeGoals: 0, AwayGoals: 2}
	knockoutStageResults["Borussia Dortmund-Chelsea FC"] = KnockoutStageMatchResult{HomeGoals: 1, AwayGoals: 0}
	knockoutStageResults["Liverpool FC-Real Madrid"] = KnockoutStageMatchResult{HomeGoals: 2, AwayGoals: 5}
	knockoutStageResults["Eintracht Frankfurt-SSC Napoli"] = KnockoutStageMatchResult{HomeGoals: 0, AwayGoals: 2}
	knockoutStageResults["RB Leipzig-Manchester City"] = KnockoutStageMatchResult{HomeGoals: 1, AwayGoals: 1}
	knockoutStageResults["FC Internazionale-FC Porto"] = KnockoutStageMatchResult{HomeGoals: 1, AwayGoals: 0}
	knockoutStageResults["SL Benfica-Club Brugge KV"] = KnockoutStageMatchResult{HomeGoals: 5, AwayGoals: 1}
	knockoutStageResults["Chelsea FC-Borussia Dortmund"] = KnockoutStageMatchResult{HomeGoals: 2, AwayGoals: 0}
	knockoutStageResults["Bayern München-Paris Saint-Germain"] = KnockoutStageMatchResult{HomeGoals: 2, AwayGoals: 0}
	knockoutStageResults["Tottenham Hotspur-AC Milan"] = KnockoutStageMatchResult{HomeGoals: 0, AwayGoals: 0}
	knockoutStageResults["Manchester City-RB Leipzig"] = KnockoutStageMatchResult{HomeGoals: 7, AwayGoals: 0}
	knockoutStageResults["FC Porto-FC Internazionale"] = KnockoutStageMatchResult{HomeGoals: 0, AwayGoals: 0}
	knockoutStageResults["Real Madrid-Liverpool FC"] = KnockoutStageMatchResult{HomeGoals: 1, AwayGoals: 0}
	knockoutStageResults["SSC Napoli-Eintracht Frankfurt"] = KnockoutStageMatchResult{HomeGoals: 3, AwayGoals: 0}
}

func init() {
	knockoutStageResults = make(map[string]KnockoutStageMatchResult, 28)
	addRoundOf16Results()
	addRoundOf8Results()
	addRoundOf4Results()
}

func getKnockoutRoundResult(knockoutRoundLeg KnockoutRoundLeg) (KnockoutStageMatchResult, error) {
	key := knockoutRoundLeg.HomeTeam.Name + "-" + knockoutRoundLeg.AwayTeam.Name
	result, ok := knockoutStageResults[key]
	if !ok {
		return KnockoutStageMatchResult{}, errors.New("unknown match")
	}
	return result, nil
}
