package dayninetwo2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type NumbersValue struct {
	OriginalValues   []int
	FinalValues      []int
	Length           int
	NextNumbersValue *NumbersValue
}

func DayNine(filename string) (int, error) {
	sum := 0
	numbersValues, err := GetInputFromFile(filename)

	if err != nil {
		return 0, err
	}

	for i, v := range numbersValues {
		err := numbersValues[i].ExtrapolateNextNumbers()

		if err != nil {
			panic(err)
		}

		tempsum, err := numbersValues[i].AddFinalNumbers()

		fmt.Println("Tempsum total: ", tempsum)

		if err != nil {
			panic(err)
		}

		sum += tempsum

		fmt.Printf("#%v - Original values: %v - Final values: %v - Sum: %v\n", i, v.OriginalValues, v.FinalValues, sum)
	}

	return sum, nil
}

func (nv *NumbersValue) AddFinalNumbers() (int, error) {
	// fmt.Printf("Adding final value for: %v \n", nv.OriginalValues)
	var tempsum int
	var err error

	if nv.Length > 1 {
		if nv.NextNumbersValue != nil {
			tempsum, err = nv.NextNumbersValue.AddFinalNumbers()

			if err != nil {
				panic(err)
			}
		}

		var added int

		added = nv.GetFirstOriginalValue() - tempsum

		nv.FinalValues = append(nv.FinalValues, added)
		for _, val := range nv.OriginalValues {
			nv.FinalValues = append(nv.FinalValues, val)
		}

		tempsum = added

		fmt.Printf("Added value: %v to this list: %v - returning: %v\n", added, nv.FinalValues, tempsum)
	}

	return tempsum, nil
}

func (nv *NumbersValue) GetLastOriginalValue() int {
	return nv.OriginalValues[nv.Length-1]
}

func (nv *NumbersValue) GetFirstOriginalValue() int {
	return nv.OriginalValues[0]
}

func (nv *NumbersValue) ExtrapolateNextNumbers() error {
	values := []int{}

	for i := 0; i < len(nv.OriginalValues)-1; i++ {
		j := i + 1
		val := nv.OriginalValues[j] - nv.OriginalValues[i]
		values = append(values, val)
	}

	// fmt.Printf("Extrapolated from: %v to these numbers: %v \n", nv.OriginalValues, values)

	nv.NextNumbersValue = NewNumbersValue(values)

	if nv.NextNumbersValue.Length > 1 && !AllZeroes(values) {
		nv.NextNumbersValue.ExtrapolateNextNumbers()
	}

	return nil
}

func AllZeroes(values []int) bool {
	for _, val := range values {
		if val != 0 {
			return false
		}
	}

	return true
}

func NewNumbersValue(values []int) *NumbersValue {
	nv := NumbersValue{
		OriginalValues: values,
		Length:         len(values),
	}

	return &nv
}

func GetInputFromFile(filename string) ([]*NumbersValue, error) {
	numbersValues := []*NumbersValue{}
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}
		}

		lineString = strings.TrimSuffix(lineString, "\n")
		numbersValue := NewNumbersValueFromString(strings.Split(lineString, " "))
		numbersValues = append(numbersValues, numbersValue)
	}

	return numbersValues, nil
}

func NewNumbersValueFromString(input []string) *NumbersValue {
	values := []int{}

	for _, v := range input {
		stringValue, err := strconv.Atoi(v)

		if err != nil {
			panic(err)
		}

		values = append(values, stringValue)
	}

	nv := NumbersValue{
		OriginalValues: values,
		FinalValues:    []int{},
		Length:         len(values),
	}

	return &nv
}
