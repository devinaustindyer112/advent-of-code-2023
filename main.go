package main

import (
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

// For novel problems, niave solution first!
// Understand the problem fully before trying to complete it!
// Visualizing the problem can help tremendously

func day_5_part_2(input string) {

	/*
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
	*/
}

// This logic is the right start, but it's incorrect.

// Matching scenarios

// |   --    |
//    |--|

//    |--|
// |   --    |

//    |--    |
// |   --|

// |     --|
//      |--     |

// |          |
//              |          |

// Matchin origin values are always the minimum origin and the maximum origin within the bounds
// Destination can be calculated once we have the origin values. I belive we can use the to origins with the range length to determine destination values. Probably can check some condition to determine if it is valid

func getDestinationValue(fromMap MapEntry, toMap MapEntry) MapEntry {

	originStart := max(fromMap.OriginStart, toMap.OriginStart)
	originEnd := min(fromMap.OriginStart+fromMap.RangeLength, toMap.OriginStart+fromMap.RangeLength)

	// This needs to be updated. There is a default value.
	if originStart >= originEnd {
		return MapEntry{
			OriginStart: fromMap.OriginStart,
			RangeLength: 1,
		}
	}

	rangeLength := originEnd - originStart
	destination := toMap.DestinationStart + originStart - toMap.OriginStart

	// Now i need to determine range length and destination start. Should be more arithmetic
	// This will be the destination values of this map, but will be origin for the next.
	return MapEntry{
		OriginStart: destination,
		RangeLength: rangeLength,
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
			DestinationStart: parseInt(entryValues[0]),
			OriginStart:      parseInt(entryValues[1]),
			RangeLength:      parseInt(entryValues[2]),
		})
	}

	return entriesMap
}
