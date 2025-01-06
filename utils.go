package main

import (
	"fmt"
	"strconv"
)

func assert(condition bool, errMessage string) {

	if !condition {
		panic(errMessage)
	}

}

func parseInt(str string) int {
	value, err := strconv.Atoi(str)

	if err != nil {
		panic(fmt.Sprintf("Error parsing string to int64 %s", err.Error()))
	}

	return value
}
