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

func parseStruct(line string) *game {
	divider := strings.Split(line, ":")
	//Parse game id
	gameId := strings.Split(divider[0], " ")[1]
	g := newGame()
	g.gameId = gameId
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

func isGamePossible(g *game) bool {
	const maxReds = 12
	const maxGreens = 13
	const maxBlues = 14

	fmt.Println("Checking game: ", *g)
	for i := 0; i < len(g.sets); i++ {
		if g.sets[i].numReds > maxReds || g.sets[i].numGreens > maxGreens || g.sets[i].numBlues > maxBlues {
			return false
		}
	}
	fmt.Println("Found possible game: ", *g)
	return true
}

func parteInput(input string) []game {
	lines := strings.Split(input, "\n")
	var games []game
	for i := 0; i < len(lines); i++ {
		game := parseStruct(lines[i])
		games = append(games, *game)
	}
	return games
}

func SolvePartOne(input string) int {
	fmt.Println("Solving part one")
	games := parteInput(input)
	var possibleGames []game
	for index := range games {
		if isGamePossible(&games[index]) {
			possibleGames = append(possibleGames, games[index])
		}
	}
	possibleGamesSum := 0
	for i := 0; i < len(possibleGames); i++ {
		sum, err := strconv.Atoi(possibleGames[i].gameId)
		if err != nil {
			panic(err)
		}
		possibleGamesSum += sum
		fmt.Println(sum)
	}
	return possibleGamesSum
}

func SolvePartTwo(input string) int {
	//TODO
	return -1
}
