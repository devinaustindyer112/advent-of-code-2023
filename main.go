package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	input, _ := os.ReadFile("./input.txt")
	day_5_part_1(string(input))
}

// NAIVE SOLUTION FIRST DUMMY!! THEN YOU CAN OPTOMIZE

func day_5_part_1(input string) {
	regex := regexp.MustCompile(`(?m)^\s*$`)
	sections := regex.Split(input, -1)
	assert(len(sections) == 8, fmt.Sprintf("Sections length incorrect: %d", len(sections)))

	sectionSplit := strings.Split(sections[0], ":")
	assert(len(sectionSplit) == 2, fmt.Sprintf("Split input length incorrect: %d", len(sectionSplit)))

	seeds := strings.Split(strings.Trim(sectionSplit[1], " "), " ")
	assert(len(seeds) == 20, fmt.Sprintf("Input values length incorrect: %d", len(seeds)))

	parseMap(sections[1])
}

func parseMap(input string) {
	inputSplit := strings.Split(input, ":")
	assert(len(inputSplit) == 2, fmt.Sprintf("Split input length incorrect: %d", len(inputSplit)))

	// Need to trim withspace

	values := strings.Split(inputSplit[1], "\n")
	assert(len(values) == 23, fmt.Sprintf("Map values incorrect length %d", len(values)))
}

func assert(condition bool, errMessage string) {
	if !condition {
		panic(errMessage)
	}
}
