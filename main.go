package main

import "os"

func main() {
	file, _ := os.ReadFile("./input.txt")
	day_3_part_2(file)
}
