package dayfive2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RangeValues struct {
	Source int
	Dest   int
	Length int
}

func SeedToLocation(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	sum := 0

	var seeds []string
	convertMap := make(map[string][]RangeValues)
	var currentString string

	for {
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			return 0, err
		}

		lineString = strings.TrimSuffix(lineString, "\n")
		titles := strings.Split(lineString, ":")

		if titles[0] == "" {
			currentString = ""
			continue
		}

		values := strings.Split(lineString, " ")

		if currentString == "" {
			switch titles[0] {
			case "seeds":
				tempseeds := strings.Trim(titles[1], " ")
				seeds = strings.Split(tempseeds, " ")
			default:
				currentString = values[0]
			}
		} else {
			source, err := strconv.Atoi(values[0])

			if err != nil {
				return 0, err
			}

			dest, err := strconv.Atoi(values[1])

			if err != nil {
				return 0, err
			}

			length, err := strconv.Atoi(values[2])

			if err != nil {
				return 0, nil
			}

			rv := RangeValues{
				Source: source,
				Dest:   dest,
				Length: length,
			}

			convertMap[currentString] = append(convertMap[currentString], rv)
		}
	}

	fmt.Printf("Seeds: %v \n", seeds)
	for k, v := range convertMap {
		fmt.Printf("Key: %v, Value: %v \n", k, v)
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
