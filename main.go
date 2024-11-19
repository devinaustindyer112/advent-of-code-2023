package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Node struct {
	Value    string
	Range    []string
	Children []*Node
}

func (n *Node) add(child Node) {
	if len(n.Children) == 0 {
		n.Children = append(n.Children, &child)
	}
}

func (n *Node) countChildren() int {
	return len(n.Children)
}

func main() {
	input, _ := os.ReadFile("./input.txt")
	day_5_part_1(string(input))
}

// This seems like a graph problem
// Every entity maps only to a single entity. I don't believe we need to account for overlapping.

func day_5_part_1(input string) {
	// Parse seeds
	// Parse soil
	// Find lowest value for soil.
	// Implement the rest of the solution

	regex := regexp.MustCompile(`(?m)^\s*$`)
	sections := regex.Split(input, -1)
	assert(len(sections) == 8, fmt.Sprintf("Sections length incorrect: %d", len(sections)))

	root := Node{
		Value: "root",
	}

	mapSeeds(&root, sections[0])
	assert(root.countChildren() == 20, fmt.Sprintf("Children length incorrect: %d", root.countChildren()))
}

func mapSeeds(root *Node, input string) {
	inputSplit := strings.Split(input, ":")
	assert(len(inputSplit) == 2, fmt.Sprintf("Split input length incorrect: %d", len(inputSplit)))

	seeds := strings.Split(strings.Trim(inputSplit[1], " "), " ")
	assert(len(seeds) == 20, fmt.Sprintf("Seeds length incorrect: %d", len(seeds)))

	for _, seed := range seeds {
		child := Node{
			Value: seed,
		}
		root.add(child)
	}
}

func mapSoil(root *Node, input string) {
	inputSplit := strings.Split(input, ":")
	assert(len(inputSplit) == 2, fmt.Sprintf("Split input length incorrect: %d", len(inputSplit)))

	regex := regexp.MustCompile(`\n$`)
	soilMap := regex.Split(inputSplit[1], -1)

}

func assert(condition bool, errMessage string) {
	if !condition {
		panic(errMessage)
	}
}
