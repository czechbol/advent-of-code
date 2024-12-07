package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	inputFilePath = "2023/06/input/input.txt"
)

type WinningStrategy struct {
	HoldLength      int
	DistanceCovered int
}

type Race struct {
	Number            int
	Length            int
	WinnerDistance    int
	WinningStrategies []WinningStrategy
}

func main() {
	races := readFile(inputFilePath)

	err := calculateWinningLengths(races)
	if err != nil {
		log.Fatalf("Failed to calculate winning lengths: %v\n", err)
	}

	result := 1
	for _, race := range races {
		result *= len(race.WinningStrategies)
	}

	log.Printf("Result: %d\n", result)

}

func readFile(inputFilePath string) []*Race {
	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("Failed to open input file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, strings.Replace(scanner.Text(), "  ", " ", -1))
	}

	lengthsStr := strings.TrimSpace(strings.TrimLeft(lines[0], "Time:"))
	distancesStr := strings.TrimSpace(strings.TrimLeft(lines[1], "Distance:"))

	length := strings.ReplaceAll(lengthsStr, " ", "")
	distances := strings.ReplaceAll(distancesStr, " ", "")

	lengthInt, err := strconv.Atoi(length)
	if err != nil {
		log.Fatalf("Failed to convert length to int: %v\n", err)
	}
	distanceInt, err := strconv.Atoi(distances)
	if err != nil {
		log.Fatalf("Failed to convert distance to int: %v\n", err)
	}

	races := make([]*Race, 0)

	races = append(races, &Race{
		Number:         1,
		Length:         lengthInt,
		WinnerDistance: distanceInt,
	})

	return races
}

func calculateWinningLengths(races []*Race) error {
	for _, race := range races {
		for j := 0; j < race.Length; j++ {
			holdLength := j
			distanceCovered := j * (race.Length - j)

			if distanceCovered > race.WinnerDistance {
				winningStrategy := WinningStrategy{
					HoldLength:      holdLength,
					DistanceCovered: distanceCovered,
				}
				race.WinningStrategies = append(race.WinningStrategies, winningStrategy)
			}
		}
	}
	return nil
}
