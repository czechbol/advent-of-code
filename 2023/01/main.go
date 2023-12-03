package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	inputFilePath = "2023/01/input/input.txt"
)

var (
	wordDigits = map[string]string{
		"zero":  "ze0ro",
		"one":   "on1e",
		"two":   "tw2o",
		"three": "th3ree",
		"four":  "fo4ur",
		"five":  "fi5ve",
		"six":   "si6x",
		"seven": "se7ven",
		"eight": "ei8ght",
		"nine":  "ni9ne",
	}
)

func main() {
	file, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Printf("Failed to open input file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()

		line = replaceWordsWithDigits(line)

		lineDigits := extractDigits(line)

		num, err := strconv.Atoi(string(lineDigits))
		if err != nil {
			fmt.Printf("Failed to convert digits to number: %v\n", err)
			return
		}

		result += num
	}

	fmt.Println("Result:", result)
}

func replaceWordsWithDigits(line string) string {
	for key, value := range wordDigits {
		if strings.Contains(line, key) {
			line = strings.ReplaceAll(line, key, value)
		}
	}
	return line
}

func extractDigits(line string) []rune {
	lineDigits := make([]rune, 2)
	numCount := 0
	for _, c := range line {
		if c >= '0' && c <= '9' {
			if numCount == 0 {
				lineDigits[0], lineDigits[1] = c, c
			} else {
				lineDigits[1] = c
			}
			numCount++
		}
	}
	return lineDigits
}
