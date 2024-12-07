package parts

import (
	"fmt"
	"regexp"
	"strconv"
)

func PartOne(input []string) (int, error) {
    result := 0
    re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

    for _, line := range input {
        matches := re.FindAllStringSubmatch(line, -1)
        for _, match := range matches {
            x, err1 := strconv.Atoi(match[1])
            y, err2 := strconv.Atoi(match[2])
            if err1 != nil || err2 != nil {
                return 0, fmt.Errorf("invalid number in match: %v", match)
            }
            result += x * y
        }
    }

    return result, nil
}
