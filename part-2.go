package main

import (
	"regexp"
	"strconv"
)

func day_3_part_2(file []byte) {
	fileStr := string(file)
	symbolIndexes := indexRegex(fileStr, "[^[:space:]0-9.]+")
	partIndexes := indexRegex(fileStr, "[0-9]+")
	gearRatios := getGearRatios(symbolIndexes, partIndexes)
	sum := 0

	for i := 0; i < len(gearRatios); i += 2 {
		// Get next two gears, multiply and add
		partFirst := stringToInt(fileStr[gearRatios[i][0]:gearRatios[i][1]])
		partSecond := stringToInt(fileStr[gearRatios[i+1][0]:gearRatios[i+1][1]])
		sum += partFirst * partSecond
	}

	println(sum)
}

func indexRegex(search string, regexStr string) [][]int {
	regex, _ := regexp.Compile(regexStr)
	return regex.FindAllStringIndex(search, -1)
}

func getGearRatios(symbolIndexes [][]int, partIndexes [][]int) [][]int {
	gears := [][]int{}

	for i := 0; i < len(symbolIndexes); i++ {
		adjacent := [][]int{}

		for j := 0; j < len(partIndexes); j++ {
			if isAdjacent(symbolIndexes[i], partIndexes[j]) {
				adjacent = append(adjacent, partIndexes[j])
			}
		}
		if len(adjacent) == 2 {
			gears = append(gears, adjacent...)
		}
	}

	return gears
}

// Checks the previous, current and next line for adjacent symbols
func isAdjacent(symbolIndex []int, partIndex []int) bool {
	return symbolIndex[0] >= partIndex[0]-1-141 && symbolIndex[1] <= partIndex[1]+1-141 ||
		symbolIndex[0] >= partIndex[0]-1 && symbolIndex[1] <= partIndex[1]+1 ||
		symbolIndex[0] >= partIndex[0]-1+141 && symbolIndex[1] <= partIndex[1]+1+141
}

func stringToInt(str string) int {
	number, _ := strconv.Atoi(str)
	return number
}
