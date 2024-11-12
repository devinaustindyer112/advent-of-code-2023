package main

import (
	"os"
	"strings"
    "fmt"
    "slices"
    "math"
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
   
    sum := 0.0
    for i:=0; i<len(games); i++ {
        points := calculatePoints(games[i])
        println(fmt.Sprintf("Game %d: %f", i+1, points))
        sum += points
    }

    println(fmt.Sprintf("%f", sum))
}

func parseGames(input string) []Game {
	games := []Game{}
	lines := strings.Split(input, "\n")
    assert(len(lines) == 200, fmt.Sprintf("Did not split input properly. Line count: %d", len(lines)))

    // We don't want to parse the last "line". It is empty.
    for i:=0; i<len(lines)-1; i++ {
        games = append(games, parseGame(lines[i]))
    }
    assert(len(games) == 199, fmt.Sprintf("Did not parse games properly. Game count: %d", len(games)))

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

// 1, 2, 4, 8, 16
// 2^i
func calculatePoints(game Game) float64 {
    matchingCount := 0  

    for i:=0; i < len(game.numbers); i++ {
        if  slices.Contains(game.winningNumbers, game.numbers[i]) {
            matchingCount++
        }
    }

    if matchingCount == 0 {
        return 0
    }
    
    return math.Pow(2, float64(matchingCount - 1))
}

func assert(condition bool, errMessage string) {
	if !condition {
		panic(errMessage)
	}
}
