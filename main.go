package main

import (
	"bufio"
	"os"
	"regexp"
)

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	day_3_part_1(scanner)
}

func day_3_part_1(scanner *bufio.Scanner) {
	var previousLine string
	scanner.Scan()
	currentLine := scanner.Text()
	scanner.Scan()
	nextLine := scanner.Text()

	for currentLine != "" {

		indexes := numberIndexes(currentLine)

		for i := 0; i < len(indexes); i++ {
			if adjacentLeft(indexes[i][0], currentLine) {
				println(currentLine[indexes[i][0]:indexes[i][1]])
			}
			if adjacentRight(indexes[i][1], currentLine) {
				println(currentLine[indexes[i][0]:indexes[i][1]])
			}
			if adjacentAbove(indexes[i][0], indexes[i][1], previousLine) {
				println(currentLine[indexes[i][0]:indexes[i][1]])
			}
			if adjacentBelow(indexes[i][0], indexes[i][1], nextLine) {
				println(currentLine[indexes[i][0]:indexes[i][1]])
			}
		}

		scanner.Scan()
		previousLine = currentLine
		currentLine = nextLine
		nextLine = scanner.Text()
	}
}

func numberIndexes(search string) [][]int {
	regex, _ := regexp.Compile("[0-9]+")
	return regex.FindAllStringIndex(search, -1)
}

func adjacentLeft(startIndex int, currentLine string) bool {
	if startIndex == 0 {
		return false
	}

	regex, _ := regexp.Compile("[^0-9.]+")

	return regex.MatchString(string(currentLine[startIndex-1]))
}

func adjacentRight(endIndex int, currentLine string) bool {
	if endIndex == len(currentLine) {
		return false
	}

	regex, _ := regexp.Compile("[^0-9.]+")

	return regex.MatchString(string(currentLine[endIndex]))
}

func adjacentAbove(startIndex int, endIndex int, previousLine string) bool {
	if previousLine == "" {
		return false
	}

	return false
}

func adjacentBelow(startIndex int, endIndex int, nextLine string) bool {
	return false
}
