package main

import (
	"fmt"

	"github.com/Irlandesee/advent_of_go/2023/day2"
	"github.com/Irlandesee/advent_of_go/2023/day3"
	"github.com/Irlandesee/advent_of_go/2023/day4"
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
	resPartOneInput := day2.SolvePartOne(util.ReadFile("2023/day2/inputs/inputs"))
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
	resPartTwoInput := day2.SolvePartTwo(util.ReadFile("2023/day2/inputs/inputs"))
	fmt.Println("Result might be: ", resPartTwoInput)
	fmt.Println("-------------END TEST PART TWO---------------")
}

func solveDayThree() {
	fmt.Println("-------------BEGIN TEST DAY 3 PART ONE---------------")
	resTestPartOne := day3.SolvePartOne(util.ReadFile("2023/day3/inputs/test"))
	if resTestPartOne != 4361 {
		fmt.Printf("Wrong result! Should be %d\n", resTestPartOne)
	} else {
		fmt.Printf("Correct result! Should be %d\n", resTestPartOne)
	}
	fmt.Println("-------------END TEST DAY 3 PART ONE---------------")
}

func solveDayFour() {
	fmt.Println("-------------BEGIN TEST DAY 4 PART ONE---------------")
	resTestPartOne := day4.SolvePartOne(util.ReadFile("2023/day4/inputs/test"))
	if resTestPartOne != 13 {
		fmt.Printf("Wrong result! Should be %d\n", resTestPartOne)
	} else {
		fmt.Printf("Correct result! Should be %d\n", resTestPartOne)
	}
	fmt.Println("-------------END TEST DAY 4 PART ONE---------------")
	resPartOne := day4.SolvePartOne(util.ReadFile("2023/day4/inputs/input"))
	fmt.Println("Res part one ", resPartOne)
	if resPartOne != 18653 {
		fmt.Printf("Wrong result! Should be %d\n", resPartOne)
	} else {
		fmt.Printf("Correct result! Should be %d\n", resPartOne)
	}
	fmt.Println("-------------BEGIN TEST DAY 4 PART TWO---------------")
	resTestPartTwo := day4.SolvePartTwo(util.ReadFile("2023/day4/inputs/test"))
	if resTestPartTwo != 30 {
		fmt.Printf("Wrong result! Should be %d, got: %d\n", 30, resTestPartTwo)
	} else {
		fmt.Printf("Correct result! Should be %d\n", resTestPartTwo)
	}
	fmt.Println("-------------END TEST DAY 4 PART TWO---------------")

}

func main() {
	solveDayFour()
}
