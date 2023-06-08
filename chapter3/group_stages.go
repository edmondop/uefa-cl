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
	inter := Team{Name: "FC Internazionale"}
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

var groupStageResults map[string]GroupStageMatchResult

func getStageMatchResult(groupStageMatch GroupStageMatch) (GroupStageMatchResult, error) {
	key := groupStageMatch.HomeTeam.Name + "-" + groupStageMatch.AwayTeam.Name
	result, ok := groupStageResults[key]
	if !ok {
		return GroupStageMatchResult{}, errors.New("unknown match")
	}
	return result, nil
}

func addGroupAResults() {
	groupStageResults["AFC Ajax-Rangers FC"] = GroupStageMatchResult{4, 0}
	groupStageResults["SSC Napoli-Liverpool FC"] = GroupStageMatchResult{4, 1}
	groupStageResults["Liverpool FC-AFC Ajax"] = GroupStageMatchResult{2, 1}
	groupStageResults["Rangers FC-SSC Napoli"] = GroupStageMatchResult{0, 3}
	groupStageResults["AFC Ajax-SSC Napoli"] = GroupStageMatchResult{1, 6}
	groupStageResults["Liverpool FC-Rangers FC"] = GroupStageMatchResult{2, 0}
	groupStageResults["SSC Napoli-AFC Ajax"] = GroupStageMatchResult{4, 2}
	groupStageResults["Rangers FC-Liverpool FC"] = GroupStageMatchResult{1, 7}
	groupStageResults["AFC Ajax-Liverpool FC"] = GroupStageMatchResult{0, 3}
	groupStageResults["SSC Napoli-Rangers FC"] = GroupStageMatchResult{3, 0}
	groupStageResults["Rangers FC-AFC Ajax"] = GroupStageMatchResult{1, 3}
	groupStageResults["Liverpool FC-SSC Napoli"] = GroupStageMatchResult{2, 0}
}

func addGroupBResults() {
	groupStageResults["Atlético Madrid-FC Porto"] = GroupStageMatchResult{HomeGoals: 2, AwayGoals: 1}
	groupStageResults["Club Brugge KV-Bayer Leverkusen"] = GroupStageMatchResult{HomeGoals: 1, AwayGoals: 0}
	groupStageResults["Bayer Leverkusen-Atlético Madrid"] = GroupStageMatchResult{HomeGoals: 2, AwayGoals: 0}
	groupStageResults["FC Porto-Club Brugge KV"] = GroupStageMatchResult{HomeGoals: 0, AwayGoals: 4}
	groupStageResults["Club Brugge KV-Atlético Madrid"] = GroupStageMatchResult{HomeGoals: 2, AwayGoals: 0}
	groupStageResults["FC Porto-Bayer Leverkusen"] = GroupStageMatchResult{HomeGoals: 2, AwayGoals: 0}
	groupStageResults["Atlético Madrid-Club Brugge KV"] = GroupStageMatchResult{HomeGoals: 0, AwayGoals: 0}
	groupStageResults["Bayer Leverkusen-FC Porto"] = GroupStageMatchResult{HomeGoals: 0, AwayGoals: 3}
	groupStageResults["Club Brugge KV-FC Porto"] = GroupStageMatchResult{HomeGoals: 0, AwayGoals: 4}
	groupStageResults["Atlético Madrid-Bayer Leverkusen"] = GroupStageMatchResult{HomeGoals: 2, AwayGoals: 2}
	groupStageResults["FC Porto-Atlético Madrid"] = GroupStageMatchResult{HomeGoals: 2, AwayGoals: 1}
	groupStageResults["Bayer Leverkusen-Club Brugge KV"] = GroupStageMatchResult{HomeGoals: 0, AwayGoals: 0}
}

