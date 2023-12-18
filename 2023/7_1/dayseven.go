package dayseven2023

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	Cards      []int
	Bid        int
	Score      []int
	TotalScore int
	Rank       int
}

func CamelCards(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	sum := 0

	hands := []Hand{}

	for {
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			return 0, err
		}

		lineString = strings.TrimSuffix(lineString, "\n")
		texts := strings.Split(lineString, " ")

		newHand, err := NewHandFromStrings(texts[0], texts[1])

		if err != nil {
			return 0, err
		}

		hands = append(hands, newHand)
	}

	fmt.Println(hands)

	for i := 0; i < len(hands); i++ {
		hands[i].SetScore()
	}

	slices.SortFunc(hands, func(a Hand, b Hand) int {
		// if a < b  = negative
		// if a > b  = positive
		// if a == b = 0
		if a.TotalScore == b.TotalScore {
			for i := 0; i < len(a.Cards); i++ {
				if a.Cards[i] == b.Cards[i] {
					continue
				}

				if a.Cards[i] > b.Cards[i] {
					return 1
				}

				return -1
			}
		} else if a.TotalScore > b.TotalScore {
			return 1
		}

		return -1
	})

	for i, hand := range hands {
		fmt.Printf("Cards: %v, Score: %v \n", hand.Cards, hand.TotalScore)
		sum += hand.Bid * (i + 1)
	}

	return sum, nil
}

func (h *Hand) SetScore() {
	scoreMap := map[int]int{}

	for _, card := range h.Cards {
		_, ok := scoreMap[card]

		if !ok {
			scoreMap[card] = 1
		} else {
			scoreMap[card]++
		}
	}

	sum := []int{}
	totalSum := 0

	for _, card := range scoreMap {
		sum = append(sum, card)

		if totalSum == 0 {
			totalSum = card * card
		} else {
			totalSum += card * card
		}
	}

	slices.Sort(sum)

	h.Score = sum
	h.TotalScore = totalSum

	fmt.Printf("Sum: %v, Totalscore: %v\n", h.Score, h.TotalScore)
}

func NewHandFromStrings(cards string, bid string) (Hand, error) {
	cardInts := []int{}
	resultHand := Hand{}

	cardValue := map[string]int{
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"J": 11,
		"Q": 12,
		"K": 13,
		"A": 14,
	}

	for _, char := range cards {
		// fmt.Printf("Trying to find: %v, with with its type: %T\n", string(char), char)
		val, ok := cardValue[string(char)]

		if !ok {
			return resultHand, errors.New("Error finding value in map.")
		}

		cardInts = append(cardInts, val)
	}

	bidInt, err := strconv.Atoi(bid)

	if err != nil {
		return resultHand, err
	}

	resultHand.Cards = cardInts
	resultHand.Bid = bidInt

	return resultHand, nil
}
