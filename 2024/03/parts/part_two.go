package parts

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

type CustomMatch struct {
    matchType string
    index int
    match []string
}

func PartTwo(input []string) (int, error) {
    result := 0
    mulEnabled := true
    doRe := regexp.MustCompile(`do\(\)`)
    dontRe := regexp.MustCompile(`don't\(\)`)
    mulRe := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

    for _, line := range input {
        doMatchesPos := doRe.FindAllStringSubmatchIndex(line, -1)
        dontMatchesPos := dontRe.FindAllStringSubmatchIndex(line, -1)
        mulMatchesPos := mulRe.FindAllStringSubmatchIndex(line, -1)

        mulMatches := mulRe.FindAllStringSubmatch(line, -1)

        matches := make([]CustomMatch, 0, len(doMatchesPos) + len(dontMatchesPos) + len(mulMatchesPos))
        for _, match := range doMatchesPos {
            matches = append(matches, CustomMatch{"do", match[0], nil})
        }
        for _, match := range dontMatchesPos {
            matches = append(matches, CustomMatch{"dont", match[0], nil})
        }
        for i, match := range mulMatchesPos {
            matches = append(matches, CustomMatch{"mul", match[0], mulMatches[i]})
        }

        // sort matches by index
        sort.Slice(matches, func(i, j int) bool {
            return matches[i].index < matches[j].index
          })

        // process matches
        for _, match := range matches {
            switch match.matchType {
            case "do":
                mulEnabled = true
            case "dont":
                mulEnabled = false
            case "mul":
                if mulEnabled {
                    x, err1 := strconv.Atoi(match.match[1])
                    y, err2 := strconv.Atoi(match.match[2])
                    if err1 != nil || err2 != nil {
                        return 0, fmt.Errorf("invalid number in match: %v", match)
                    }
                    result += x * y
                }
            }
        }


    }

    return result, nil
}
