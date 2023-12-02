package daytwo2023

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func CubeGame(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		// log.Println("Error opening file: ", err)
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	sum := 0

	for {
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			// log.Println("Error reading line: ", err)
			return 0, err
		}

		lineString = strings.TrimSuffix(lineString, "\n")

		split := strings.Split(lineString, ":")
		gameId := getGameId(split[0])
		ok := isCubesInInterval(split[1])

		log.Printf("gameId: %v content: %v", gameId, ok)

		if ok {
			sum += gameId
		}

	}

	return sum, nil
}

func CubeGamePartTwo(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		// log.Println("Error opening file: ", err)
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	sum := 0

	for {
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			// log.Println("Error reading line: ", err)
			return 0, err
		}

		lineString = strings.TrimSuffix(lineString, "\n")

		split := strings.Split(lineString, ":")
		gameId := getGameId(split[0])
		partialSum := isCubesInIntervalPartTwo(split[1])

		log.Printf("gameId: %v partialSum: %v", gameId, partialSum)

		sum += partialSum
	}

	return sum, nil
}

func isCubesInIntervalPartTwo(input string) int {
	rules := map[string]int{
		"blue":  14,
		"red":   12,
		"green": 13,
	}

	log.Println(rules)

	rounds := strings.Split(input, ";")

	maxCubes := map[string]int{
		"blue":  0,
		"red":   0,
		"green": 0,
	}

	for _, round := range rounds {
		cubes := strings.Split(round, ",")

		for _, cube := range cubes {
			tempString := strings.Split(cube, " ")

			qty, err := strconv.Atoi(tempString[1])

			if err != nil {
				log.Println("Issue converting: ", err)
				panic(err)
			}

			color := tempString[2]

			if qty > maxCubes[color] {
				maxCubes[color] = qty
			}
		}
	}

	sum := 0

	for _, v := range maxCubes {
		if sum == 0 {
			sum = v
		} else {
			sum *= v
		}
	}

	return sum
}

func isCubesInInterval(input string) bool {
	rules := map[string]int{
		"blue":  14,
		"red":   12,
		"green": 13,
	}

	log.Println(rules)

	rounds := strings.Split(input, ";")

	for _, round := range rounds {
		cubes := strings.Split(round, ",")
		for _, cube := range cubes {
			tempString := strings.Split(cube, " ")

			qty, err := strconv.Atoi(tempString[1])

			if err != nil {
				log.Println("Issue converting: ", err)
				panic(err)
			}

			color := tempString[2]

			if qty > rules[color] {
				return false
			}
		}
	}

	return true
}

func getGameId(input string) int {
	numberString := strings.TrimPrefix(input, "Game ")
	number, err := strconv.Atoi(numberString)

	if err != nil {
		log.Println("Issue converting string to int:", err)
	}

	return number
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
