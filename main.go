package main

import (
	"fmt"
	"os"
	"regexp"
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

func day_5_part_2(input string) {

	regex := regexp.MustCompile(`\n\n`)
	sections := regex.Split(input, -1)
	assert(len(sections) == 8, fmt.Sprintf("Sections length incorrect: %d", len(sections)))

	seedMap := parseSeeds(sections[0])
	assert(len(seedMap) == 10, fmt.Sprintf("Seeds map length incorrect: %d", len(seedMap)))

	// This can be converted into a loop
	seedToSoilMap := parseMap(sections[1])
	soilToFertilizerMap := parseMap(sections[2])
	fertilizerToWaterMap := parseMap(sections[3])
	waterToLight := parseMap(sections[4])
	lightToTemperature := parseMap(sections[5])
	temperatureToHumidity := parseMap(sections[6])
	humidityToLocation := parseMap(sections[7])

	destinations := getDestinationMaps(seedMap, seedToSoilMap)
	destinations = getDestinationMaps(destinations, soilToFertilizerMap)
	destinations = getDestinationMaps(destinations, fertilizerToWaterMap)
	destinations = getDestinationMaps(destinations, waterToLight)
	destinations = getDestinationMaps(destinations, lightToTemperature)
	destinations = getDestinationMaps(destinations, temperatureToHumidity)
	destinations = getDestinationMaps(destinations, humidityToLocation)

	for i := 0; i < len(destinations); i++ {
		fmt.Printf(`Origin start: %d`, destinations[i].OriginStart)
		fmt.Printf(`Destination start: %d`, destinations[i].DestinationStart)
		fmt.Printf(`Range length: %d`, destinations[i].DestinationStart)
		fmt.Println()
	}
}

func getDestinationMaps(fromMaps []MapEntry, toMaps []MapEntry) []MapEntry {
	// I dont think this works anymore and probably and issue with get Destination maps. Its addding a new map each time one doesn exist. I think we've handled optimizing it enought. Now we just need to make sure it works appropriately.

	// This is definitely where we're spending most of our time.
	var destinationMaps []MapEntry
	for i := 0; i < len(fromMaps); i++ {
		destinationMap := getDestinationMap(fromMaps[i], toMaps)
		destinationMaps = append(destinationMaps, destinationMap)
	}

	return destinationMaps
}

// I think I can do this recursively. Will need to update this. This will return an array now.
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
		originEnd := min(fromMap.OriginStart+fromMap.RangeLength, toMaps[i].OriginStart+fromMap.RangeLength)

		if originStart < originEnd {
			rangeLength := originEnd - originStart
			destination := toMaps[i].DestinationStart + originStart - toMaps[i].OriginStart

			destinationMap = append(destinationMap, MapEntry{
				OriginStart: destination,
				RangeLength: rangeLength,
			})

			if fromMap.OriginStart < originStart && fromMap.OriginStart+fromMap.RangeLength > originEnd {
				// If the intersection in the middle
				// 	append(getDestinationMap(left-side))
				// 	append(getDestinationMap(right-side))
			}

			if fromMap.OriginStart < originStart && fromMap.OriginStart+fromMap.RangeLength <= originEnd {
				// If the intersection is on the right side
				//	append(getDestinationMap(left-side))
			}

			if fromMap.OriginStart < originStart && fromMap.OriginStart+fromMap.RangeLength > originEnd {
				// If the intersection is on the left side
				// 	append(getDestinationMap(right-side))
			}
		}
	}

	// if there are no matches then we should retur
	return []MapEntry{
		{
			OriginStart: fromMap.OriginStart,
			RangeLength: fromMap.RangeLength,
		},
	} // This still needs some work. How do I know when there is no mapping?
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
