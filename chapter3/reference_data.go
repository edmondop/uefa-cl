package chapter3

import (
	"errors"
	"time"
)

func GroupA() Group {
	ajax := Team{Name: "AFC Ajax"}
	rangers := Team{Name: "Rangers FC"}
	napoli := Team{Name: "SSC Napoli"}
	liverpool := Team{Name: "Liverpool FC"}

	matches := []GroupStageMatch{
		{
			HomeTeam:  ajax,
			AwayTeam:  rangers,
			MatchDate: time.Date(2022, 9, 7, 17, 45, 0, 0, time.UTC),
		},
		{
			HomeTeam:  napoli,
			AwayTeam:  liverpool,
			MatchDate: time.Date(2022, 9, 7, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  liverpool,
			AwayTeam:  ajax,
			MatchDate: time.Date(2022, 9, 13, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  rangers,
			AwayTeam:  napoli,
			MatchDate: time.Date(2022, 9, 14, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  ajax,
			AwayTeam:  napoli,
			MatchDate: time.Date(2022, 10, 4, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  liverpool,
			AwayTeam:  rangers,
			MatchDate: time.Date(2022, 10, 4, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  napoli,
			AwayTeam:  ajax,
			MatchDate: time.Date(2022, 10, 12, 17, 45, 0, 0, time.UTC),
		},
		{
			HomeTeam:  rangers,
			AwayTeam:  liverpool,
			MatchDate: time.Date(2022, 10, 20, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  liverpool,
			AwayTeam:  napoli,
			MatchDate: time.Date(2022, 10, 26, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  napoli,
			AwayTeam:  rangers,
			MatchDate: time.Date(2022, 11, 1, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  rangers,
			AwayTeam:  ajax,
			MatchDate: time.Date(2022, 11, 2, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  liverpool,
			AwayTeam:  napoli,
			MatchDate: time.Date(2022, 11, 22, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  napoli,
			AwayTeam:  liverpool,
			MatchDate: time.Date(2022, 12, 6, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  rangers,
			AwayTeam:  liverpool,
			MatchDate: time.Date(2022, 12, 7, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  napoli,
			AwayTeam:  rangers,
			MatchDate: time.Date(2022, 12, 13, 20, 0, 0, 0, time.UTC),
		},
	}

	return Group{
		GroupName: "Group A",
		Teams:     [4]Team{ajax, rangers, napoli, liverpool},
		GroupCalendar: GroupCalendar{
			Matches: matches,
		},
	}
}

func GroupB() Group {
	atleticoMadrid := Team{Name: "Atlético Madrid"}
	fcPorto := Team{Name: "FC Porto"}
	clubBruggeKV := Team{Name: "Club Brugge KV"}
	bayerLeverkusen := Team{Name: "Bayer Leverkusen"}

	matches := []GroupStageMatch{
		{
			HomeTeam:  atleticoMadrid,
			AwayTeam:  fcPorto,
			MatchDate: time.Date(2022, 9, 7, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  clubBruggeKV,
			AwayTeam:  bayerLeverkusen,
			MatchDate: time.Date(2022, 9, 7, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  bayerLeverkusen,
			AwayTeam:  atleticoMadrid,
			MatchDate: time.Date(2022, 9, 13, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  fcPorto,
			AwayTeam:  clubBruggeKV,
			MatchDate: time.Date(2022, 9, 13, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  clubBruggeKV,
			AwayTeam:  atleticoMadrid,
			MatchDate: time.Date(2022, 10, 4, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  fcPorto,
			AwayTeam:  bayerLeverkusen,
			MatchDate: time.Date(2022, 10, 4, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  atleticoMadrid,
			AwayTeam:  clubBruggeKV,
			MatchDate: time.Date(2022, 10, 12, 17, 45, 0, 0, time.UTC),
		},
		{
			HomeTeam:  bayerLeverkusen,
			AwayTeam:  fcPorto,
			MatchDate: time.Date(2022, 10, 12, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  clubBruggeKV,
			AwayTeam:  fcPorto,
			MatchDate: time.Date(2022, 10, 26, 17, 45, 0, 0, time.UTC),
		},
		{
			HomeTeam:  atleticoMadrid,
			AwayTeam:  bayerLeverkusen,
			MatchDate: time.Date(2022, 10, 26, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  fcPorto,
			AwayTeam:  atleticoMadrid,
			MatchDate: time.Date(2022, 11, 1, 17, 45, 0, 0, time.UTC),
		},
		{
			HomeTeam:  bayerLeverkusen,
			AwayTeam:  clubBruggeKV,
			MatchDate: time.Date(2022, 11, 1, 17, 45, 0, 0, time.UTC),
		},
	}

	return Group{
		GroupName: "Group B",
		Teams:     [4]Team{atleticoMadrid, fcPorto, clubBruggeKV, bayerLeverkusen},
		GroupCalendar: GroupCalendar{
			Matches: matches,
		},
	}
}

func GroupC() Group {
	inter := Team{Name: "Inter"}
	bayernMunchen := Team{Name: "Bayern München"}
	barcelona := Team{Name: "FC Barcelona"}
	plzen := Team{Name: "Viktoria Plzeň"}

	matches := []GroupStageMatch{
		{
			HomeTeam:  inter,
			AwayTeam:  bayernMunchen,
			MatchDate: time.Date(2022, 9, 7, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  barcelona,
			AwayTeam:  plzen,
			MatchDate: time.Date(2022, 9, 7, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  plzen,
			AwayTeam:  inter,
			MatchDate: time.Date(2022, 9, 13, 17, 45, 0, 0, time.UTC),
		},
		{
			HomeTeam:  bayernMunchen,
			AwayTeam:  barcelona,
			MatchDate: time.Date(2022, 9, 13, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  bayernMunchen,
			AwayTeam:  plzen,
			MatchDate: time.Date(2022, 10, 4, 17, 45, 0, 0, time.UTC),
		},
		{
			HomeTeam:  inter,
			AwayTeam:  barcelona,
			MatchDate: time.Date(2022, 10, 4, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  plzen,
			AwayTeam:  bayernMunchen,
			MatchDate: time.Date(2022, 10, 12, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  barcelona,
			AwayTeam:  inter,
			MatchDate: time.Date(2022, 10, 12, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  inter,
			AwayTeam:  plzen,
			MatchDate: time.Date(2022, 10, 26, 17, 45, 0, 0, time.UTC),
		},
		{
			HomeTeam:  barcelona,
			AwayTeam:  bayernMunchen,
			MatchDate: time.Date(2022, 10, 26, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  bayernMunchen,
			AwayTeam:  inter,
			MatchDate: time.Date(2022, 11, 1, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  plzen,
			AwayTeam:  barcelona,
			MatchDate: time.Date(2022, 11, 1, 20, 0, 0, 0, time.UTC),
		},
	}

	return Group{
		GroupName: "Group C",
		Teams:     [4]Team{inter, bayernMunchen, barcelona, plzen},
		GroupCalendar: GroupCalendar{
			Matches: matches,
		},
	}
}

func GroupD() Group {
	frankfurt := Team{Name: "Eintracht Frankfurt"}
	sportingCP := Team{Name: "Sporting CP"}
	tottenham := Team{Name: "Tottenham Hotspur"}
	marseille := Team{Name: "Olympique Marseille"}

	matches := []GroupStageMatch{
		{
			HomeTeam:  frankfurt,
			AwayTeam:  sportingCP,
			MatchDate: time.Date(2022, 9, 7, 17, 45, 0, 0, time.UTC),
		},
		{
			HomeTeam:  tottenham,
			AwayTeam:  marseille,
			MatchDate: time.Date(2022, 9, 7, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  sportingCP,
			AwayTeam:  tottenham,
			MatchDate: time.Date(2022, 9, 13, 17, 45, 0, 0, time.UTC),
		},
		{
			HomeTeam:  marseille,
			AwayTeam:  frankfurt,
			MatchDate: time.Date(2022, 9, 13, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  marseille,
			AwayTeam:  sportingCP,
			MatchDate: time.Date(2022, 10, 4, 17, 45, 0, 0, time.UTC),
		},
		{
			HomeTeam:  frankfurt,
			AwayTeam:  tottenham,
			MatchDate: time.Date(2022, 10, 4, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  tottenham,
			AwayTeam:  frankfurt,
			MatchDate: time.Date(2022, 10, 12, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  sportingCP,
			AwayTeam:  marseille,
			MatchDate: time.Date(2022, 10, 26, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  frankfurt,
			AwayTeam:  marseille,
			MatchDate: time.Date(2022, 10, 26, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  tottenham,
			AwayTeam:  sportingCP,
			MatchDate: time.Date(2022, 11, 1, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  sportingCP,
			AwayTeam:  frankfurt,
			MatchDate: time.Date(2022, 11, 1, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  marseille,
			AwayTeam:  tottenham,
			MatchDate: time.Date(2022, 11, 1, 20, 0, 0, 0, time.UTC),
		},
	}

	return Group{
		GroupName: "D",
		Teams:     [4]Team{frankfurt, sportingCP, tottenham, marseille},
		GroupCalendar: GroupCalendar{
			Matches: matches,
		},
	}
}

func GroupE() Group {
	dinamoZagreb := Team{Name: "Dinamo Zagreb"}
	chelseaFC := Team{Name: "Chelsea FC"}
	rbSalzburg := Team{Name: "RB Salzburg"}
	acMilan := Team{Name: "AC Milan"}

	matches := []GroupStageMatch{
		{
			HomeTeam:  dinamoZagreb,
			AwayTeam:  chelseaFC,
			MatchDate: time.Date(2022, 9, 6, 17, 45, 0, 0, time.UTC),
		},
		{
			HomeTeam:  rbSalzburg,
			AwayTeam:  acMilan,
			MatchDate: time.Date(2022, 9, 6, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  acMilan,
			AwayTeam:  dinamoZagreb,
			MatchDate: time.Date(2022, 9, 14, 17, 45, 0, 0, time.UTC),
		},
		{
			HomeTeam:  chelseaFC,
			AwayTeam:  rbSalzburg,
			MatchDate: time.Date(2022, 9, 14, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  rbSalzburg,
			AwayTeam:  dinamoZagreb,
			MatchDate: time.Date(2022, 10, 5, 17, 45, 0, 0, time.UTC),
		},
		{
			HomeTeam:  chelseaFC,
			AwayTeam:  acMilan,
			MatchDate: time.Date(2022, 10, 5, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  acMilan,
			AwayTeam:  chelseaFC,
			MatchDate: time.Date(2022, 10, 11, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  dinamoZagreb,
			AwayTeam:  rbSalzburg,
			MatchDate: time.Date(2022, 10, 11, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  rbSalzburg,
			AwayTeam:  chelseaFC,
			MatchDate: time.Date(2022, 10, 25, 17, 45, 0, 0, time.UTC),
		},
		{
			HomeTeam:  dinamoZagreb,
			AwayTeam:  acMilan,
			MatchDate: time.Date(2022, 10, 25, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  acMilan,
			AwayTeam:  rbSalzburg,
			MatchDate: time.Date(2022, 11, 2, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  chelseaFC,
			AwayTeam:  dinamoZagreb,
			MatchDate: time.Date(2022, 11, 2, 20, 0, 0, 0, time.UTC),
		},
	}

	return Group{
		GroupName: "E",
		Teams:     [4]Team{dinamoZagreb, chelseaFC, rbSalzburg, acMilan},
		GroupCalendar: GroupCalendar{
			Matches: matches,
		},
	}
}

func GroupF() Group {
	celticFC := Team{Name: "Celtic FC"}
	realMadrid := Team{Name: "Real Madrid"}
	rbLeipzig := Team{Name: "RB Leipzig"}
	shakhtarDonetsk := Team{Name: "Shakhtar Donetsk"}

	matches := []GroupStageMatch{
		{
			HomeTeam:  celticFC,
			AwayTeam:  realMadrid,
			MatchDate: time.Date(2022, 9, 6, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  rbLeipzig,
			AwayTeam:  shakhtarDonetsk,
			MatchDate: time.Date(2022, 9, 6, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  shakhtarDonetsk,
			AwayTeam:  celticFC,
			MatchDate: time.Date(2022, 9, 14, 17, 45, 0, 0, time.UTC),
		},
		{
			HomeTeam:  realMadrid,
			AwayTeam:  rbLeipzig,
			MatchDate: time.Date(2022, 9, 14, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  rbLeipzig,
			AwayTeam:  celticFC,
			MatchDate: time.Date(2022, 10, 5, 17, 45, 0, 0, time.UTC),
		},
		{
			HomeTeam:  realMadrid,
			AwayTeam:  shakhtarDonetsk,
			MatchDate: time.Date(2022, 10, 5, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  celticFC,
			AwayTeam:  rbLeipzig,
			MatchDate: time.Date(2022, 10, 11, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  shakhtarDonetsk,
			AwayTeam:  realMadrid,
			MatchDate: time.Date(2022, 10, 11, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  celticFC,
			AwayTeam:  shakhtarDonetsk,
			MatchDate: time.Date(2022, 10, 25, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  rbLeipzig,
			AwayTeam:  realMadrid,
			MatchDate: time.Date(2022, 10, 25, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  realMadrid,
			AwayTeam:  celticFC,
			MatchDate: time.Date(2022, 11, 2, 17, 45, 0, 0, time.UTC),
		},
		{
			HomeTeam:  shakhtarDonetsk,
			AwayTeam:  rbLeipzig,
			MatchDate: time.Date(2022, 11, 2, 17, 45, 0, 0, time.UTC),
		},
	}

	return Group{
		GroupName: "F",
		Teams:     [4]Team{celticFC, realMadrid, rbLeipzig, shakhtarDonetsk},
		GroupCalendar: GroupCalendar{
			Matches: matches,
		},
	}
}

func GroupG() Group {
	borussiaDortmund := Team{Name: "Borussia Dortmund"}
	fcKobenhavn := Team{Name: "FC København"}
	sevillaFC := Team{Name: "Sevilla FC"}
	manchesterCity := Team{Name: "Manchester City"}

	matches := []GroupStageMatch{
		{
			HomeTeam:  borussiaDortmund,
			AwayTeam:  fcKobenhavn,
			MatchDate: time.Date(2022, 9, 6, 17, 45, 0, 0, time.UTC),
		},
		{
			HomeTeam:  sevillaFC,
			AwayTeam:  manchesterCity,
			MatchDate: time.Date(2022, 9, 6, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  manchesterCity,
			AwayTeam:  borussiaDortmund,
			MatchDate: time.Date(2022, 9, 14, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  fcKobenhavn,
			AwayTeam:  sevillaFC,
			MatchDate: time.Date(2022, 9, 14, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  sevillaFC,
			AwayTeam:  borussiaDortmund,
			MatchDate: time.Date(2022, 10, 5, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  manchesterCity,
			AwayTeam:  fcKobenhavn,
			MatchDate: time.Date(2022, 10, 5, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  fcKobenhavn,
			AwayTeam:  manchesterCity,
			MatchDate: time.Date(2022, 10, 11, 17, 45, 0, 0, time.UTC),
		},
		{
			HomeTeam:  borussiaDortmund,
			AwayTeam:  sevillaFC,
			MatchDate: time.Date(2022, 10, 11, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  sevillaFC,
			AwayTeam:  fcKobenhavn,
			MatchDate: time.Date(2022, 10, 25, 17, 45, 0, 0, time.UTC),
		},
		{
			HomeTeam:  borussiaDortmund,
			AwayTeam:  manchesterCity,
			MatchDate: time.Date(2022, 10, 25, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  fcKobenhavn,
			AwayTeam:  borussiaDortmund,
			MatchDate: time.Date(2022, 11, 2, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  manchesterCity,
			AwayTeam:  sevillaFC,
			MatchDate: time.Date(2022, 11, 2, 20, 0, 0, 0, time.UTC),
		},
	}

	return Group{
		GroupName: "Group G",
		Teams:     [4]Team{borussiaDortmund, fcKobenhavn, sevillaFC, manchesterCity},
		GroupCalendar: GroupCalendar{
			Matches: matches,
		},
	}
}

func GroupH() Group {
	psg := Team{Name: "Paris Saint-Germain"}
	juventus := Team{Name: "Juventus"}
	benfica := Team{Name: "SL Benfica"}
	maccabi := Team{Name: "Maccabi Haifa"}

	matches := []GroupStageMatch{
		{
			HomeTeam:  psg,
			AwayTeam:  juventus,
			MatchDate: time.Date(2022, 9, 6, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  benfica,
			AwayTeam:  maccabi,
			MatchDate: time.Date(2022, 9, 6, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  juventus,
			AwayTeam:  benfica,
			MatchDate: time.Date(2022, 9, 14, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  maccabi,
			AwayTeam:  psg,
			MatchDate: time.Date(2022, 9, 14, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  juventus,
			AwayTeam:  maccabi,
			MatchDate: time.Date(2022, 10, 5, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  benfica,
			AwayTeam:  psg,
			MatchDate: time.Date(2022, 10, 5, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  maccabi,
			AwayTeam:  juventus,
			MatchDate: time.Date(2022, 10, 11, 17, 45, 0, 0, time.UTC),
		},
		{
			HomeTeam:  psg,
			AwayTeam:  benfica,
			MatchDate: time.Date(2022, 10, 11, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  benfica,
			AwayTeam:  juventus,
			MatchDate: time.Date(2022, 10, 25, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  psg,
			AwayTeam:  maccabi,
			MatchDate: time.Date(2022, 10, 25, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  juventus,
			AwayTeam:  psg,
			MatchDate: time.Date(2022, 11, 2, 20, 0, 0, 0, time.UTC),
		},
		{
			HomeTeam:  maccabi,
			AwayTeam:  benfica,
			MatchDate: time.Date(2022, 11, 2, 20, 0, 0, 0, time.UTC),
		},
	}

	return Group{
		GroupName: "H",
		Teams:     [4]Team{psg, juventus, benfica, maccabi},
		GroupCalendar: GroupCalendar{
			Matches: matches,
		},
	}
}

var results map[string]GroupStageMatchResult

func getResult(groupStageMatch GroupStageMatch) (GroupStageMatchResult, error) {
	key := groupStageMatch.HomeTeam.Name + "-" + groupStageMatch.AwayTeam.Name
	result, ok := results[key]
	if !ok {
		return GroupStageMatchResult{}, errors.New("unknown match")
	}
	return result, nil
}

func addGroupAResults() {
	results["AFC Ajax-Rangers FC"] = GroupStageMatchResult{4, 0}
	results["SSC Napoli-Liverpool FC"] = GroupStageMatchResult{4, 1}
	results["Liverpool FC-AFC Ajax"] = GroupStageMatchResult{2, 1}
	results["Rangers FC-SSC Napoli"] = GroupStageMatchResult{0, 3}
	results["AFC Ajax-SSC Napoli"] = GroupStageMatchResult{1, 6}
	results["Liverpool FC-Rangers FC"] = GroupStageMatchResult{2, 0}
	results["SSC Napoli-AFC Ajax"] = GroupStageMatchResult{4, 2}
	results["Rangers FC-Liverpool FC"] = GroupStageMatchResult{1, 7}
	results["AFC Ajax-Liverpool FC"] = GroupStageMatchResult{0, 3}
	results["SSC Napoli-Rangers FC"] = GroupStageMatchResult{3, 0}
	results["Rangers FC-AFC Ajax"] = GroupStageMatchResult{1, 3}
	results["Liverpool FC-SSC Napoli"] = GroupStageMatchResult{2, 0}
}

func addGroupBResults() {
	results["Atlético Madrid-FC Porto"] = GroupStageMatchResult{HomeGoals: 2, AwayGoals: 1}
	results["Club Brugge KV-Bayer Leverkusen"] = GroupStageMatchResult{HomeGoals: 1, AwayGoals: 0}
	results["Bayer Leverkusen-Atlético Madrid"] = GroupStageMatchResult{HomeGoals: 2, AwayGoals: 0}
	results["FC Porto-Club Brugge KV"] = GroupStageMatchResult{HomeGoals: 0, AwayGoals: 4}
	results["Club Brugge KV-Atlético Madrid"] = GroupStageMatchResult{HomeGoals: 2, AwayGoals: 0}
	results["FC Porto-Bayer Leverkusen"] = GroupStageMatchResult{HomeGoals: 2, AwayGoals: 0}
	results["Atlético Madrid-Club Brugge KV"] = GroupStageMatchResult{HomeGoals: 0, AwayGoals: 0}
	results["Bayer Leverkusen-FC Porto"] = GroupStageMatchResult{HomeGoals: 0, AwayGoals: 3}
	results["Club Brugge KV-FC Porto"] = GroupStageMatchResult{HomeGoals: 0, AwayGoals: 4}
	results["Atlético Madrid-Bayer Leverkusen"] = GroupStageMatchResult{HomeGoals: 2, AwayGoals: 2}
	results["FC Porto-Atlético Madrid"] = GroupStageMatchResult{HomeGoals: 2, AwayGoals: 1}
	results["Bayer Leverkusen-Club Brugge KV"] = GroupStageMatchResult{HomeGoals: 0, AwayGoals: 0}
}

func addGroupCResults() {
	results["Inter-Bayern München"] = GroupStageMatchResult{HomeGoals: 0, AwayGoals: 2}
	results["FC Barcelona-Viktoria Plzeň"] = GroupStageMatchResult{HomeGoals: 5, AwayGoals: 1}
	results["Viktoria Plzeň-Inter"] = GroupStageMatchResult{HomeGoals: 0, AwayGoals: 2}
	results["Bayern München-FC Barcelona"] = GroupStageMatchResult{HomeGoals: 2, AwayGoals: 0}
	results["Bayern München-Viktoria Plzeň"] = GroupStageMatchResult{HomeGoals: 5, AwayGoals: 0}
	results["Inter-FC Barcelona"] = GroupStageMatchResult{HomeGoals: 1, AwayGoals: 0}
	results["Viktoria Plzeň-Bayern München"] = GroupStageMatchResult{HomeGoals: 2, AwayGoals: 4}
	results["FC Barcelona-Inter"] = GroupStageMatchResult{HomeGoals: 3, AwayGoals: 3}
	results["Inter-Viktoria Plzeň"] = GroupStageMatchResult{HomeGoals: 4, AwayGoals: 0}
	results["FC Barcelona-Bayern München"] = GroupStageMatchResult{HomeGoals: 0, AwayGoals: 3}
	results["Bayern München-Inter"] = GroupStageMatchResult{HomeGoals: 2, AwayGoals: 0}
	results["Viktoria Plzeň-FC Barcelona"] = GroupStageMatchResult{HomeGoals: 2, AwayGoals: 4}
}

func addGroupDResults() {
	results["Eintracht Frankfurt-Sporting CP"] = GroupStageMatchResult{0, 3}
	results["Tottenham Hotspur-Olympique Marseille"] = GroupStageMatchResult{2, 0}
	results["Sporting CP-Tottenham Hotspur"] = GroupStageMatchResult{2, 0}
	results["Olympique Marseille-Eintracht Frankfurt"] = GroupStageMatchResult{0, 1}
	results["Olympique Marseille-Sporting CP"] = GroupStageMatchResult{4, 1}
	results["Eintracht Frankfurt-Tottenham Hotspur"] = GroupStageMatchResult{0, 0}
	results["Tottenham Hotspur-Eintracht Frankfurt"] = GroupStageMatchResult{3, 2}
	results["Sporting CP-Olympique Marseille"] = GroupStageMatchResult{0, 2}
	results["Eintracht Frankfurt-Olympique Marseille"] = GroupStageMatchResult{2, 1}
	results["Tottenham Hotspur-Sporting CP"] = GroupStageMatchResult{1, 1}
	results["Sporting CP-Eintracht Frankfurt"] = GroupStageMatchResult{1, 2}
	results["Olympique Marseille-Tottenham Hotspur"] = GroupStageMatchResult{1, 2}
}

func addGroupEResults() {
	results["Dinamo Zagreb-Chelsea FC"] = GroupStageMatchResult{1, 0}
	results["RB Salzburg-AC Milan"] = GroupStageMatchResult{1, 1}
	results["AC Milan-Dinamo Zagreb"] = GroupStageMatchResult{3, 1}
	results["Chelsea FC-RB Salzburg"] = GroupStageMatchResult{1, 1}
	results["RB Salzburg-Dinamo Zagreb"] = GroupStageMatchResult{1, 0}
	results["Chelsea FC-AC Milan"] = GroupStageMatchResult{3, 0}
	results["AC Milan-Chelsea FC"] = GroupStageMatchResult{0, 2}
	results["Dinamo Zagreb-RB Salzburg"] = GroupStageMatchResult{1, 1}
	results["RB Salzburg-Chelsea FC"] = GroupStageMatchResult{1, 2}
	results["Dinamo Zagreb-AC Milan"] = GroupStageMatchResult{0, 4}
	results["AC Milan-RB Salzburg"] = GroupStageMatchResult{4, 0}
	results["Chelsea FC-Dinamo Zagreb"] = GroupStageMatchResult{2, 1}
}

func addGroupFResults() {
	results["Celtic FC-Real Madrid"] = GroupStageMatchResult{0, 3}
	results["RB Leipzig-Shakhtar Donetsk"] = GroupStageMatchResult{1, 4}
	results["Shakhtar Donetsk-Celtic FC"] = GroupStageMatchResult{1, 1}
	results["Real Madrid-RB Leipzig"] = GroupStageMatchResult{2, 0}
	results["RB Leipzig-Celtic FC"] = GroupStageMatchResult{3, 1}
	results["Real Madrid-Shakhtar Donetsk"] = GroupStageMatchResult{2, 1}
	results["Celtic FC-RB Leipzig"] = GroupStageMatchResult{0, 2}
	results["Shakhtar Donetsk-Real Madrid"] = GroupStageMatchResult{1, 1}
	results["Celtic FC-Shakhtar Donetsk"] = GroupStageMatchResult{1, 1}
	results["RB Leipzig-Real Madrid"] = GroupStageMatchResult{3, 2}
	results["Real Madrid-Celtic FC"] = GroupStageMatchResult{5, 1}
	results["Shakhtar Donetsk-RB Leipzig"] = GroupStageMatchResult{0, 4}
}

func addGroupGResults() {
	results["Borussia Dortmund-FC København"] = GroupStageMatchResult{3, 0}
	results["Sevilla FC-Manchester City"] = GroupStageMatchResult{0, 4}
	results["Manchester City-Borussia Dortmund"] = GroupStageMatchResult{2, 1}
	results["FC København-Sevilla FC"] = GroupStageMatchResult{0, 0}
	results["Sevilla FC-Borussia Dortmund"] = GroupStageMatchResult{1, 4}
	results["Manchester City-FC København"] = GroupStageMatchResult{5, 0}
	results["FC København-Manchester City"] = GroupStageMatchResult{0, 0}
	results["Borussia Dortmund-Sevilla FC"] = GroupStageMatchResult{1, 1}
	results["Sevilla FC-FC København"] = GroupStageMatchResult{3, 0}
	results["Borussia Dortmund-Manchester City"] = GroupStageMatchResult{0, 0}
	results["FC København-Borussia Dortmund"] = GroupStageMatchResult{1, 1}
	results["Manchester City-Sevilla FC"] = GroupStageMatchResult{3, 1}
}

func addGroupHResults() {
	results["Paris Saint-Germain-Juventus"] = GroupStageMatchResult{2, 1}
	results["SL Benfica-Maccabi Haifa"] = GroupStageMatchResult{2, 0}
	results["Juventus-SL Benfica"] = GroupStageMatchResult{1, 2}
	results["Maccabi Haifa-Paris Saint-Germain"] = GroupStageMatchResult{1, 3}
	results["Juventus-Maccabi Haifa"] = GroupStageMatchResult{3, 1}
	results["SL Benfica-Paris Saint-Germain"] = GroupStageMatchResult{1, 1}
	results["Maccabi Haifa-Juventus"] = GroupStageMatchResult{2, 0}
	results["Paris Saint-Germain-SL Benfica"] = GroupStageMatchResult{1, 1}
	results["SL Benfica-Juventus"] = GroupStageMatchResult{4, 3}
	results["Paris Saint-Germain-Maccabi Haifa"] = GroupStageMatchResult{7, 2}
	results["Juventus-Paris Saint-Germain"] = GroupStageMatchResult{1, 2}
	results["Maccabi Haifa-SL Benfica"] = GroupStageMatchResult{1, 6}
}

func init() {
	results = make(map[string]GroupStageMatchResult)
	addGroupAResults()
	addGroupBResults()
	addGroupCResults()
	addGroupDResults()
	addGroupEResults()
	addGroupFResults()
	addGroupGResults()
	addGroupHResults()
}
