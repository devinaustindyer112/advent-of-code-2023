package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	print("start day 2")
	
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("error opening file")
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		println(scanner.Text())
	}
}