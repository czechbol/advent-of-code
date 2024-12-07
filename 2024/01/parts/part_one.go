package parts

import (
	"fmt"
	"log/slog"
	"math"
	"slices"
)

func PartOne(input []string) (int, error) {
	col_one := make([]int, len(input))
	col_two := make([]int, len(input))

	for i, line := range input {
		var num_one, num_two int
		_, err := fmt.Sscanf(line, "%d   %d", &num_one, &num_two)
		if err != nil {
			slog.Error("Failed to parse line", "line", line, "error", err)
			return 0, err
		}

		col_one[i] = num_one
		col_two[i] = num_two
	}

	var result int

	slices.Sort(col_one)
	slices.Sort(col_two)

	for i := 0; i < len(col_one); i++ {
		result += int(math.Abs(float64(col_one[i]-col_two[i])))
	}

	return result, nil
}