func addGroupCResults() {
	groupStageResults["FC Internazionale-Bayern München"] = GroupStageMatchResult{HomeGoals: 0, AwayGoals: 2}
	groupStageResults["FC Barcelona-Viktoria Plzeň"] = GroupStageMatchResult{HomeGoals: 5, AwayGoals: 1}
	groupStageResults["Viktoria Plzeň-FC Internazionale"] = GroupStageMatchResult{HomeGoals: 0, AwayGoals: 2}
	groupStageResults["Bayern München-FC Barcelona"] = GroupStageMatchResult{HomeGoals: 2, AwayGoals: 0}
	groupStageResults["Bayern München-Viktoria Plzeň"] = GroupStageMatchResult{HomeGoals: 5, AwayGoals: 0}
	groupStageResults["FC Internazionale-FC Barcelona"] = GroupStageMatchResult{HomeGoals: 1, AwayGoals: 0}
	groupStageResults["Viktoria Plzeň-Bayern München"] = GroupStageMatchResult{HomeGoals: 2, AwayGoals: 4}
	groupStageResults["FC Barcelona-FC Internazionale"] = GroupStageMatchResult{HomeGoals: 3, AwayGoals: 3}
	groupStageResults["FC Internazionale-Viktoria Plzeň"] = GroupStageMatchResult{HomeGoals: 4, AwayGoals: 0}
	groupStageResults["FC Barcelona-Bayern München"] = GroupStageMatchResult{HomeGoals: 0, AwayGoals: 3}
	groupStageResults["Bayern München-FC Internazionale"] = GroupStageMatchResult{HomeGoals: 2, AwayGoals: 0}
	groupStageResults["Viktoria Plzeň-FC Barcelona"] = GroupStageMatchResult{HomeGoals: 2, AwayGoals: 4}
}

func addGroupDResults() {
	groupStageResults["Eintracht Frankfurt-Sporting CP"] = GroupStageMatchResult{0, 3}
	groupStageResults["Tottenham Hotspur-Olympique Marseille"] = GroupStageMatchResult{2, 0}
	groupStageResults["Sporting CP-Tottenham Hotspur"] = GroupStageMatchResult{2, 0}
	groupStageResults["Olympique Marseille-Eintracht Frankfurt"] = GroupStageMatchResult{0, 1}
	groupStageResults["Olympique Marseille-Sporting CP"] = GroupStageMatchResult{4, 1}
	groupStageResults["Eintracht Frankfurt-Tottenham Hotspur"] = GroupStageMatchResult{0, 0}
	groupStageResults["Tottenham Hotspur-Eintracht Frankfurt"] = GroupStageMatchResult{3, 2}
	groupStageResults["Sporting CP-Olympique Marseille"] = GroupStageMatchResult{0, 2}
	groupStageResults["Eintracht Frankfurt-Olympique Marseille"] = GroupStageMatchResult{2, 1}
	groupStageResults["Tottenham Hotspur-Sporting CP"] = GroupStageMatchResult{1, 1}
	groupStageResults["Sporting CP-Eintracht Frankfurt"] = GroupStageMatchResult{1, 2}
	groupStageResults["Olympique Marseille-Tottenham Hotspur"] = GroupStageMatchResult{1, 2}
}

func addGroupEResults() {
	groupStageResults["Dinamo Zagreb-Chelsea FC"] = GroupStageMatchResult{1, 0}
	groupStageResults["RB Salzburg-AC Milan"] = GroupStageMatchResult{1, 1}
	groupStageResults["AC Milan-Dinamo Zagreb"] = GroupStageMatchResult{3, 1}
	groupStageResults["Chelsea FC-RB Salzburg"] = GroupStageMatchResult{1, 1}
	groupStageResults["RB Salzburg-Dinamo Zagreb"] = GroupStageMatchResult{1, 0}
	groupStageResults["Chelsea FC-AC Milan"] = GroupStageMatchResult{3, 0}
	groupStageResults["AC Milan-Chelsea FC"] = GroupStageMatchResult{0, 2}
	groupStageResults["Dinamo Zagreb-RB Salzburg"] = GroupStageMatchResult{1, 1}
	groupStageResults["RB Salzburg-Chelsea FC"] = GroupStageMatchResult{1, 2}
	groupStageResults["Dinamo Zagreb-AC Milan"] = GroupStageMatchResult{0, 4}
	groupStageResults["AC Milan-RB Salzburg"] = GroupStageMatchResult{4, 0}
	groupStageResults["Chelsea FC-Dinamo Zagreb"] = GroupStageMatchResult{2, 1}
}

