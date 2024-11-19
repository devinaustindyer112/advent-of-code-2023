package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type Game struct {
	numbers        []string
	winningNumbers []string
}

func main() {
	input, _ := os.ReadFile("./input.txt")
	day_4_part_1(string(input))
}

// Use a map of game count?

func day_4_part_1(input string) {
	games := parseGames(input)
	assert(len(games) == 199, fmt.Sprintf("Did not parse games properly. Game count: %d", len(games)))

	sum := 0
	cardCount := calculateCards(games)
	for _, count := range cardCount {
		sum = sum + count
	}

	println(sum)
}

func parseGames(input string) []Game {
	games := []Game{}
	lines := strings.Split(input, "\n")
	assert(len(lines) == 200, fmt.Sprintf("Did not split input properly. Line count: %d", len(lines)))

	// We don't want to parse the last "line". It is empty.
	for i := 0; i < len(lines)-1; i++ {
		games = append(games, parseGame(lines[i]))
	}
	assert(len(games) == 199, fmt.Sprintf("Did not parse games properly. Game count: %d", len(games)))

	return games
}

func parseGame(inputLine string) Game {
	numbers := strings.Fields(inputLine[strings.Index(inputLine, ":")+1 : strings.Index(inputLine, "|")])
	winningNumbers := strings.Fields(inputLine[strings.Index(inputLine, "|")+1:])

	assert(len(numbers) == 10, fmt.Sprintf("Did not parse numbers correctly. Length: %d", len(numbers)))
	assert(len(winningNumbers) == 25, fmt.Sprintf("Did not parse winningNumbers properly. Length: %d", len(winningNumbers)))

	return Game{
		numbers,
		winningNumbers,
	}
}

func calculateCards(games []Game) [199]int {
	// Gotta use a map or array to index the amount to iterate
	cardCount := [199]int{}

	for i := 0; i < len(games); i++ {
		cardCount[i] = cardCount[i] + 1
		for j := 0; j < cardCount[i]; j++ {
			winCount := 0
			for k := 0; k < len(games[i].numbers); k++ {
				if slices.Contains(games[i].winningNumbers, games[i].numbers[k]) {
					winCount++
					cardCount[i+winCount] = cardCount[i+winCount] + 1
				}
			}
		}
	}

	for i, count := range cardCount {
		println(fmt.Sprintf("card %d has count %d", i+1, count))
	}

	return cardCount
}

func calculatePoints(game Game) float64 {
	matchingCount := 0

	for i := 0; i < len(game.numbers); i++ {
		if slices.Contains(game.winningNumbers, game.numbers[i]) {
			matchingCount++
		}
	}

	if matchingCount == 0 {
		return 0
	}

	return math.Pow(2, float64(matchingCount-1))
}

func assert(condition bool, errMessage string) {
	if !condition {
		panic(errMessage)
	}
}
