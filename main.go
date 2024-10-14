package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {	
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("error opening file")
	}

	scanner := bufio.NewScanner(file)

	sum:= 0
	for i:=1;  scanner.Scan(); i++ {

		// The issue is that I'm looking at the total of the game, not each round
		game := parseGame(scanner.Text())
		if isValid(game) {

			println(i)

			sum += i
		}
	}

	print(sum)
}

func parseGame(line string) map[string]int {

	gameString := line[strings.Index(line, ":"):]
	rounds := strings.Split(gameString, ";")
	roundsMap := map[string]int{}

	for _, round := range rounds {
		drawsMap := parseRound(round)
		roundsMap["red"] = roundsMap["red"] + drawsMap["red"]
		roundsMap["blue"] = roundsMap["blue"] + drawsMap["blue"]
		roundsMap["green"] = roundsMap["green"] + drawsMap["green"]
	}

	return roundsMap
}

func parseRound(round string) map[string]int  {

	draws := strings.Split(round, ",")
	drawsMap := map[string]int {}

	colorRegex, _ := regexp.Compile("red|green|blue")
	numberRegex, _ := regexp.Compile("[0-9]+")

	for _, draw := range draws {
		color := colorRegex.FindString(draw)
		count, _ := strconv.Atoi(numberRegex.FindString(draw))
		drawsMap[color] = drawsMap[colorRegex.FindString(draw)] + count;
	}

	return drawsMap
}

func isValid(game map[string]int) bool {
	return game["red"] <= 12 && game["green"] <= 13 && game["blue"] <= 14 
}

// Which games possible with only 12 red cubes, 13 green cubes, and 14 blue cubes?
