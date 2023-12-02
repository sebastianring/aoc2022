package dayone2023

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Calibrate(filename string) (int, error) {
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
		// lineLength := len(lineString)

		numbers := findAllNumbers(lineString)

		firstNumber := numbers[0]
		lastNumber := numbers[len(numbers)-1]

		stringNumber := fmt.Sprint(firstNumber) + fmt.Sprint(lastNumber)

		value, err := strconv.Atoi(stringNumber)

		if err != nil {
			log.Println(err)
			panic(err)
		}

		sum += value

		// log.Printf("Numbers: %v - First number: %v - Last: %v - Value: %v - stringNumber: %v \n",
		// numbers, firstNumber, lastNumber, value, stringNumber)

	}

	return sum, nil
}

func findAllNumbers(input string) []int {
	allNumbers := []int{}
	// log.Println(allNumbers)

	for _, char := range input {
		// log.Println(char)
		number, err := strconv.Atoi(string(char))

		if err != nil {
			continue
		}

		allNumbers = append(allNumbers, number)
	}

	return allNumbers
}

func CalibratePartTwo(filename string) (int, error) {
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
		// lineLength := len(lineString)

		numbers := findAllNumbersPartTwo(lineString)

		firstNumber := numbers[0]
		lastNumber := numbers[len(numbers)-1]

		stringNumber := fmt.Sprint(firstNumber) + fmt.Sprint(lastNumber)

		value, err := strconv.Atoi(stringNumber)

		if err != nil {
			log.Println(err)
			panic(err)
		}

		sum += value

		log.Printf("Numbers: %v - First number: %v - Last: %v - Value: %v - stringNumber: %v \n",
			numbers, firstNumber, lastNumber, value, stringNumber)

	}

	return sum, nil
}

func findAllNumbersPartTwo(input string) []int {
	allNumbers := []int{}
	// log.Println(input)

	currString := ""
	convertMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"zero":  0,
	}

	for _, char := range input {
		// log.Println(char)
		number, err := strconv.Atoi(string(char))

		if err != nil {
			currString += string(char)
			stringLength := len(currString)

			for i := stringLength; i >= 0; i-- {
				// log.Println("checking: ", currString[i:stringLength])
				value, ok := convertMap[currString[i:stringLength]]

				if ok {
					// log.Println("Found a text value: ", currString[i:stringLength], " with value: ", value)
					allNumbers = append(allNumbers, value)
					// currString = ""
					break
				}
			}

		} else {
			currString = ""
			allNumbers = append(allNumbers, number)
		}

	}

	return allNumbers
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
