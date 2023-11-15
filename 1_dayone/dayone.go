package dayone

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
	// "strings"
)

func CalorieCountingTopThree(filename string) (int, error) {
	highscores := []int{0, 0, 0}

	file, err := os.Open(filename)

	if err != nil {
		log.Println("Error opening file: ", err)
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	currSum := 0

	for {
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			log.Println("Error reading line: ", err)
			return 0, err
		}

		lineString = strings.TrimSuffix(lineString, "\n")
		lineLength := len(lineString)

		if lineLength == 0 {
			highscores = CompareAndSetHighscores(highscores, currSum)
			currSum = 0
		} else {
			line, err := strconv.Atoi(lineString)

			if err != nil {
				log.Println("Error converting string to int")
				return 0, err
			}

			currSum += line
		}
	}

	sum := 0

	for _, val := range highscores {
		sum += val
	}

	return sum, nil
}

func CompareAndSetHighscores(highscores []int, score int) []int {
	// score = 250
	// {100, 200, 300}
	// {100, 200, 250, 300}
	//             ^ Added

	highscores = append(highscores, score)
	slices.Sort(highscores)

	return highscores[1:]
}

func CalorieCounting(filename string) (int, error) {
	highscore := 0

	file, err := os.Open(filename)

	if err != nil {
		log.Println("Error opening file: ", err)
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	currSum := 0

	for {
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			log.Println("Error reading line: ", err)
			return 0, err
		}

		lineString = strings.TrimSuffix(lineString, "\n")
		lineLength := len(lineString)

		if lineLength == 0 {
			highscore = max(highscore, currSum)
			currSum = 0
		} else {
			line, err := strconv.Atoi(lineString)

			if err != nil {
				log.Println("Error converting string to int")
				return 0, err
			}

			currSum += line
		}
	}

	return highscore, nil
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
