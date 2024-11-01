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

	sum := 0

	for j := 0; currentLine != ""; j++ {

		indexes := indexPartNumbers(currentLine)

		// Just pass in the string that we want to inspect for each case, then have a simple function that checks for the values
		for i := 0; i < len(indexes); i++ {
			if adjacentHorizontal(indexes[i][0], indexes[i][1], currentLine) ||
				adjacentVertical(indexes[i][0], indexes[i][1], previousLine) ||
				adjacentVertical(indexes[i][0], indexes[i][1], nextLine) {

				number := stringToInt(string(currentLine[indexes[i][0]:indexes[i][1]]))
				sum = sum + number

				if j == 139 {
					println(number)
				}

			}
		}
		// println(" Done")

		previousLine = currentLine
		currentLine = nextLine
		scanner.Scan()
		nextLine = scanner.Text()

	}

	println(sum)

}

func indexPartNumbers(search string) [][]int {
	regex, _ := regexp.Compile("[0-9]+")
	return regex.FindAllStringIndex(search, -1)
}

// Something is wrong with these
func adjacentHorizontal(startIndex int, endIndex int, currentLine string) bool {
	if startIndex != 0 {
		startIndex--
	}

	if endIndex != len(currentLine) {
		endIndex++
	}

	regex, _ := regexp.Compile("[^0-9.]+")

	return regex.MatchString(string(currentLine[startIndex:endIndex]))
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

	regex, _ := regexp.Compile("[^0-9.]+")

	return regex.MatchString(string(currentLine[startIndex:endIndex]))
}

func stringToInt(str string) int {
	number, _ := strconv.Atoi(str)
	return number
}
