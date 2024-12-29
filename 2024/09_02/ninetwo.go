package ninetwo

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type fileid struct {
	id     int
	index  int
	length int
}

func (f *fileid) getValues() []string {
	var idStr string
	if f.id == -1 {
		idStr = "."
	} else {
		idStr = strconv.Itoa(f.id)
	}
	result := []string{}
	for i := 0; i < f.length; i++ {
		result = append(result, idStr)
	}

	return result
}

func (f *fileid) getDots() []string {
	result := []string{}
	for i := 0; i < f.length; i++ {
		result = append(result, ".")
	}

	return result
}

func (f *fileid) getEndIndex() int {
	return f.index + f.length
}

func nineTwo(filename string) (int, error) {
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

	files := []fileid{}
	empties := []fileid{}

	stringResult := []string{}
	fileCheck := true
	fileCtr := 0
	for i := 0; i < len(data); i++ {
		repeater, _ := strconv.Atoi(data[i : i+1])
		if fileCheck {
			files = append(files, fileid{
				id:     fileCtr,
				index:  len(stringResult),
				length: repeater,
			})

			for ctr := 0; ctr < repeater; ctr++ {
				stringResult = append(stringResult, strconv.Itoa(fileCtr))
			}
			fileCtr++

			fileCheck = false
		} else {
			empties = append(empties, fileid{
				id:     -1,
				index:  len(stringResult),
				length: repeater,
			})
			for ctr := 0; ctr < repeater; ctr++ {
				stringResult = append(stringResult, ".")
			}
			fileCheck = true
		}
	}

	// fmt.Println(stringResult)
	// fmt.Println(files)
	// fmt.Println(empties)

	for right := len(files) - 1; right > -1; right-- {
		for left := 0; left < len(empties); left++ {
			// fmt.Println(empties)
			if empties[left].length >= files[right].length {
				// fmt.Printf("left index: %d, left end index: %d\n", empties[left].index, empties[left].index+files[right].length)
				// fmt.Printf("REPLACING: %v\n", stringResult[empties[left].index:empties[left].index+files[right].length])
				stringResult = slices.Replace(stringResult, empties[left].index, empties[left].index+files[right].length, files[right].getValues()...)

				// fmt.Printf("right index: %d, right end index: %d\n", files[right].index, files[right].getEndIndex())
				stringResult = slices.Replace(stringResult, files[right].index, files[right].getEndIndex(), files[right].getDots()...)

				if empties[left].length == files[right].length {
					empties = removeAtIndex(empties, left)
				} else {
					empties[left].index += files[right].length
					empties[left].length -= files[right].length
				}

				break
			}
		}
	}

	fmt.Println(stringResult)
	// fmt.Println(files)
	// fmt.Println(empties)

	for i := 0; i < len(stringResult); i++ {
		if stringResult[i] == "." {
			continue
		}

		id, _ := strconv.Atoi(stringResult[i])

		sum += id * i
	}

	return sum, nil
}

func removeAtIndex[T any](s []T, index int) []T {
	if index < 0 || index >= len(s) {
		return s
	}
	result := []T{}
	result = append(result, s[:index]...)
	result = append(result, s[index+1:]...)

	return result
}
