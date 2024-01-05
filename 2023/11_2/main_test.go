package dayeleventwo2023_test

import (
	"fmt"
	dayeleven2023 "github.com/sebastianring/aoc2022/2023/11_2"
	"testing"
)

func TestDayElevenDataOne(t *testing.T) {
	result, err := dayeleven2023.DayEleven("data.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}

func TestDayElevenDataTwo(t *testing.T) {
	result, err := dayeleven2023.DayEleven("data2.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}

func TestDayElevenDataThree(t *testing.T) {
	result, err := dayeleven2023.DayEleven("data3.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}
