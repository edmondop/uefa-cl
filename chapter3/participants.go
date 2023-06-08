package chapter3

type Team struct {
	Name string
}

type Pot struct {
	Teams []Team
}

type Participants struct {
	Pots []Pot
}

func (p *Participants) TeamCount() int {
	count := 0
	for _, pot := range p.Pots {
		count += len(pot.Teams)
	}
	return count
}

type KnockoutRoundLeg struct {
	HomeTeam Team
	AwayTeam Team
}
type KnockoutRoundFixture struct {
	FirstLeg  KnockoutRoundLeg
	SecondLeg KnockoutRoundLeg
}

type Result struct {
	Winner Team
}

func Pot1() Pot {
	return Pot{
		Teams: []Team{
			{Name: "Real Madrid"},
			{Name: "Eintracht Frankfurt"},
			{Name: "Manchester City"},
			{Name: "AC Milan"},
			{Name: "Bayern Munich"},
			{Name: "Paris Saint-Germain"},
			{Name: "FC Porto"},
			{Name: "Ajax"},
		},
	}
}

func Pot2() Pot {
	return Pot{
		[]Team{
			{Name: "Liverpool"},
			{Name: "Chelsea"},
			{Name: "Barcelona"},
			{Name: "Juventus"},
			{Name: "Atletico Madrid"},
			{Name: "Sevilla"},
			{Name: "RB Leipzig"},
			{Name: "Tottenham Hotspur"},
		},
	}
}

func Pot3() Pot {
	return Pot{
		[]Team{
			{Name: "Borussia Dortmund"},
			{Name: "FC Salzburg"},
			{Name: "Shakhtar Donetsk"},
			{Name: "FC Internazionale"},
			{Name: "Napoli"},
			{Name: "Benfica"},
			{Name: "Sporting CP"},
			{Name: "Bayer Leverkusen"},
		},
	}
}

func Pot4() Pot {
	return Pot{
		[]Team{
			{Name: "Marseille"},
			{Name: "Brugge"},
			{Name: "Celtic"},
			{Name: "Viktoria Plzen"},
			{Name: "Maccabi Haifa"},
			{Name: "Rangers"},
			{Name: "FC Copenhagen"},
			{Name: "Dinamo Zagreb"},
		},
	}
}

func GetParticipants() Participants {
	return Participants{
		Pots: []Pot{
			Pot1(),
			Pot2(),
			Pot3(),
			Pot4(),
		},
	}
}
