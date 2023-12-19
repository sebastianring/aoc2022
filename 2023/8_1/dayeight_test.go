package dayeight2023_test

import (
	"fmt"
	"testing"

	dayeight2023 "github.com/sebastianring/aoc2022/2023/8_1"
)

func TestHauntedWastelandDataOne(t *testing.T) {
	result, err := dayeight2023.HauntedWasteland("data.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}

func TestHauntedWastelandDataTwo(t *testing.T) {
	result, err := dayeight2023.HauntedWasteland("data2.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}

func TestHauntedWastelandDataThree(t *testing.T) {
	result, err := dayeight2023.HauntedWasteland("data3.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}

func TestHauntedWastelandPartTwoDataOne(t *testing.T) {
	result, err := dayeight2023.HauntedWastelandPartTwo("data4.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}
