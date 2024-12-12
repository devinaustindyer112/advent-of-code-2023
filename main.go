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
	day_5_part_2(string(input))
}

// For novel problems, niave solution first!
// Understand the problem fully before trying to complete it!
// Visualizing the problem can help tremendously

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
}

func getDestinationValues(fromMap []MapEntry, toMap []MapEntry) []MapEntry {

	// Think of it this wat

	// Origin: |     -----|
	// Map:
	// Origin: 		|-----     |
	// Destination 			|-----     |

	// There is a clear pattern we can take advantage of.
	// Either origin 1 or origin 2 will be a value that we can use.
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

func parseSeeds(section string) []MapEntry {
	seedStrings := regexp.MustCompile(`[0-9]+`).FindAllString(section, -1)
	var seeds []MapEntry

	for i := 0; i < len(seedStrings); i += 2 {
		seed := MapEntry{
			OriginStart: parseInt(seedStrings[i]),
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
