package seven

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type calc struct {
	sum    int
	values []int
}

func PartOne(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	sum := 0
	calcs := []calc{}

	for {
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			return 0, err
		}

		lineString = strings.TrimSuffix(lineString, "\n")

		newCalc := calc{}
		parts := strings.Split(lineString, ":")
		sumStr := parts[0]

		newCalc.sum, _ = strconv.Atoi(sumStr)

		parts[1] = strings.TrimPrefix(parts[1], " ")
		parts = strings.Split(parts[1], " ")

		for _, value := range parts {
			newVal, _ := strconv.Atoi(value)
			newCalc.values = append(newCalc.values, newVal)
		}

		calcs = append(calcs, newCalc)
	}

	for _, c := range calcs {
		c.findSolution()
	}

	return sum, nil
}

// 10 23 9
// 10 + 23 + 9
// 10 * 23 + 9
// 10 * 23 * 9
// 10 + 23 * 9
//
// 10 23 9 12
// 10 + 23 + 9 + 12
// 10 + 23 + 9 * 12
// 10 + 23 * 9 * 12
// 10 * 23 * 9 * 12
// 10 * 23 * 9 + 12
// 10 * 23 + 9 + 12
// 10 + 23 * 9 + 12

type operation func(...int) int

func (c *calc) findSolution() int {
	operations := []operation{multiply, addition}
	placements := len(c.values) - 1
	_ = operations

	a := "+"
	b := "*"
	_ = a
	_ = b

	for i := 0; i < placements; i++ {
		for j := 0; j < placements; j++ {

		}
	}

	return 0
}

func recursive(input []int) {

}

func multiply(v ...int) int {
	result := 0
	for _, val := range v {
		result *= val
	}

	return result
}

func addition(v ...int) int {
	result := 0
	for _, val := range v {
		result += val
	}

	return result
}
