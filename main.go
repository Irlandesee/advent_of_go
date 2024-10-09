package main

import (
	"fmt"
	"github.com/Irlandesee/advent_of_go/2023/day2"
	"github.com/Irlandesee/advent_of_go/2023/util"
)

func main() {
	resPartOne := day2.SolvePartOne(util.ReadFile("2023/day2/inputs/test"))
	testPartOneResult := 8
	if resPartOne != testPartOneResult {
		fmt.Printf("Wrong result! Should be %d, is %d\n", testPartOneResult, resPartOne)
	} else {
		fmt.Printf("Correct result! Should be %d\n", resPartOne)
	}
	resPartOneInput := day2.SolvePartOne(util.ReadFile("2023/day2/inputs/input")) //wrong
	fmt.Println("Result might be: ", resPartOneInput)

}
