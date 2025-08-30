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

type ScratchCard struct {
	numOfCopies       int
	game              Game
	scratchCardCopies []*Game
}

func printGame(game *Game) {
	fmt.Printf("[GAME ID] : %d\n", game.gameId)
	for index := range game.winNumbers {
		if index < len(game.winNumbers)-1 {
			fmt.Printf("[%d]", game.winNumbers[index])
		} else {
			fmt.Printf("[%d]\n", game.winNumbers[index])
		}
	}

	for index := range game.numbers {
		if index < len(game.numbers)-1 {
			fmt.Printf("[%d]", game.numbers[index])
		} else {
			fmt.Printf("[%d]\n", game.numbers[index])
		}
	}
	fmt.Println("============================")
}

func printScratchCard(scratch *ScratchCard) {
	fmt.Println("Scratch card Game ID: ", scratch.game.gameId)
	fmt.Println("Scratch card number of copies: ", scratch.numOfCopies)
	for index := range scratch.scratchCardCopies {
		if index < len(scratch.scratchCardCopies)-1 {
			fmt.Printf("[%d]", scratch.scratchCardCopies[index].gameId)
		} else {
			fmt.Printf("[%d]\n", scratch.scratchCardCopies[index].gameId)
		}
	}
}

func parseGame(line string) *Game {
	if line == "" {
		return nil
	}
	divider := strings.Split(line, ":")
	reGameId := regexp.MustCompile(`Card\s+(\d+)`)
	gameId, err := strconv.Atoi(reGameId.FindStringSubmatch(line)[1])
	if err != nil {
		panic(err)
	}

	divider[1] = strings.TrimSpace(divider[1])
	cardWinningNumbers := strings.Split(divider[1], "|")[0]
	cardNumbers := strings.Split(divider[1], "|")[1]

	cardWinningNumbers = strings.TrimLeft(cardWinningNumbers, " \t\n\r")
	cardWinningNumbers = strings.TrimRight(cardWinningNumbers, " \t\n\r")
	cardNumbers = strings.TrimLeft(cardNumbers, " \t\n\r")
	cardNumbers = strings.TrimRight(cardNumbers, " \t\n\r")

	re := regexp.MustCompile(`\s{2}`)
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
	return g
}

/*
*
Returns total sum of points for a given game
*/
func calcGame(g *Game) int {
	var firstMatch = true
	var sumOfPoints int
	for i := range g.winNumbers {
		for j := range g.numbers {
			if g.winNumbers[i] == g.numbers[j] {
				if firstMatch {
					sumOfPoints += 1
					firstMatch = false
				} else {
					sumOfPoints *= 2
				}
			}
		}
	}
	return sumOfPoints
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

func parseInput(input *string) (games []*Game) {
	lines := strings.Split(*input, "\n")
	games = []*Game{}
	for i := 0; i < len(lines); i++ {
		g := parseGame(lines[i])
		if g != nil {
			games = append(games, g)
		}
	}
	return games
}

func SolvePartOne(input string) int {
	games := parseInput(&input)
	totalSumOfPoints := 0
	for index := range games {
		printGame(games[index])
		totalSumOfPoints += calcGame(games[index])
	}
	return totalSumOfPoints
}

func extractCopies(game *Game) int {
	numberOfCopies := 0
	for i := range game.winNumbers {
		for j := range game.numbers {
			if game.winNumbers[i] == game.numbers[j] {
				numberOfCopies += 1
			}
		}
	}
	return numberOfCopies
}

func SolvePartTwo(input string) int {
	scratchCardList := make([]*ScratchCard, 0)
	games := parseInput(&input)
	for index := range games {
		numOfCopies := extractCopies(games[index])
		if numOfCopies > 0 {
			endIndex := index + numOfCopies + 1
			scratchCard := ScratchCard{
				numOfCopies:       numOfCopies,
				game:              *games[index],
				scratchCardCopies: games[index+1 : endIndex],
			}
			scratchCardList = append(scratchCardList, &scratchCard)
		}
	}

	totalNumberOfScratchCards := 0
	for index := range scratchCardList {
		printScratchCard(scratchCardList[index])
		totalNumberOfScratchCards += scratchCardList[index].numOfCopies
	}

	return totalNumberOfScratchCards
}
