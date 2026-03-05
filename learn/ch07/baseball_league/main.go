package main

import (
	"fmt"
	"os"
)

func main() {
	var league League
	league.AddTeam(Team{Name: "Inky Squid", Members: []string{"Ja", "Oma", "Kuba"}})
	league.AddTeam(Team{Name: "Angry Flounder", Members: []string{"Kumo", "Ono", "Leno"}})
	league.AddTeam(Team{Name: "Happy Clam", Members: []string{"Clara", "Boe", "Neo"}})
	league.AddTeam(Team{Name: "Pink Crab", Members: []string{"Boye", "Tomo", "Miu"}})
	league.AddTeam((Team{Name: "Blue Beluga", Members: []string{"Mina", "Mino", "Mine"}}))

	league.MatchResult("Inky Squid", 16, "Angry Flounder", 9)
	league.MatchResult("Inky Squid", 6, "Happy Clam", 14)
	league.MatchResult("Inky Squid", 6, "Pink Crab", 4)
	league.MatchResult("Blue Beluga", 32, "Inky Squid", 31)
	league.MatchResult("Blue Beluga", 12, "Happy Clam", 2)
	league.MatchResult("Blue Beluga", 26, "Angry Flounder", 11)

	fmt.Println("\nTEAMS")
	for _, team := range league.Teams {
		team.Print()
	}

	fmt.Println("\nSCORE TEAM")
	fmt.Println("----- ----")
	for team, score := range league.Wins {
		fmt.Printf("%5d %s\n", score, team)

	}
	fmt.Println("\nRANKING")
	RankPrinter(&league, os.Stdout)
}
