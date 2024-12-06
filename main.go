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

	seeds := parseSeeds(sections[0])
	assert(len(seeds) == 20, fmt.Sprintf("Seeds length incorrect: %d", len(seeds)))

	seedToSoilMap := parseMap(sections[1])
	soilToFertilizerMap := parseMap(sections[2])

	soils := getSoils(seeds, seedToSoilMap)
	assert(len(soils) >= 0, fmt.Sprintf("Invalid soils length %d", soils))

	fertilizers := getSoils(soils, soilToFertilizerMap)

	println("soils:")
	for _, soil := range soils {
		println(soil)
	}

	println("----------------------")
	println("fertilizers")
	for _, fertilizer := range fertilizers {
		println(fertilizer)
	}

	// Start with a single map and getting the lowest value.
}

func getSoils(seeds []int, seedToSoilArray []MapEntry) []int {

	var soils []int
	for _, seed := range seeds {
		soil := getSoil(seed, seedToSoilArray)
		soils = append(soils, soil)
	}

	return soils
}

func getSoil(seed int, seedToSoilArray []MapEntry) int {

	for _, seedToSoil := range seedToSoilArray {

		if !isWithinRange(seed, seedToSoil) {
			continue
		}

		for i := 0; i < seedToSoil.RangeLength; i++ {
			if seedToSoil.OriginStart+i == seed {
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

func parseSeeds(section string) []int {
	seedStrings := regexp.MustCompile(`[0-9]+`).FindAllString(section, -1)
	var seedIntegers []int
	for _, seed := range seedStrings {
		seedIntegers = append(seedIntegers, parseInt(seed))
	}

	return seedIntegers
}

func parseMap(section string) []MapEntry {
	entriesList := strings.Split(strings.Trim(strings.Split(strings.Trim(section, "\n"), ":")[1], "\n"), "\n")
	var entriesMap []MapEntry

	for _, entry := range entriesList {
		entryValues := strings.Split(entry, " ")
		entriesMap = append(entriesMap, MapEntry{
			DestinationStart: parseInt(entryValues[0]),
			OriginStart:      parseInt(entryValues[1]),
			RangeLength:      parseInt(entryValues[2]),
		})
	}

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
