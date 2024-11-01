package day_1

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Panic("error opening input file")
	}

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		first, _ := findFirst(scanner.Text())
		second, _ := findLast(scanner.Text())
		sum += first + second
	}

	print(sum)
}

var stringsToDigits = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func findFirst(str string) (int, error) {

	for i := 0; i < len(str); i++ {

		if str[i] >= 48 && str[i] <= 57 {
			digit, _ := strconv.Atoi(string(str[i]))
			return digit * 10, nil
		}

		for key, _ := range stringsToDigits {
			if strings.HasPrefix(str[i:], key) {
				return stringsToDigits[key] * 10, nil
			}
		}
	}

	return -1, errors.New("sucks to suck")
}

func findLast(str string) (int, error) {

	for i := len(str) - 1; i >= 0; i-- {

		if str[i] >= 48 && str[i] <= 57 {
			digit, _ := strconv.Atoi(string(str[i]))
			return digit, nil
		}

		for key, _ := range stringsToDigits {
			if strings.HasPrefix(str[i:], key) {
				return stringsToDigits[key], nil
			}
		}
	}

	return -1, errors.New("sucks to suck")
}
