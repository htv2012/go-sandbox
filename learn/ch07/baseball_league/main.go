package main

import "fmt"

func main() {
	var league League
	league.AddTeam(Team{Name: "Inky Squid", Members: []string{"Ja", "Oma", "Kuba"}})
	league.AddTeam(Team{Name: "Angry Flounder", Members: []string{"Kumo", "Ono", "Leno"}})
	league.AddTeam(Team{Name: "Happy Clam", Members: []string{"Clara", "Boe", "Neo"}})
	league.AddTeam(Team{Name: "Pink Crab", Members: []string{"Boye", "Tomo", "Miu"}})

	league.MatchResult("Inky Squid", 16, "Angry Flounder", 9)
	league.MatchResult("Inky Squid", 6, "Happy Clam", 14)
	league.MatchResult("Inky Squid", 6, "Pink Crab", 4)

	fmt.Println(league)
	ranking := league.Ranking()
	for rank, name := range ranking {
		fmt.Printf("%-2d. %s\n", rank+1, name)
	}
}
