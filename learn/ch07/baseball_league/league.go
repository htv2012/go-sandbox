package main

import (
	"io"
	"slices"
	"strconv"
)

type League struct {
	Teams []Team
	Wins  map[string]int
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(ranker Ranker, writer io.Writer) {
	teams := ranker.Ranking()
	for rank, team := range teams {
		writer.Write([]byte(strconv.Itoa(rank + 1)))
		writer.Write([]byte(". "))
		writer.Write([]byte(team))
		writer.Write([]byte("\n"))
	}
}

func (l *League) AddTeam(t Team) {
	l.Teams = append(l.Teams, t)
}

func (l *League) MatchResult(team1Name string, team1Score int, team2Name string, team2Score int) {
	if l.Wins == nil {
		l.Wins = map[string]int{}
	}

	if team1Score > team2Score {
		l.Wins[team1Name] += 1
	} else if team2Score > team1Score {
		l.Wins[team2Name] += 1
	}
}

func (l *League) Ranking() []string {
	names := make([]string, 0, len(l.Wins))
	for name := range l.Wins {
		names = append(names, name)
	}

	slices.SortFunc(names, func(a, b string) int { return l.Wins[b] - l.Wins[a] })
	return names
}
