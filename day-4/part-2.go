package day_4

import (
	"fmt"
	"slices"
)

func day_4_part_2(input string) {
	games := parseGames(input)
	assert(len(games) == 199, fmt.Sprintf("Did not parse games properly. Game count: %d", len(games)))

	sum := 0
	cardCount := calculateCards(games)
	for _, count := range cardCount {
		sum = sum + count
	}

	println(sum)
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
