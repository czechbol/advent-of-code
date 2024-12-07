package parts

import (
	"log/slog"
	"math"
	"strconv"
	"strings"
)

func PartOne(input []string) (int, error) {
    safeCount := 0

    for _, line := range input {
        parts := strings.Fields(line)
        levels := make([]int, len(parts))

        for i, part := range parts {
            num, err := strconv.Atoi(part)
            if err != nil {
                slog.Error("Failed to parse number", "part", part, "error", err)
                return 0, err
            }
            levels[i] = num
        }

        if isSafePartOne(levels) {
            safeCount++
        }
    }

    return safeCount, nil
}

func isSafePartOne(levels []int) bool {
	increasing := false
	decreasing := false

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		absDiff := int(math.Abs(float64(diff)))
		if absDiff < 1 || absDiff > 3 {
			return false
		}
		if diff < 0 {
			decreasing = true
		}
		if diff > 0 {
			increasing = true
		}
	}

	// only one of the conditions should be true
	safe := increasing != decreasing
	return safe

}
