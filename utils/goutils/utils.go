package goutils

import (
	"bufio"
	"log/slog"
	"os"
)

func ReadFile(path string) ([]string,error) {
	file, err := os.Open(path)
	if err != nil {
		slog.Error("Failed to open input file", "error", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
