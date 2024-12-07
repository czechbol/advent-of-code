package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	inputFilePath = "2023/04/input/input.txt"
)

type Card struct {
	Number         int
	WinningNumbers []int
	YourNumbers    []int
	Score          int
	Copies         int
}

func (c Card) String() string {
	return fmt.Sprintf("Card: {Number: %d, WinningNumbers: %v, YourNumbers: %v, Score: %d, Copies: %d}", c.Number, c.WinningNumbers, c.YourNumbers, c.Score, c.Copies)
}

func main() {
	lines := readFile(inputFilePath)

	cards := parseInput(lines)

	calculateScoreCopies(cards)

	score := 0
	copies := 0
	for _, card := range cards {

		fmt.Println(card)
		score += card.Score
		copies += 1 + card.Copies
	}

	fmt.Printf("Score: %d\n", score)
	fmt.Printf("Copies: %d\n", copies)
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

func parseInput(lines []string) []*Card {
	cards := make([]*Card, len(lines))

	for i, line := range lines {
		cardData := strings.Split(line, ":")
		cardNumber, err := strconv.Atoi(strings.TrimLeft(cardData[0], "Card "))
		if err != nil {
			log.Fatalf("Failed to convert card number to int: %v", err)
		}
		numbersData := strings.Split(cardData[1], "|")

		winningNumbersStr := strings.Fields(strings.TrimSpace(numbersData[0]))
		yourNumbersStr := strings.Fields(strings.TrimSpace(numbersData[1]))

		winningNumbers := make([]int, len(winningNumbersStr))
		yourNumbers := make([]int, len(yourNumbersStr))

		for j, numStr := range winningNumbersStr {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatalf("Failed to convert winning number to int: %v", err)
			}
			winningNumbers[j] = num
		}

		for j, numStr := range yourNumbersStr {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatalf("Failed to convert your number to int: %v", err)
			}
			yourNumbers[j] = num
		}

		card := &Card{
			Number:         cardNumber,
			WinningNumbers: winningNumbers,
			YourNumbers:    yourNumbers,
			Copies:         0,
			Score:          0,
		}

		cards[i] = card
	}

	return cards
}

func calculateScoreCopies(cards []*Card) {
	for i, card := range cards {
		card.Score = 0

		for _, yourNum := range card.YourNumbers {
			for _, winningNum := range card.WinningNumbers {
				if yourNum == winningNum {
					card.Score++
				}
			}
		}
		for j := 0; j < card.Score; j++ {
			cards[i+j+1].Copies += 1 + card.Copies
		}

	}
}
