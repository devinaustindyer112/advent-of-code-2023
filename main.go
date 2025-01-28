package main

import (
	"os"
	"regexp"
)

func main() {
	input, _ := os.ReadFile("./input.txt")
	day_6_part_1(string(input))
}

// I need to find the values where (time - holding) * holding > distance
// Lowest value and the highest

// Naive would be to:
// 1. Iterate from the bottom, stop on first match. This is lower bound.
// 2. Iterate from the top, stop on first match. This is upper bound.
// 3. Use the difference

func day_6_part_1(input string) {

	regex := regexp.MustCompile(`\n\n`)
	sections := regex.Split(input, -1)

}
