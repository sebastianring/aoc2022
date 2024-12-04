package daythreepartone

import (
	"bufio"
	// "fmt"
	"os"
	"strconv"
	"strings"
)

func DayOne(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
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

			return 0, err
		}

		lineString = strings.TrimSuffix(lineString, "\n")
		i := 0
		for i < len(lineString)-1 {
			newPos, score := findMul(lineString[i:])
			if score == -1 {
				break
			}

			i += newPos
			sum += score
		}
	}

	return sum, nil
}

func findMul(s string) (int, int) {
	i := strings.Index(s, "mul(")
	if i == -1 {
		return 0, -1
	}

	parts := strings.SplitN(s[i+4:], ")", 2)
	// fmt.Printf("trying to add this: %s\n", parts[0])
	mathParts := strings.SplitN(parts[0], ",", 2)
	if len(mathParts) != 2 {
		return i + 4, 0
	}

	numbers := []int{}
	for ctr, math := range mathParts {
		_ = ctr
		// fmt.Printf("part #%d: %s\n", ctr, math)
		number, err := strconv.Atoi(math)
		if err != nil {
			return i + 4, 0
		}

		numbers = append(numbers, number)
	}

	result := numbers[0] * numbers[1]

	return i + 4, result
}
