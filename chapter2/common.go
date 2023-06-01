package chapter2

import (
	"bufio"
	"fmt"
	"os"
)

type Team struct {
	Name string
}

type Pot struct {
	Teams []Team
}

type Participants struct {
	Pots []Pot
}

type Result struct {
	Winner Team
}

func readPot(filepath string) (Pot, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return Pot{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	teams := []Team{}
	for scanner.Scan() {
		teams = append(teams, Team{Name: scanner.Text()})
	}

	if err := scanner.Err(); err != nil {
		return Pot{}, err
	}

	return Pot{Teams: teams}, nil
}

func ReadParticipants(potPathTemplate string) (Participants, error) {
	pots := []Pot{}
	for i := 1; i <= 4; i++ {
		filepath := fmt.Sprintf(potPathTemplate, i)
		pot, err := readPot(filepath)
		if err != nil {
			return Participants{}, err
		}
		pots = append(pots, pot)
	}
	return Participants{Pots: pots}, nil
}
