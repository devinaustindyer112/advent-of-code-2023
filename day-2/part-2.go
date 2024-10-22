package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	sum:= 0.0

	for i:=1;  scanner.Scan(); i++ {
		game := parseGame(scanner.Text())
		power := game["red"] * game["blue"] * game["green"]
		sum += power
	}

	fmt.Printf("%f\n", sum)
}

func parseGame(line string) map[string]float64 {

	gameString := line[strings.Index(line, ":"):]
	rounds := strings.Split(gameString, ";")
	roundsMap := map[string]float64{}

	for _, round := range rounds {
		drawsMap := parseRound(round)
		roundsMap["red"] = math.Max(float64(roundsMap["red"]), float64(drawsMap["red"]))
		roundsMap["blue"] = math.Max(float64(roundsMap["blue"]), float64(drawsMap["blue"]))
		roundsMap["green"] = math.Max(float64(roundsMap["green"]), float64(drawsMap["green"]))
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
