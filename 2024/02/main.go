package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/czechbol/advent-of-code/2024/02/parts"
	"github.com/czechbol/advent-of-code/utils/goutils"
)

const (
	inputDir = "2024/02/input/"
	inputFile = "input.txt"
)

func main() {
	goutils.SetLogger()

	start := time.Now()
	input, err := goutils.ReadFile(inputDir+inputFile)
	if err != nil {
		slog.Error("Failed to read input", "error", err)
		os.Exit(1)
	}
	slog.Info("Reading input", "took", time.Since(start))

	partOneStart := time.Now()
	partOneResult, err := parts.PartOne(input)
	if err != nil {
		slog.Error("Failed to execute Part One", "error", err)
		os.Exit(1)
	}
	slog.Info("Part One", "result",partOneResult, "took", time.Since(partOneStart))

	partTwoStart := time.Now()
	partTwoResult , err := parts.PartTwo(input)
	if err != nil {
		slog.Error("Failed to execute Part One", "error", err)
		os.Exit(2)
	}
	slog.Info("Part Two", "result", partTwoResult, "took", time.Since(partTwoStart))
}
