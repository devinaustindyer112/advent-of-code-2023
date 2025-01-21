package day_5_part_2

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"

	"adventofcode.com/2023/utils"
)

type MapEntry struct {
	DestinationStart int
	OriginStart      int
	RangeLength      int
}

func (m MapEntry) originEnd() int {
	return m.OriginStart + m.RangeLength
}

func main() {
	input, _ := os.ReadFile("./input.txt")
	day_5_part_2(string(input))
}

func day_5_part_2(input string) {

	regex := regexp.MustCompile(`\n\n`)
	sections := regex.Split(input, -1)
	utils.Assert(len(sections) == 8, fmt.Sprintf("Sections length incorrect: %d", len(sections)))

	value := parseSeeds(sections[0])

	for i := 1; i < len(sections); i++ {
		mapp := parseMap(sections[i])
		value = getDestinationMaps(value, mapp)
	}

	lowest := math.MaxInt
	for i := 0; i < len(value); i++ {
		if value[i].OriginStart < lowest {
			lowest = value[i].OriginStart
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
				RangeLength: fromMap.originEnd() - originEnd,
			}

			left := MapEntry{
				OriginStart: fromMap.OriginStart,
				RangeLength: originStart - fromMap.OriginStart,
			}

			newMap := toMaps[1:]

			// 1. All matches are on the left side
			if fromMap.OriginStart == originStart && fromMap.originEnd() > originEnd {
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
			OriginStart: utils.ParseInt(seedStrings[i]),
			RangeLength: utils.ParseInt(seedStrings[i+1]),
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
			OriginStart:      utils.ParseInt(entryValues[1]),
			DestinationStart: utils.ParseInt(entryValues[0]),
			RangeLength:      utils.ParseInt(entryValues[2]),
		})
	}

	return entriesMap
}
