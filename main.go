package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
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

	var sum int

	for currentLine != "" {

		indexes := indexPartNumbers(currentLine)

		// Just pass in the string that we want to inspect for each case, then have a simple function that checks for the values
		for i := 0; i < len(indexes); i++ {
			if adjacentHorizontal(indexes[i][0], indexes[i][1], currentLine) {
				sum += stringToInt(currentLine[indexes[i][0]:indexes[i][1]])
				print(stringToInt(currentLine[indexes[i][0]:indexes[i][1]]))
			} else if adjacentVertical(indexes[i][0], indexes[i][1], previousLine) {
				sum += stringToInt(currentLine[indexes[i][0]:indexes[i][1]])
				print(stringToInt(currentLine[indexes[i][0]:indexes[i][1]]))
			} else if adjacentVertical(indexes[i][0], indexes[i][1], nextLine) {
				sum += stringToInt(currentLine[indexes[i][0]:indexes[i][1]])
				print(stringToInt(currentLine[indexes[i][0]:indexes[i][1]]))
			}

			print("-")
		}

		println("")

		scanner.Scan()
		previousLine = currentLine
		currentLine = nextLine
		nextLine = scanner.Text()
	}

	println(sum)
}

func indexPartNumbers(search string) [][]int {
	regex, _ := regexp.Compile("[0-9]+")
	return regex.FindAllStringIndex(search, -1)
}

func adjacentHorizontal(startIndex int, endIndex int, currentLine string) bool {
	if startIndex == 0 {
		return false
	}

	regex, _ := regexp.Compile("[^.]")

	if regex.MatchString(string(currentLine[startIndex-1])) {
		return true
	}

	if endIndex == len(currentLine) {
		return false
	}

	return regex.MatchString(string(currentLine[endIndex : endIndex+1]))
}

func adjacentVertical(startIndex int, endIndex int, currentLine string) bool {
	if currentLine == "" {
		return false
	}

	if startIndex != 0 {
		startIndex--
	}

	if endIndex != len(currentLine) {
		endIndex++
	}

	regex, _ := regexp.Compile("[^.]+")

	return regex.MatchString(string(currentLine[startIndex:endIndex]))
}

func stringToInt(str string) int {
	number, _ := strconv.Atoi(str)
	return number
}
