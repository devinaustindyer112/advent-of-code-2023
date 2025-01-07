package day_5

import (
	"fmt"
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

// For novel problems, niave solution first!
// Understand the problem fully before trying to complete it!

func day_5_part_1(input string) {

	testGetDestinationFromValue()

	regex := regexp.MustCompile(`\n\n`)
	sections := regex.Split(input, -1)
	assert(len(sections) == 8, fmt.Sprintf("Sections length incorrect: %d", len(sections)))

	seeds := parseSeeds(sections[0])
	assert(len(seeds) == 20, fmt.Sprintf("Seeds length incorrect: %d", len(seeds)))

	// We can do this in one go
	seedToSoilMap := parseMap(sections[1])
	soilToFertilizerMap := parseMap(sections[2])
	fertilizerToWaterMap := parseMap(sections[3])
	waterToLight := parseMap(sections[4])
	lightToTemperature := parseMap(sections[5])
	temperatureToHumidity := parseMap(sections[6])
	humidityToLocation := parseMap(sections[7])

	// We can convert this to a loop. Use previous output as input.
	soils := getDestinationValues(seeds, seedToSoilMap)
	fertilizers := getDestinationValues(soils, soilToFertilizerMap)
	waters := getDestinationValues(fertilizers, fertilizerToWaterMap)
	lights := getDestinationValues(waters, waterToLight)
	temperatures := getDestinationValues(lights, lightToTemperature)
	humidities := getDestinationValues(temperatures, temperatureToHumidity)
	locations := getDestinationValues(humidities, humidityToLocation)

	println(slices.Min(locations))
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

		// I can do this much more efficiently
		for i := 0; i < to.RangeLength; i++ {

			if to.OriginStart+i == originValue {
				return to.DestinationStart + i
			}
		}
	}

	return originValue
}

func isWithinRange(originValue int, mapEntry MapEntry) bool {

	if mapEntry.OriginStart <= originValue && originValue < mapEntry.OriginStart+mapEntry.RangeLength {
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
		panic(fmt.Sprintf("Error parsing string to int %s", err.Error()))
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
