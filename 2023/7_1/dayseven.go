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
	Cards           []int
	ReplacementCard int
	Bid             int
	Score           []int
	TotalScore      int
	Rank            int
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
		fmt.Printf("Cards: %v, Score: %v Sum: %v Bid: %v Rank: %v \n", hand.Cards, hand.TotalScore, hand.Bid*(i+1), hand.Bid, i+1)
		sum += hand.Bid * (i + 1)
	}

	return sum, nil
}

func (h *Hand) SetScore() {
	scoreMap := map[int]int{}
	highestCount := 0
	highestCard := 0
	fmt.Printf("Started calculated this hand: %v\n", h.Cards)

	for _, card := range h.Cards {
		_, ok := scoreMap[card]

		if !ok {
			scoreMap[card] = 1
		} else {
			scoreMap[card]++
		}

		if scoreMap[card] > highestCount || highestCount == 1 && card > highestCard {
			if card != 1 {
				highestCount = scoreMap[card]
				highestCard = card
				fmt.Printf("Highest card updated: %v, number of cards: %v\n", highestCard, highestCount)
			}
		}
	}

	fmt.Printf("Scoremap: %v\n", scoreMap)

	replacementCard := 0
	// _ = replacementCard

	// if highestCount > 1 {
	replacementCard = highestCard
	// } else {
	// replacementCard = 14
	// }

	for i := 0; i < len(h.Cards); i++ {
		if h.Cards[i] == 1 {
			fmt.Printf("Replaced card: %v ", h.Cards[i])
			// h.Cards[i] = replacementCard
			fmt.Printf(" with card: %v \n", replacementCard)

			scoreMap[1]--
			scoreMap[replacementCard]++

			if scoreMap[1] == 0 {
				delete(scoreMap, 1)
			}
		}
	}

	fmt.Printf("Scoremap: %v\n", scoreMap)
	fmt.Printf("Updated hand looks like this: %v \n", h.Cards)

	sum := []int{}
	totalSum := 0

	for i, card := range scoreMap {
		sum = append(sum, card)

		if i == 0 {
			totalSum = card * card
		} else {
			totalSum += card * card
		}
	}

	slices.Sort(sum)

	h.Score = sum
	h.TotalScore = totalSum

	fmt.Printf("Sum: %v, Totalscore: %v\n", h.Score, h.TotalScore)
	fmt.Printf("-------------- \n")
}

func NewHandFromStrings(cards string, bid string) (Hand, error) {
	cardInts := []int{}
	resultHand := Hand{}

	cardValue := map[string]int{
		"J": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
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
