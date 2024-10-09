package main

import (
	"fmt"
	"github.com/Irlandesee/advent_of_go/2023/day2"
	"github.com/Irlandesee/advent_of_go/2023/util"
)

func solveDayTwo() {
	fmt.Println("-------------BEGIN TEST PART ONE---------------")
	resPartOne := day2.SolvePartOne(util.ReadFile("2023/day2/inputs/test"))
	testPartOneResult := 8
	if resPartOne != testPartOneResult {
		fmt.Printf("Wrong result! Should be %d, is %d\n", testPartOneResult, resPartOne)
	} else {
		fmt.Printf("Correct result! Should be %d\n", resPartOne)
	}
	fmt.Println("-------------END TEST PART ONE---------------")
	fmt.Println("-------------BEGIN PART ONE---------------")
	resPartOneInput := day2.SolvePartOne(util.ReadFile("2023/day2/inputs/input")) //wrong
	fmt.Println("Result might be: ", resPartOneInput)
	fmt.Println("-------------END PART ONE---------------")

	fmt.Println("-------------BEGIN TEST PART TWO---------------")
	resPartTwoTest := day2.SolvePartTwo(util.ReadFile("2023/day2/inputs/test"))
	testPartTwoResult := 2286
	if resPartTwoTest != testPartTwoResult {
		fmt.Printf("Wrong result! Should be %d, is %d\n", testPartTwoResult, resPartTwoTest)
	} else {
		fmt.Printf("Correct result! Should be %d\n", resPartTwoTest)
	}
	fmt.Println("-------------END TEST PART TWO---------------")
	fmt.Println("-------------BEGIN TEST PART TWO---------------")
	resPartTwoInput := day2.SolvePartTwo(util.ReadFile("2023/day2/inputs/input"))
	fmt.Println("Result might be: ", resPartTwoInput)
	fmt.Println("-------------END TEST PART TWO---------------")
}

func main() {
	solveDayTwo()
}
