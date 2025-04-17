package day4

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
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

	cardWinningNumbers = strings.TrimLeft(cardWinningNumbers, " \t\n\r")
	cardWinningNumbers = strings.TrimRight(cardWinningNumbers, " \t\n\r")
	cardNumbers = strings.TrimLeft(cardNumbers, " \t\n\r")
	cardNumbers = strings.TrimRight(cardNumbers, " \t\n\r")

	var re = regexp.MustCompile(`\s{2}`)
	cardWinningNumbers = re.ReplaceAllString(cardWinningNumbers, " ")
	cardNumbers = re.ReplaceAllString(cardNumbers, " ")

	winNumbers := strings.Split(cardWinningNumbers, " ")
	numbers := strings.Split(cardNumbers, " ")

	g := newGame(len(winNumbers), len(numbers))
	g.gameId = gameId
	for index := range winNumbers {
		g.winNumbers[index], err = strconv.Atoi(winNumbers[index])
		if err != nil {
			panic(err)
		}
	}
	for index := range numbers {
		g.numbers[index], err = strconv.Atoi(numbers[index])
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(g)
}

/*
*
A real game has 10 winning numbers
and 25 "normal" numbers
*/
func newGame(winNumbersLength int, numbersLength int) *Game {
	var g Game
	g.winNumbers = make([]int, winNumbersLength)
	g.numbers = make([]int, numbersLength)
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
