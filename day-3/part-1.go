package day_3

import (
	"regexp"
	"strconv"
)

func day_3_part_1(file []byte) {
	fileStr := string(file)
	symbolIndexes := indexRegex(fileStr, "[^[:space:]0-9.]+")
	partIndexes := indexRegex(fileStr, "[0-9]+")
	parts := getParts(symbolIndexes, partIndexes)
	sum := 0

	for i := 0; i < len(parts); i++ {
		number := stringToInt(fileStr[parts[i][0]:parts[i][1]])
		sum += number
	}

	println(sum)
}

func indexRegex(search string, regexStr string) [][]int {
	regex, _ := regexp.Compile(regexStr)
	return regex.FindAllStringIndex(search, -1)
}

func getParts(symbolIndexes [][]int, partsIndex [][]int) [][]int {
	var adjacent [][]int

	for i := 0; i < len(partsIndex); i++ {
		for j := 0; j < len(symbolIndexes); j++ {
			if isAdjacent(symbolIndexes[j], partsIndex[i]) {
				adjacent = append(adjacent, partsIndex[i])
				break
			}
		}
	}

	return adjacent
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
