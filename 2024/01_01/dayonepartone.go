package dayonepartone

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	leftList := []int{}
	rightList := []int{}

	for {
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			return 0, err
		}

		lineString = strings.TrimSuffix(lineString, "\n")
		parts := strings.Split(lineString, "   ")

		leftInt, _ := strconv.Atoi(parts[0])
		rightInt, _ := strconv.Atoi(parts[1])

		leftList = append(leftList, leftInt)
		rightList = append(rightList, rightInt)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	for i := 0; i < len(leftList); i++ {
		sum += absolute(leftList[i], rightList[i])
		fmt.Printf("iteration #%d: left: %d right:%d abs: %d\n", i, leftList[i], rightList[i], absolute(leftList[i], rightList[i]))
	}

	return sum, nil
}

func absolute(a int, b int) int {
	var c int
	if a > b {
		c = a - b
	} else {
		c = b - a
	}

	if c < 0 {
		return c * -1
	}

	return c
}
