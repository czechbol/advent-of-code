package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	inputFilePath = "2023/03/input/input.txt"
)

type lineMatch struct {
	lineIdx         int
	symbol          string
	startIdx        int
	endIdx          int
	number          int
	adjacentMatches []lineMatch
}

func main() {
	lines := readFile(inputFilePath)

	//for _, line := range lines {
	//	fmt.Println(line)
	//}

	validNumbers := findValidNumbers(findMatches(lines))

	result := 0
	for _, match := range validNumbers {
		if match.number != -1 {
			fmt.Println(match.number)
			result += match.number

		} else if match.symbol == "*" && len(match.adjacentMatches) == 2 {
			result += match.adjacentMatches[0].number * match.adjacentMatches[1].number
			fmt.Println(match.adjacentMatches[0].number * match.adjacentMatches[1].number)

		}
	}
	fmt.Println("Result:", result) // 467835

}

func readFile(inputFilePath string) []string {
	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("Failed to open input file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func findMatches(lines []string) []lineMatch {
	re := regexp.MustCompile(`(?P<numbers>\d+)|(?P<symbols>[^\d.\w\n])`)
	lineMatches := make([]lineMatch, 0)

	for lineIdx, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		startIndex := 0
		for _, match := range matches {
			for i, value := range match {
				idx := strings.Index(line[startIndex:], value)
				if i == 1 && value != "" {
					num, err := strconv.Atoi(value)
					if err != nil {
						log.Fatalf("Failed to convert digits to number: %v\n", err)
					}
					lineMatches = append(lineMatches, lineMatch{
						number:   num,
						lineIdx:  lineIdx,
						startIdx: startIndex + idx,
						endIdx:   startIndex + idx + len(value) - 1,
					})
					startIndex += idx + len(value)
				} else if i == 2 && value != "" {
					lineMatches = append(lineMatches, lineMatch{
						number:   -1,
						symbol:   value,
						lineIdx:  lineIdx,
						startIdx: startIndex + idx,
						endIdx:   startIndex + idx + len(value) - 1,
					})
					startIndex += idx + len(value)
				}
			}
		}
	}
	return lineMatches
}

func findValidNumbers(lineMatches []lineMatch) []lineMatch {
	validNumbers := make([]lineMatch, 0)
	for _, match := range lineMatches {
		if match.number == -1 && match.symbol != "*" {
			continue
		} else if match.number == 114 {
			fmt.Println(match)
		}

		var valid bool
		for _, testMatch := range lineMatches {
			if match.number != -1 && testMatch.number != -1 {
				continue
			} else if match.symbol != "*" && testMatch.number == -1 {
				continue
			}
			if testMatch.lineIdx < match.lineIdx-1 {
				continue
			} else if testMatch.lineIdx > match.lineIdx+1 {
				break
			}
			// check if testMatch starts or ends at most one character before or after match
			if testMatch.endIdx < match.startIdx-1 {
				continue
			} else if testMatch.startIdx > match.endIdx+1 {
				continue
			}
			if match.symbol == "*" && testMatch.number != -1 {
				match.adjacentMatches = append(match.adjacentMatches, testMatch)
			}

			if match.number == -1 && match.symbol != "*" {
				continue
			}
			valid = true
		}
		if valid {
			validNumbers = append(validNumbers, match)
		}
	}
	fmt.Println(validNumbers)
	return validNumbers
}
