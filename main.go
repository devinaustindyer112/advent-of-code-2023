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
// Understand the problem fully before trying to complete it!

func day_5_part_1(input string) {
	regex := regexp.MustCompile(`(?m)^\s*$`)
	sections := regex.Split(input, -1)
	assert(len(sections) == 8, fmt.Sprintf("Sections length incorrect: %d", len(sections)))

	seeds := regexp.MustCompile(`[0-9]+`).FindAllString(sections[0], -1)
	assert(len(seeds) == 20, fmt.Sprintf("Seeds length incorrect: %d", len(seeds)))

	seedToSoilMap := parseMap(sections[1])

	soils := getSoils(seeds, seedToSoilMap)
	assert(len(soils) >= 0, fmt.Sprintf("Invalid soils length %d", soils))

	for _, soil := range soils {
		println(soil)
	}

	// Start with a single map and getting the lowest value.
}

func getSoils(seeds []string, seedToSoilArray []MapEntry) []int {

	var soils []int
	for _, seed := range seeds {
		soil := getSoil(seed, seedToSoilArray)
		soils = append(soils, soil)
	}

	return soils
}

func getSoil(seed string, seedToSoilArray []MapEntry) int {

	for _, seedToSoil := range seedToSoilArray {

		if !isWithinRange(parseInt(seed), seedToSoil) {
			continue
		}

		for i := 0; i < seedToSoil.RangeLength; i++ {
			if seedToSoil.OriginStart+i == parseInt(seed) {
				return seedToSoil.DestinationStart + i
			}
		}

	}

	return -1
}

func isWithinRange(seed int, mapEntry MapEntry) bool {

	if mapEntry.OriginStart <= seed && seed <= mapEntry.OriginStart+mapEntry.RangeLength {
		return true
	}

	return false
}

func parseMap(section string) []MapEntry {
	entriesList := strings.Split(strings.Trim(strings.Split(strings.Trim(section, "\n"), ":")[1], "\n"), "\n")
	assert(len(entriesList) == 23, fmt.Sprintf("Map entries incorrect length %d", len(entriesList)))
	var entriesMap []MapEntry

	for _, entry := range entriesList {
		entryValues := strings.Split(entry, " ")
		entriesMap = append(entriesMap, MapEntry{
			DestinationStart: parseInt(entryValues[0]),
			OriginStart:      parseInt(entryValues[1]),
			RangeLength:      parseInt(entryValues[2]),
		})
	}
	assert(len(entriesMap) == 23, fmt.Sprintf("Map entry slice incorrect length %d", len(entriesMap)))

	return entriesMap
}

func parseInt(str string) int {
	value, err := strconv.Atoi(str)

	if err != nil {
		panic(fmt.Sprintf("Error parsing string to int %s", err.Error()))
	}

	return value
}

func assert(condition bool, errMessage string) {

	if !condition {
		panic(errMessage)
	}

}
