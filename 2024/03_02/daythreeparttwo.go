package daythreeparttwo

import (
	"bufio"
	"fmt"
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
	totalData := ""

	for {
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			return 0, err
		}

		lineString = strings.TrimSuffix(lineString, "\n")
		totalData += lineString
	}

	do := 0
	for do != -1 {
		dont := findStr(totalData[do:], "don't()", do)
		mul := findStr(totalData[do:], "mul(", do)
		fmt.Printf("next do: %d, next dont: %d, next mul: %d\n", do, dont, mul)

		for mul < dont || mul != -1 && dont == -1 {
			sum += getMulValues(totalData[mul+4:])
			do = mul + 4

			mul = findStr(totalData[mul+4:], "mul(", do)
		}

		// if nextMul != -1 && nextMul < nextDont || nextDont == -1 && nextMul > 0 {
		// 	score := getMulValues(lineString[nextMul+4:])
		// 	sum += score
		// }

		if do > len(totalData)-1 || mul == -1 {
			break
		}

		do = findStr(totalData[do:], "do()", do)
	}

	return sum, nil
}

func findStr(s, substring string, ctr int) int {
	res := strings.Index(s, substring)
	if res == -1 {
		return -1
	}

	return res + ctr
}

func getMulValues(s string) int {
	parts := strings.SplitN(s, ")", 2)
	fmt.Printf("trying to add this: %s\n", parts[0])
	mathParts := strings.SplitN(parts[0], ",", 2)
	if len(mathParts) != 2 {
		return 0
	}

	numbers := []int{}
	for ctr, math := range mathParts {
		_ = ctr
		fmt.Printf("part #%d: %s\n", ctr, math)
		number, err := strconv.Atoi(math)
		if err != nil {
			return 0
		}

		numbers = append(numbers, number)
	}

	result := numbers[0] * numbers[1]
	fmt.Printf("result: %d\n", result)

	return result
}
