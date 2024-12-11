package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type MapEntry struct {
	DestinationStart int
	OriginStart      int
	RangeLength      int
}

type SeedMap struct {
	Start       int
	RangeLength int
}

func main() {
	input, _ := os.ReadFile("./input.txt")
	day_5_part_2(string(input))
}

// For novel problems, niave solution first!
// Understand the problem fully before trying to complete it!

func day_5_part_2(input string) {

	testGetDestinationFromValue()

	regex := regexp.MustCompile(`\n\n`)
	sections := regex.Split(input, -1)
	assert(len(sections) == 8, fmt.Sprintf("Sections length incorrect: %d", len(sections)))

	seedMap := parseSeeds(sections[0])
	assert(len(seedMap) == 10, fmt.Sprintf("Seeds map length incorrect: %d", len(seedMap)))

	// Can probably be converted to a for loop
	seedToSoilMap := parseMap(sections[1])
	soilToFertilizerMap := parseMap(sections[2])
	fertilizerToWaterMap := parseMap(sections[3])
	waterToLight := parseMap(sections[4])
	lightToTemperature := parseMap(sections[5])
	temperatureToHumidity := parseMap(sections[6])
	humidityToLocation := parseMap(sections[7])

	print("done")

	// We can convert this to a loop. Use previous output as input.
	// This is SLOOWWW for part 2. Probably need to find a new way to solve this part.
	// Yeah, shes just dying on me.
	// I need to figure out a completely different strategy for this.

	// I don't think I need to store indiviual values. I can pass ranges. Once I have the final range, I can determine the lowest value.
	// Instead of passin indivual seeds, I can pass the ranges. So given seed range, return the appropriate map ranges
	// Iterating through the seeds is the biggest slowdown here, and comparing to each range.

	soils := getDestinationValues(seeds, seedToSoilMap)
	fertilizers := getDestinationValues(soils, soilToFertilizerMap)
	waters := getDestinationValues(fertilizers, fertilizerToWaterMap)
	lights := getDestinationValues(waters, waterToLight)
	temperatures := getDestinationValues(lights, lightToTemperature)
	humidities := getDestinationValues(temperatures, temperatureToHumidity)
	locations := getDestinationValues(humidities, humidityToLocation)

	println(slices.Min(locations))
}

func getDestinationRange() {

}

func getDestinationRanges() {

}

func getDestinationValues(originValues []int, toMap []MapEntry) []int {

	var destinationValues []int

	for _, originValue := range originValues {
		destinationValue := getDestinationValue(originValue, toMap)
		destinationValues = append(destinationValues, destinationValue)
	}

	return destinationValues
}

func getDestinationValue(originValue int, toMap []MapEntry) int {

	for _, to := range toMap {

		if !isWithinRange(originValue, to) {
			continue
		}

		// Updated logic
		return originValue - to.OriginStart + to.DestinationStart
	}

	return originValue
}

func isWithinRange(originValue int, mapEntry MapEntry) bool {

	if mapEntry.OriginStart <= originValue && originValue < mapEntry.OriginStart+mapEntry.RangeLength {
		return true
	}

	return false
}

func parseSeeds(section string) []SeedMap {
	seedStrings := regexp.MustCompile(`[0-9]+`).FindAllString(section, -1)
	var seeds []SeedMap

	for i := 0; i < len(seedStrings); i += 2 {
		seed := SeedMap{
			Start:       parseInt(seedStrings[i]),
			RangeLength: parseInt(seedStrings[i+1]),
		}
		seeds = append(seeds, seed)
	}

	return seeds
}

func parseMap(section string) []MapEntry {
	entriesList := strings.Split(strings.Trim(strings.Split(strings.Trim(section, "\n"), ":")[1], "\n"), "\n")

	var entriesMap []MapEntry

	for _, entry := range entriesList {
		entryValues := regexp.MustCompile(`[0-9]+`).FindAllString(entry, -1)
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
		panic(fmt.Sprintf("Error parsing string to int64 %s", err.Error()))
	}

	return value
}

func assert(condition bool, errMessage string) {

	if !condition {
		panic(errMessage)
	}

}

// Write some tests. First go round was no bueno
func testGetDestinationFromValue() {
	entryMap := []MapEntry{
		{
			DestinationStart: 679195301,
			OriginStart:      529385087,
			RangeLength:      505408118,
		},
	}

	value := getDestinationValue(763445965, entryMap)
	expected := 763445965 - entryMap[0].OriginStart + entryMap[0].DestinationStart

	println(expected)
	assert(value == expected, fmt.Sprintf("Expected: %d but got: %d", expected, value))
}
