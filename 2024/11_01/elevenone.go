package elevenone

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sebastianring/aoc2022/utils"
)

func eleven(filename string) (int, error) {
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
	// rule 1: if 0 -> 1
	// rule 2: even number of digits -> split into 2
	// rule 3: else -> multiplied with 2024
	// fmt.Printf("checking #%d, val: %s\n", i+j, data[i+j])

	maxBlinks := 25
	dataresult := []string{}

	for i := 0; i < len(data); i++ {
		oldData := []string{data[i]}
		for blinkCtr := 0; blinkCtr < maxBlinks; blinkCtr++ {
			newData := []string{}
			for j := 0; j < len(oldData); j++ {
				val := oldData[j]
				operation := checkValue(val)
				result := operation(val)
				newData = append(newData, result...)
			}

			oldData = newData
		}

		dataresult = append(dataresult, oldData...)
	}

	sum = len(dataresult)

	return sum, nil
}

type manipulate func(string) []string

func checkValue(s string) manipulate {
	// fmt.Println("triggered check val")
	if s == "0" {
		return one
	} else if utils.IsEven(len(s)) {
		return split
	}

	return multiply2024
}

func split(s string) []string {
	// fmt.Println("triggered split")
	split1 := checkZero(s[:len(s)/2])
	split2 := checkZero(s[len(s)/2:])

	return []string{split1, split2}
}

func checkZero(s string) string {
	sInt, _ := strconv.Atoi(s)

	return strconv.Itoa(sInt)
}

func multiply2024(s string) []string {
	// fmt.Println("triggered multiply")
	val, _ := strconv.Atoi(s)

	sum := val * 2024

	return []string{strconv.Itoa(sum)}
}

func one(s string) []string {
	// fmt.Println("triggered one")

	return []string{"1"}
}

// for ctr, r := range result {
// 	if ctr == 0 {
// 		data = utils.ReplaceAtIndex(data, i+j, r)
// 	} else {
// 		data = utils.AddAtIndex(data, i+j+ctr, r)
// 		j++
// 	}
// }
