package dayoneparttwo

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func DayOnePartTwo(filename string) (int, error) {
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

	multiplier := map[int]int{}

	for _, number := range leftList {
		// _, exists := multiplier[number]
		// if exists {
		// 	continue
		// }

		multiplier[number] = 0
		for _, number2 := range rightList {
			if number == number2 {
				multiplier[number]++
			}
		}
	}

	for _, leftVal := range leftList {
		multi, ok := multiplier[leftVal]
		if !ok {
			log.Fatalf("something went really wrong bro, %d", leftVal)
		}

		sum += leftVal * multi
		fmt.Printf("val: %d, multi %d\n", leftVal, multi)
		fmt.Printf("adding %d on the sum: %d\n", (leftVal * multi), sum)
	}

	return sum, nil
}
