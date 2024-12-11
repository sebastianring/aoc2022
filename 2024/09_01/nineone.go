package nineone

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func nineOne(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	sum := 0
	var data string

	for {
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			return 0, err
		}

		lineString = strings.TrimSuffix(lineString, "\n")
		data = lineString
	}

	stringResult := []string{}
	fileCheck := true
	fileCtr := 0
	for i := 0; i < len(data); i++ {
		repeater, _ := strconv.Atoi(data[i : i+1])
		if fileCheck {
			for ctr := 0; ctr < repeater; ctr++ {
				stringResult = append(stringResult, strconv.Itoa(fileCtr))
			}
			fileCtr++
			fileCheck = false
		} else {
			for ctr := 0; ctr < repeater; ctr++ {
				stringResult = append(stringResult, ".")
			}
			fileCheck = true
		}
	}

	// fmt.Println(stringResult)

	// left = free space
	left := 0
	// right = data
	right := len(stringResult) - 1
	for left < right {
		for stringResult[left] != "." {
			left++
		}

		for stringResult[right] == "." {
			right--
		}

		if left >= right {
			break
		}

		// fmt.Printf("left: %d, right: %d\n", left, right)
		// fmt.Printf("leftVal: %s, rightVal: %s\n", stringResult[left], stringResult[right])
		stringResult[left] = stringResult[right]
		stringResult[right] = "."
		left++
	}

	for i := 0; i < left; i++ {
		if stringResult[i] == "." {
			break
		}

		id, _ := strconv.Atoi(stringResult[i])

		sum += id * i
	}

	// fmt.Println(stringResult)

	return sum, nil
}
