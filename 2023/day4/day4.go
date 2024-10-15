package day4

import (
	"fmt"
	"strconv"
	"strings"
)

type game struct {
	gameId     int
	winNumbers []int
	numbers    []int
}

func parseGame(line string) {
	divider := strings.Split(line, ":")
	gameId, err := strconv.Atoi(strings.Split(divider[0], " ")[1])
	if err != nil {
		panic(err)
	}
	//parse numbers
	divider[1] = strings.TrimSpace(divider[1])
	cardWinningNumbers := strings.Split(divider[1], "|")[0]
	cardNumbers := strings.Split(divider[1], "|")[1]

	cardWinningNumbers = strings.TrimSpace(cardWinningNumbers)
	cardNumbers = strings.TrimSpace(cardNumbers)

	winNumbers := strings.Split(cardWinningNumbers, " ")
	numbers := strings.Split(cardNumbers, " ")

	g := newGame()
	g.gameId = gameId
	for index := range winNumbers {
		g.winNumbers[index], err = strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
	}
	for index := range numbers {
		fmt.Println(num)
		g.numbers[index], err = strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(g)
}

func newGame() *game {
	var g game
	g.winNumbers = make([]int, 5)
	g.numbers = make([]int, 8)
	return &g
}

func parseInput(input *string) {
	lines := strings.Split(*input, "\n")
	for i := 0; i < len(lines); i++ {
		parseGame(lines[i])
	}
}

func SolvePartOne(input string) int {
	parseInput(&input)
	return -1
}

func SolvePartTwo(input string) int {
	//TODO
	return -1
}
