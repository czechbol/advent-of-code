package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	inputFilePath = "2023/02/input/input.txt"
	redCubes      = 12
	greenCubes    = 13
	blueCubes     = 14
)

type Set struct {
	Blue  int
	Red   int
	Green int
}

type Game struct {
	ID         int
	Sets       []Set
	MinimumSet *Set
}

func main() {
	games, err := parseFile(inputFilePath)
	if err != nil {
		fmt.Printf("Failed to parse file: %v\n", err)
		return
	}

	validSum := 0
	minSetPower := 0
	for _, game := range games {
		if isValidGame(game) {
			validSum += game.ID
		}
		minSetPower += game.MinimumSet.Blue * game.MinimumSet.Red * game.MinimumSet.Green
	}

	fmt.Println("Valid sum:", validSum)
	fmt.Println("Minimum set power:", minSetPower)
}

func parseFile(filePath string) ([]Game, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	games := []Game{}
	currentGame := Game{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		gameNumber, err := parseGameNumber(line)
		if err != nil {
			return nil, fmt.Errorf("failed to parse game number: %w", err)
		}

		currentGame = Game{
			ID: gameNumber,
		}

		line = strings.Split(line, ":")[1]
		sets := strings.Split(line, ";")

		err = parseSets(&currentGame, sets)
		if err != nil {
			return nil, fmt.Errorf("failed to parse sets: %w", err)
		}

		currentGame.MinimumSet = getMinimumSet(currentGame)

		games = append(games, currentGame)
	}

	return games, nil
}

func parseGameNumber(line string) (int, error) {
	gameNumberStr := line[len("Game "):strings.Index(line, ":")]
	gameNumber, err := strconv.Atoi(gameNumberStr)
	if err != nil {
		return 0, err
	}
	return gameNumber, nil
}

func parseSets(currentGame *Game, sets []string) error {
	for _, set := range sets {
		set = strings.TrimSpace(set)
		if set == "" {
			continue
		}
		setParts := strings.Split(set, ", ")
		if len(setParts) > 3 {
			return fmt.Errorf("invalid set: %s", set)
		}
		blue, red, green, err := parseSetParts(setParts)
		if err != nil {
			return fmt.Errorf("failed to parse set: %w", err)
		}
		currentGame.Sets = append(currentGame.Sets, Set{
			Blue:  blue,
			Red:   red,
			Green: green,
		})
	}
	return nil
}

func parseSetParts(setParts []string) (int, int, int, error) {
	var blue, red, green int
	var err error
	for i := range setParts {
		if strings.Contains(setParts[i], "blue") {
			blue, err = strconv.Atoi(strings.Split(setParts[i], " ")[0])
			if err != nil {
				return 0, 0, 0, err
			}
		} else if strings.Contains(setParts[i], "red") {

			red, err = strconv.Atoi(strings.Split(setParts[i], " ")[0])
			if err != nil {
				return 0, 0, 0, err
			}
		} else if strings.Contains(setParts[i], "green") {

			green, err = strconv.Atoi(strings.Split(setParts[i], " ")[0])
			if err != nil {
				return 0, 0, 0, err
			}
		}
	}
	return blue, red, green, nil
}

func getMinimumSet(game Game) *Set {
	minimumSet := Set{
		Blue:  0,
		Red:   0,
		Green: 0,
	}
	for _, set := range game.Sets {
		if set.Blue > minimumSet.Blue {
			minimumSet.Blue = set.Blue
		}
		if set.Red > minimumSet.Red {
			minimumSet.Red = set.Red
		}
		if set.Green > minimumSet.Green {
			minimumSet.Green = set.Green
		}
	}
	return &minimumSet
}

func isValidGame(game Game) bool {
	for _, set := range game.Sets {
		if set.Blue > blueCubes || set.Red > redCubes || set.Green > greenCubes {
			return false
		}
	}
	return true
}
