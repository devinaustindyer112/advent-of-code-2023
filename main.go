package main

import (
	"os"
	"strings"
    "fmt"
)

type Game struct {
	numbers        []string
	winningNumbers []string
}

func main() {
	input, _ := os.ReadFile("./input.txt")
	day_4_part_1(string(input))
}

func day_4_part_1(input string) {
	games := parseGames(input)
	assert(len(games) == 199, fmt.Sprintf("Did not parse games properly. Game count: %d", len(games)))
}

func parseGames(input string) []Game {
	games := []Game{}
	lines := strings.Split(input, "\n")

    // We don't want to parse the last "line". It is empty.
    for i:=0; i<len(lines)-1; i++ {
        parseGame(lines[i])
    }

    assert(len(lines) == 200, fmt.Sprintf("Did not split input properly. Line count: %d", len(lines)))
    return games
}

func parseGame(inputLine string) Game {
    numbers := strings.Fields(inputLine[strings.Index(inputLine, ":") + 1:strings.Index(inputLine, "|")])
    winningNumbers := strings.Fields(inputLine[strings.Index(inputLine, "|") + 1:]) 
    assert(len(numbers) == 10, fmt.Sprintf("Did not parse numbers correctly. Length: %d", len(numbers)))
    assert(len(winningNumbers) == 25, fmt.Sprintf("Did not parse winningNumbers properly. Length: %d", len(winningNumbers)))

    return Game{
        numbers,
        winningNumbers,
    }
}

func assert(condition bool, errMessage string) {
	if !condition {
		panic(errMessage)
	}
}
