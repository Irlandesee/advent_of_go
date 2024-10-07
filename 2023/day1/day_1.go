package day1

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// ReadFile -> Read the whole file and return it as a string
func ReadFile(input string) string {
	contents, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}
	return string(contents)
}

func CleanString(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Trim(s, "\x00")
	return s
}

func LinkStrings(first, last string) string {
	return first + last
}

func FindFirstLast(s string) (string, string) {
	s = CleanString(s)
	first, last := "", ""

	// Find the first digit
	for i := 0; i < len(s); i++ {
		if _, err := strconv.Atoi(string(s[i])); err == nil {
			first = string(s[i])
			break
		}
	}

	// Find the last digit
	for i := 0; i < len(s); i++ {
		if _, err := strconv.Atoi(string(s[i])); err == nil {
			last = string(s[i])
		}
	}
	return first, last
}

func Solve(input string) int {
	in := ReadFile(input)
	lines := strings.Split(in, "\n")
	totalSum := 0
	for index := 0; index < len(lines); index++ {
		fmt.Println("Parsing: ", lines[index])
		first, last := FindFirstLast(lines[index])
		num, err := strconv.Atoi(LinkStrings(first, last))
		if err != nil {
			panic(err)
		}
		totalSum += num
	}
	return totalSum
}

func SolvePartTwo(fileName string) int {
	in := ReadFile(fileName)
	lines := strings.Split(in, "\n")
	regex := "(one|two|three|four|five|six|seven|eight|nine)"
	for index := 0; index < len(lines); index++ {
		fmt.Println("Parsing: ", lines[index])
		r, err := regexp.Compile(regex)
		if err != nil {
			panic(err)
		}
		fmt.Println(r.FindAllString(lines[index], -1))
	}
	return -1
}
