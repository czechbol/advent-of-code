package parts

import (
	"log/slog"
	"strconv"
	"strings"
)

func PartTwo(input []string) (int, error) {
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

        isSafe, _ := isSafePartTwo(levels)

        if isSafe {
            safeCount++
        } else {
            for i := 0; i < len(levels); i++ {
                newLevels := removeElementFromSlice(levels, i)
                isSafe, _ := isSafePartTwo(newLevels)
                if isSafe {
                    safeCount++
                    break
                }
            }
        }
    }

    return safeCount, nil
}

func isSafePartTwo(levels []int) (bool, int) {
    diffMap := make(map[int]int)

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		absDiff := Abs(diff)
		if absDiff < 1 || absDiff > 3 {
            return false, i
		}
        diffMap[i] = diff
	}

    lastIncreasing := -1
    lastDecreasing := -1
    numIncreasing := 0

    for i, diff := range diffMap {
        if diff > 0 {
            numIncreasing++
            lastIncreasing = i
        } else if diff < 0 {
            lastDecreasing = i
        }
    }

    if numIncreasing == len(diffMap) || numIncreasing == 0 {
        return true, -1
    } else if numIncreasing == 1 {
        return false, lastIncreasing
    } else if numIncreasing == len(diffMap) - 1 {
        return false, lastDecreasing
    }

	return false, -1
}

func removeElementFromSlice(slice []int, index int) []int {
    newSlice := make([]int, len(slice))
    copy(newSlice, slice)
    return append(newSlice[:index], newSlice[index+1:]...)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