func addGroupFResults() {
	groupStageResults["Celtic FC-Real Madrid"] = GroupStageMatchResult{0, 3}
	groupStageResults["RB Leipzig-Shakhtar Donetsk"] = GroupStageMatchResult{1, 4}
	groupStageResults["Shakhtar Donetsk-Celtic FC"] = GroupStageMatchResult{1, 1}
	groupStageResults["Real Madrid-RB Leipzig"] = GroupStageMatchResult{2, 0}
	groupStageResults["RB Leipzig-Celtic FC"] = GroupStageMatchResult{3, 1}
	groupStageResults["Real Madrid-Shakhtar Donetsk"] = GroupStageMatchResult{2, 1}
	groupStageResults["Celtic FC-RB Leipzig"] = GroupStageMatchResult{0, 2}
	groupStageResults["Shakhtar Donetsk-Real Madrid"] = GroupStageMatchResult{1, 1}
	groupStageResults["Celtic FC-Shakhtar Donetsk"] = GroupStageMatchResult{1, 1}
	groupStageResults["RB Leipzig-Real Madrid"] = GroupStageMatchResult{3, 2}
	groupStageResults["Real Madrid-Celtic FC"] = GroupStageMatchResult{5, 1}
	groupStageResults["Shakhtar Donetsk-RB Leipzig"] = GroupStageMatchResult{0, 4}
}

func addGroupGResults() {
	groupStageResults["Borussia Dortmund-FC København"] = GroupStageMatchResult{3, 0}
	groupStageResults["Sevilla FC-Manchester City"] = GroupStageMatchResult{0, 4}
	groupStageResults["Manchester City-Borussia Dortmund"] = GroupStageMatchResult{2, 1}
	groupStageResults["FC København-Sevilla FC"] = GroupStageMatchResult{0, 0}
	groupStageResults["Sevilla FC-Borussia Dortmund"] = GroupStageMatchResult{1, 4}
	groupStageResults["Manchester City-FC København"] = GroupStageMatchResult{5, 0}
	groupStageResults["FC København-Manchester City"] = GroupStageMatchResult{0, 0}
	groupStageResults["Borussia Dortmund-Sevilla FC"] = GroupStageMatchResult{1, 1}
	groupStageResults["Sevilla FC-FC København"] = GroupStageMatchResult{3, 0}
	groupStageResults["Borussia Dortmund-Manchester City"] = GroupStageMatchResult{0, 0}
	groupStageResults["FC København-Borussia Dortmund"] = GroupStageMatchResult{1, 1}
	groupStageResults["Manchester City-Sevilla FC"] = GroupStageMatchResult{3, 1}
}

func addGroupHResults() {
	groupStageResults["Paris Saint-Germain-Juventus"] = GroupStageMatchResult{2, 1}
	groupStageResults["SL Benfica-Maccabi Haifa"] = GroupStageMatchResult{2, 0}
	groupStageResults["Juventus-SL Benfica"] = GroupStageMatchResult{1, 2}
	groupStageResults["Maccabi Haifa-Paris Saint-Germain"] = GroupStageMatchResult{1, 3}
	groupStageResults["Juventus-Maccabi Haifa"] = GroupStageMatchResult{3, 1}
	groupStageResults["SL Benfica-Paris Saint-Germain"] = GroupStageMatchResult{1, 1}
	groupStageResults["Maccabi Haifa-Juventus"] = GroupStageMatchResult{2, 0}
	groupStageResults["Paris Saint-Germain-SL Benfica"] = GroupStageMatchResult{1, 1}
	groupStageResults["SL Benfica-Juventus"] = GroupStageMatchResult{4, 3}
	groupStageResults["Paris Saint-Germain-Maccabi Haifa"] = GroupStageMatchResult{7, 2}
	groupStageResults["Juventus-Paris Saint-Germain"] = GroupStageMatchResult{1, 2}
	groupStageResults["Maccabi Haifa-SL Benfica"] = GroupStageMatchResult{1, 6}
}

func init() {
	groupStageResults = make(map[string]GroupStageMatchResult)
	addGroupAResults()
	addGroupBResults()
	addGroupCResults()
	addGroupDResults()
	addGroupEResults()
	addGroupFResults()
	addGroupGResults()
	addGroupHResults()
}
