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

func EngineSchematic(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		// log.Println("Error opening file: ", err)
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	sum := 0
	y := 0
	symbols := []Pos{}
	numbers := []NumberSchematic{}
	lineLength := 0
	lineHeight := 0

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
		currString := ""
		lineLength = len(lineString) - 1

		for x, c := range lineString {
			char := string(c)

			if ifSymbol(c) {
				symbols = append(symbols, Pos{
					x: x,
					y: y,
				})

				err := AddCurrentStringAsNumber(&numbers, &currString, x, y)

				if err != nil {
					return 0, err
				}

			} else if char != "." {
				currString += char
			}

			if currString != "" {
				if char == "." || x == lineLength {
					err := AddCurrentStringAsNumber(&numbers, &currString, x, y)

					if err != nil {
						return 0, err
					}
				}
			}
		}

		if currString != "" {
			err := AddCurrentStringAsNumber(&numbers, &currString, lineLength, y)

			if err != nil {
				return 0, nil
			}
		}

		y++
	}

	lineHeight = y
	log.Printf("RESULT: Numbers: %v, Symbols: %v \n", numbers, symbols)

	for _, number := range numbers {
		ok := isSymbolNearby(number, &symbols, lineLength, lineHeight)

		if ok {
			sum += number.Value
		}
	}

	return sum, nil
}

func isSymbolNearby(number NumberSchematic, symbols *[]Pos, lineLength int, lineHeight int) bool {
	minx := max(number.x-1, 0)
	maxx := min(number.x+number.Length+1, lineLength)

	miny := max(number.y-1, 0)
	maxy := min(number.y+1, lineHeight)

	for _, symbol := range *symbols {
		// log.Printf("Checking this number: %v at pos %v if symbol is nearby: %v \n", number.Value, number.Pos, symbol)
		// log.Printf("minx: %v, maxx: %v, miny: %v, maxy: %v \n", minx, maxx, miny, maxy)
		if symbol.x <= maxx && symbol.x >= minx &&
			symbol.y >= miny && symbol.y <= maxy {
			log.Printf("VAL: %v POS: %v NEARBY SYMBOL %v", number.Value, number.Pos, symbol)
			// log.Printf("Checking this number: %v at pos %v if symbol is nearby: %v \n", number.Value, number.Pos, symbol)
			// log.Printf("minx: %v, maxx: %v, miny: %v, maxy: %v \n", minx, maxx, miny, maxy)
			return true
		}
	}

	return false
}

func AddCurrentStringAsNumber(numbers *[]NumberSchematic, currTempNumber *string, x int, y int) error {
	if *currTempNumber == "" {
		return nil
	}

	log.Println("X, Y:", x, y)

	value, err := strconv.Atoi(*currTempNumber)

	if err != nil {
		log.Println(err)
		return err
	}

	numberLength := len(*currTempNumber)

	number := NumberSchematic{
		Pos: Pos{
			x: x - numberLength,
			y: y,
		},
		Length: numberLength,
		Value:  value,
	}

	*numbers = append(*numbers, number)
	*currTempNumber = ""

	return nil
}

func ifSymbol(c rune) bool {
	if c == rune('.') ||
		unicode.IsNumber(c) {
		return false
	}

	return true
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
