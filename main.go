package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

type MapEntry struct {
	DestinationStart int
	OriginStart      int
	RangeLength      int
}

func (m MapEntry) originEnd() int {
	return m.OriginStart + m.RangeLength
}

func (m MapEntry) destinationEnd() int {
	return m.DestinationStart + m.RangeLength
}

func main() {
	input, _ := os.ReadFile("./input.txt")
	day_5_part_2(string(input))
}

// Answer: 79753136 | Too high
// It works for single values, why not ranges?
// Need more tests
// I'm seeing negative numbers!!!!!!!!!!!!!!

func day_5_part_2(input string) {

	regex := regexp.MustCompile(`\n\n`)
	sections := regex.Split(input, -1)
	assert(len(sections) == 8, fmt.Sprintf("Sections length incorrect: %d", len(sections)))

	seedMap := parseSeeds(sections[0])
	assert(len(seedMap) == 10, fmt.Sprintf("Seeds map length incorrect: %d", len(seedMap)))

	// This can be converted into a loop
	seedToSoilMap := parseMap(sections[1])
	soilToFertilizerMap := parseMap(sections[2])
	fertilizerToWater := parseMap(sections[3])
	waterToLight := parseMap(sections[4])
	lightToTemperature := parseMap(sections[5])
	temperatureToHumidity := parseMap(sections[6])
	humidityToLocation := parseMap(sections[7])

	soil := getDestinationMaps(seedMap, seedToSoilMap)
	fertilizer := getDestinationMaps(soil, soilToFertilizerMap)
	water := getDestinationMaps(fertilizer, fertilizerToWater)
	light := getDestinationMaps(water, waterToLight)
	temp := getDestinationMaps(light, lightToTemperature)
	humidity := getDestinationMaps(temp, temperatureToHumidity)
	location := getDestinationMaps(humidity, humidityToLocation)

	lowest := math.MaxInt
	for i := 0; i < len(location); i++ {
		if location[i].OriginStart < lowest {
			lowest = location[i].OriginStart
		}
	}

	println(lowest)
}

func getDestinationMaps(fromMaps []MapEntry, toMaps []MapEntry) []MapEntry {

	var destinationMap []MapEntry

	for i := 0; i < len(fromMaps); i++ {
		destinationMap = append(destinationMap, getDestinationMap(fromMaps[i], toMaps)...)
	}

	return destinationMap
}

func getDestinationMap(fromMap MapEntry, toMaps []MapEntry) []MapEntry {

	var destinationMap []MapEntry

	// If no toMaps are provided, map to itself.
	if len(toMaps) == 0 {
		return []MapEntry{
			{
				OriginStart: fromMap.OriginStart,
				RangeLength: fromMap.RangeLength,
			},
		}
	}

	for i := 0; i < len(toMaps); i++ {
		originStart := max(fromMap.OriginStart, toMaps[i].OriginStart)
		originEnd := min(fromMap.OriginStart+fromMap.RangeLength, toMaps[i].OriginStart+toMaps[i].RangeLength)

		if originStart < originEnd {
			rangeLength := originEnd - originStart
			destination := toMaps[i].DestinationStart + originStart - toMaps[i].OriginStart

			destinationMap = append(destinationMap, MapEntry{
				OriginStart: destination,
				RangeLength: rangeLength,
			})

			right := MapEntry{
				OriginStart: originEnd + 1,
				RangeLength: fromMap.destinationEnd() - originEnd,
			}

			left := MapEntry{
				OriginStart: originStart,
				RangeLength: originStart - fromMap.OriginStart,
			}

			newMap := toMaps[1:]

			// 1. All matches are on the left side
			if fromMap.OriginStart == originStart && fromMap.destinationEnd() > originEnd {
				destinationMap = append(destinationMap, getDestinationMap(right, newMap)...)
				return destinationMap
			}

			// 2. All matches are on the right side
			if originStart > fromMap.OriginStart && originEnd == fromMap.originEnd() {
				destinationMap = append(destinationMap, getDestinationMap(left, newMap)...)
				return destinationMap
			}

			// 3. Matches are in the middle
			if originStart > fromMap.OriginStart && originEnd < fromMap.originEnd() {
				destinationMap = append(destinationMap, getDestinationMap(right, newMap)...)
				destinationMap = append(destinationMap, getDestinationMap(left, newMap)...)
				return destinationMap
			}

			return destinationMap
		}
	}

	return []MapEntry{
		{
			OriginStart: fromMap.OriginStart,
			RangeLength: fromMap.RangeLength,
		},
	}
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
			OriginStart:      parseInt(entryValues[1]),
			DestinationStart: parseInt(entryValues[0]),
			RangeLength:      parseInt(entryValues[2]),
		})
	}

	return entriesMap
}
