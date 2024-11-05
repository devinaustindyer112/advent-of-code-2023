package main

import (
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, _ := os.ReadFile("./input.txt")
	day_3_part_2(string(file))
}

// Pretty sure we can use a map and make this more effiecent

func day_3_part_2(file string) {
	symbolIndexes := indexRegex(file, "[^[:space:]0-9.]+")
	partIndexes := indexRegex(file, "[0-9]+")
	adjacent := getAdjecentParts(symbolIndexes, partIndexes)
	sum := 0

	for i := 0; i < len(adjacent); i++ {
		number := stringToInt(file[adjacent[i][0]:adjacent[i][1]])
		sum += number
	}

	println(sum)
}

func indexRegex(search string, regexStr string) [][]int {
	regex, _ := regexp.Compile(regexStr)
	return regex.FindAllStringIndex(search, -1)
}

func getAdjecentParts(symbolsIndex [][]int, partsIndex [][]int) [][]int {
	var adjacent [][]int

	for i := 0; i < len(partsIndex); i++ {
		for j := 0; j < len(symbolsIndex); j++ {
			if isAdjacent(symbolsIndex[j], partsIndex[i]) {
				adjacent = append(adjacent, partsIndex[i])
				break
			}
		}
	}

	return adjacent
}

func isAdjacent(symbolIndex []int, partIndex []int) bool {
	// Check previous line
	if symbolIndex[0] >= partIndex[0]-1-141 && symbolIndex[1] <= partIndex[1]+1-141 {
		return true
	}

	// Check current line
	if symbolIndex[0] >= partIndex[0]-1 && symbolIndex[1] <= partIndex[1]+1 {
		return true
	}

	// Check next line
	if symbolIndex[0] >= partIndex[0]-1+141 && symbolIndex[1] <= partIndex[1]+1+141 {
		return true
	}

	return false
}

func stringToInt(str string) int {
	number, _ := strconv.Atoi(str)
	return number
}
