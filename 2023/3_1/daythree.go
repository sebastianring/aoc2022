package daythree2023

import (
	"bufio"
	"strconv"
	// "fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

type Pos struct {
	x int
	y int
}

type NumberSchematic struct {
	Pos
	Value  int
	Length int
}

type TempNumbersSchematic struct {
	Pos
	Value string
}

func EngineSchematic(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		// log.Println("Error opening file: ", err)
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	sum := 0
	symbols := []Pos{}
	numbers := []NumberSchematic{}
	y := 0

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

		currString := TempNumbersSchematic{
			Value: "",
		}

		lineLength := len(lineString) - 1

		for x, c := range lineString {
			char := string(c)

			if ifSymbol(char) {
				symbols = append(symbols, Pos{
					x: x,
					y: y,
				})

				err := AddCurrentStringAsNumber(&numbers, currString, x, y)

				if err != nil {
					return 0, err
				}

				currString.Value = ""

			} else if char != "." {
				currString.Value += char
			}

			if currString.Value != "" {
				if char == "." || x == lineLength {
					err := AddCurrentStringAsNumber(&numbers, currString, x, y)

					if err != nil {
						return 0, err
					}

					currString.Value = ""
				}
			}
		}

		if currString.Value != "" {
			err := AddCurrentStringAsNumber(&numbers, currString, lineLength, y)

			if err != nil {
				return 0, nil
			}

			currString.Value = ""
		}

		y++
		// log.Println(currString)
	}

	log.Printf("RESULT: Numbers: %v, Symbols: %v \n", numbers, symbols)

	return sum, nil
}

func AddCurrentStringAsNumber(numbers *[]NumberSchematic, currTempNumber TempNumbersSchematic, x int, y int) error {
	log.Println("X, Y:", x, y)
	value, err := strconv.Atoi(currTempNumber.Value)

	if err != nil {
		return err
	}

	numberLength := len(currTempNumber.Value)

	number := NumberSchematic{
		Pos: Pos{
			x: x - numberLength,
			y: y,
		},
		Length: numberLength,
		Value:  value,
	}

	*numbers = append(*numbers, number)

	return nil
}

func ifSymbol(c string) bool {
	if c == "." {
		return false
	}

	if unicode.IsNumber(rune(c[0])) {
		return false
	}

	return true
}
