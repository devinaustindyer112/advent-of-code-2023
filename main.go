package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type MapEntry struct {
	DestinationStart int
	OriginStart      int
	RangeLength      int
}

func main() {
	input, _ := os.ReadFile("./input.txt")
	day_5_part_1(string(input))
}

// For novel problems, niave solution first!
// Understand the problem fully before trying to complete it

func day_5_part_1(input string) {
	regex := regexp.MustCompile(`(?m)^\s*$`)
	sections := regex.Split(input, -1)
	assert(len(sections) == 8, fmt.Sprintf("Sections length incorrect: %d", len(sections)))

	seeds := strings.Split(strings.Trim(strings.Split(sections[0], ":")[1], " "), " ")
	assert(len(seeds) == 20, fmt.Sprintf("Seeds length incorrect: %d", len(seeds)))

	parseMap(sections[1])

	// Start with a single map and getting the appropriate values

}

func parseMap(section string) {
	entriesList := strings.Split(strings.Trim(strings.Split(strings.Trim(section, "\n"), ":")[1], "\n"), "\n")
	assert(len(entriesList) == 23, fmt.Sprintf("Map values incorrect length %d", len(entriesList)))

	var entriesMap []MapEntry
	for _, entry := range entriesList {
		entryValues := strings.Split(entry, " ")
		entriesMap = append(entriesMap, MapEntry{
			DestinationStart: parseInt(entryValues[0]),
			OriginStart:      parseInt(entryValues[1]),
			RangeLength:      parseInt(entryValues[2]),
		})
	}
}

func parseInt(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		panic("error parsing string to int")
	}
	return value
}

func assert(condition bool, errMessage string) {
	if !condition {
		panic(errMessage)
	}
}
