package dayfour2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//        winning numbers |  your numbers
// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53

func ScratchCards(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		// log.Println("Error opening file: ", err)
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	sum := 0

	// points := make(map[int]int, 0)
	// _ = points

	for {
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			return 0, err
		}

		lineString = strings.TrimSuffix(lineString, "\n")
		titleSplit := strings.Split(lineString, ":")

		for i, v := range titleSplit {
			fmt.Printf("I: %v, V: %v \n", i, v)
		}

		numbersSplit := strings.Split(titleSplit[1], "|")

		for _, v := range numbersSplit {
			v = strings.TrimSuffix(v, " ")
			v = strings.TrimPrefix(v, " ")
		}

		winningNumbers := strings.Split(numbersSplit[0], " ")
		winningNumbers = winningNumbers[1 : len(winningNumbers)-1]
		playerNumbers := strings.Split(numbersSplit[1], " ")
		playerNumbers = playerNumbers[1:]

		fmt.Printf("Winning numbers: %v ---- Player numbers: %v \n", winningNumbers, playerNumbers)

		outcome := ComparePlayerToWinning(playerNumbers, winningNumbers)
		fmt.Printf("Outcome: %v\n", outcome)

		cardSum := 0

		for i := 0; i < len(outcome); i++ {
			if cardSum == 0 {
				cardSum = 1
			} else {
				cardSum = cardSum * 2
			}
		}

		sum += cardSum

		fmt.Printf("Total value for card: %v \n", sum)
		fmt.Println("----------------")
	}

	return sum, nil
}

func ComparePlayerToWinning(playerNumbers []string, winningNumbers []string) []int {
	result := []int{}

	for _, winningNumber := range winningNumbers {
		if winningNumber == " " || winningNumber == "" {
			continue
		}

		for _, playerNumber := range playerNumbers {
			if playerNumber == " " || playerNumber == "" {
				continue
			}

			n1 := strings.TrimLeft(winningNumber, " ")
			n2 := strings.TrimLeft(playerNumber, " ")

			// fmt.Println("winningNumber: ", winningNumber, "playerNumber:", playerNumber)
			if n1 == n2 {
				number, err := strconv.Atoi(n2)

				if err != nil {
					panic(err)
				}

				result = append(result, number)
			}
		}
	}

	return result
}

type CardValue struct {
	Number        int
	Value         int
	Instances     int
	TrailingCards int
}

func ScratchCardsPartTwo(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		// log.Println("Error opening file: ", err)
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	// allCards := make(map[int]CardValue, 0)
	allCards := []CardValue{}
	sum := 0

	// points := make(map[int]int, 0)
	// _ = points

	for {
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			return 0, err
		}

		lineString = strings.TrimSuffix(lineString, "\n")
		titleSplit := strings.Split(lineString, ":")

		for i, v := range titleSplit {
			fmt.Printf("I: %v, V: %v \n", i, v)
		}

		cardString := strings.TrimPrefix(titleSplit[0], "Card ")
		cardString = strings.TrimSpace(cardString)
		fmt.Println(cardString, titleSplit)
		card, err := strconv.Atoi(cardString)

		if err != nil {
			return 0, err
		}

		fmt.Printf("CARD: %v\n", card)

		numbersSplit := strings.Split(titleSplit[1], "|")

		for _, v := range numbersSplit {
			v = strings.TrimLeft(v, " ")
			v = strings.TrimRight(v, " ")
		}

		winningNumbers := strings.Split(numbersSplit[0], " ")
		playerNumbers := strings.Split(numbersSplit[1], " ")

		fmt.Printf("Winning numbers: %v ---- Player numbers: %v \n", winningNumbers, playerNumbers)

		outcome := ComparePlayerToWinning(playerNumbers, winningNumbers)
		fmt.Printf("Outcome: %v\n", outcome)

		cardSum := 0

		for i := 0; i < len(outcome); i++ {
			if cardSum == 0 {
				cardSum = 1
			} else {
				cardSum *= 2
			}
		}

		currentCard := CardValue{
			Number:        card,
			Value:         cardSum,
			Instances:     1,
			TrailingCards: len(outcome),
		}

		allCards = append(allCards, currentCard)

		// if val, ok := allCards[card]; ok {
		// 	val.Instances += 1
		// 	allCards[card] = val
		// } else {
		// 	newCardValue := CardValue{
		// 		// Value:         cardSum,
		// 		Instances:     1,
		// 		TrailingCards: len(outcome),
		// 	}
		//
		// 	allCards[card] = newCardValue
		// }

		// sum += cardSum

		fmt.Printf("Total value for card: %v %v \n", sum, allCards[card-1])
		fmt.Println("----------------")
	}

	for k, v := range allCards {
		fmt.Printf("This card: %v with trailing values: %v and instances: %v ------- Trailing: ", k, v.TrailingCards, v.Instances)
		// fmt.Printf("Card #%v Value: %v Multiplier: %v Trailing: %v \n", k, card.Value, card.Multiplier, card.TrailingCards)
		for i := 0; i < v.TrailingCards; i++ {
			nextCard := k + i + 1

			if nextCard >= len(allCards) {
				continue
			}

			fmt.Printf("%v, ", nextCard)
			allCards[nextCard].Instances += v.Instances
		}

		fmt.Printf("\n")
	}

	for _, v := range allCards {
		sum += v.Instances
	}

	fmt.Println("TOTAL RESULT:", len(allCards))

	for i := 0; i < len(allCards); i++ {
		fmt.Printf("Card #%v Value: %v Instances: %v Trailing: %v \n", i, allCards[i].Value, allCards[i].Instances, allCards[i].TrailingCards)
	}

	return sum, nil
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}
