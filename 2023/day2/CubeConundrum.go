package day2

import (
	"fmt"
	"strconv"
	"strings"
)

type set struct {
	numReds   int
	numGreens int
	numBlues  int
}

type game struct {
	gameId string
	sets   map[int]set
}

func newGame() *game {
	var g game
	g.sets = make(map[int]set)
	return &g
}

func cleanSets(sets []string) []string {
	var cleanedSets []string
	for i := 0; i < len(sets); i++ {
		s := strings.TrimSpace(sets[i])
		cleanedSets = append(cleanedSets, s)
	}
	return cleanedSets
}

func parseStruct(line string) game {
	divider := strings.Split(line, ":")
	//Parse game id
	gameId := strings.Split(divider[0], " ")[1]
	g := game{gameId: gameId}
	g.sets = make(map[int]set)
	numReds := 0
	numGreens := 0
	numBlues := 0
	//Parse sets
	divider[1] = strings.TrimSpace(divider[1])
	sets := strings.Split(divider[1], ";")
	for i := 0; i < len(sets); i++ {
		sets[i] = strings.TrimSpace(sets[i])
		s := strings.Split(sets[i], ",")
		for j := 0; j < len(s); j++ {
			t := strings.Split(strings.TrimSpace(s[j]), " ")
			if t[1] == "red" {
				reds, err := strconv.Atoi(t[0])
				if err != nil {
					panic(err)
				}
				numReds += reds
			} else if t[1] == "green" {
				greens, err := strconv.Atoi(t[0])
				if err != nil {
					panic(err)
				}
				numGreens += greens
			} else { //blue
				blues, err := strconv.Atoi(t[0])
				if err != nil {
					panic(err)
				}
				numBlues += blues
			}
		}
		gameSet := set{numReds, numGreens, numBlues}
		g.sets[i] = gameSet
		numReds = 0
		numGreens = 0
		numBlues = 0
	}
	return g
}

func parteInput(input string) []game {
	lines := strings.Split(input, "\n")
	var games []game
	for i := 0; i < len(lines); i++ {
		game := parseStruct(lines[i])
		games = append(games, game)
	}
	return games
}

func SolvePartOne(input string) int {
	fmt.Println("Solving part one")
	games := parteInput(input)
	for index := range games {
		fmt.Println(games[index])
	}
	return -1
}

func SolvePartTwo(input string) int {
	//TODO
	return -1
}
