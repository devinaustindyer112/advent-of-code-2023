package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Node struct {
	Type      string
	ChildType string
	Value     int
	Range     []int
	Children  []*Node
}

// I'm not sure this is the best solution, but I'm going to solve this using
// a graph and nodes.

func (n *Node) add(node *Node) {

	// Root node
	if n.Value < 0 {
		n.Children = append(n.Children, node)
	}

	for _, child := range n.Children {
		// Create function for checking if in range
		if node.Type == child.ChildType {
			child.Children = append(child.Children, node)
		}
		child.add(node)
	}
}

func (n *Node) isCompatible(node *Node) {

}

func (n *Node) countChildren() int {
	return len(n.Children)
}

func main() {
	input, _ := os.ReadFile("./input.txt")
	day_5_part_1(string(input))
}

func day_5_part_1(input string) {

	regex := regexp.MustCompile(`(?m)^\s*$`)
	sections := regex.Split(input, -1)
	assert(len(sections) == 8, fmt.Sprintf("Sections length incorrect: %d", len(sections)))

	root := Node{
		Value: -1,
	}

	mapElement(&root, sections[0])
	assert(root.countChildren() == 20, fmt.Sprintf("Children length incorrect: %d", root.countChildren()))
}

func mapElement(root *Node, input string) {
	inputSplit := strings.Split(input, ":")
	assert(len(inputSplit) == 2, fmt.Sprintf("Split input length incorrect: %d", len(inputSplit)))

	inputValues := strings.Split(strings.Trim(inputSplit[1], " "), " ")
	assert(len(inputValues) == 20, fmt.Sprintf("Input values length incorrect: %d", len(inputValues)))

	for _, inputValue := range inputValues {
		elementValue, _ := strconv.Atoi(inputValue)
		elementType := inputSplit[0]

		child := Node{
			Type:  elementType,
			Value: elementValue,
		}

		root.add(&child)
	}
}

func assert(condition bool, errMessage string) {
	if !condition {
		panic(errMessage)
	}
}
