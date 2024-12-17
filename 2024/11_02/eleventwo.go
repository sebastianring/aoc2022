package eleventwo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sebastianring/aoc2022/utils"
)

func eleventwo(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	sum := 0

	data := []string{}

	for {
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			return 0, err
		}

		lineString = strings.TrimSuffix(lineString, "\n")
		data = strings.Split(lineString, " ")
	}

	fmt.Println(data)

	stoneMap := dataToMap(data)
	blinks := 75

	for i := 0; i < blinks; i++ {
		stoneMap = blink(stoneMap)
	}

	for _, ctr := range stoneMap {
		sum += ctr
	}

	return sum, nil
}

func blink(data map[string]int) map[string]int {
	newMap := map[string]int{}

	for val, ctr := range data {
		operation := checkValue(val)
		result := operation(val)

		for _, r := range result {
			newMap[r] += ctr
		}
	}

	return newMap
}

func dataToMap(data []string) map[string]int {
	result := map[string]int{}
	for _, d := range data {
		result[d] += 1
	}

	return result
}

// val := oldData[j]
// 				operation := checkValue(val)
// 				result := operation(val)
// 				memo[oldData[j]] = result
//
// 				newData = append(newData, result...)

type manipulate func(string) []string

func checkValue(s string) manipulate {
	if s == "0" {
		return one
	} else if utils.IsEven(len(s)) {
		return split
	}

	return multiply2024
}

func split(s string) []string {
	split1 := checkZero(s[:len(s)/2])
	split2 := checkZero(s[len(s)/2:])

	return []string{split1, split2}
}

func checkZero(s string) string {
	sInt, _ := strconv.Atoi(s)

	return strconv.Itoa(sInt)
}

func multiply2024(s string) []string {
	val, _ := strconv.Atoi(s)

	sum := val * 2024

	return []string{strconv.Itoa(sum)}
}

func one(s string) []string {
	return []string{"1"}
}
