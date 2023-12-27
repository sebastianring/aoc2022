package daytwo2023_test

import (
	"fmt"
	dayone2023 "github.com/sebastianring/aoc2022/2023/02_1"
	"testing"
)

func TestCubeGame(t *testing.T) {
	result, err := dayone2023.CubeGame("data2.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}

func TestCubeGamePartTwo(t *testing.T) {
	result, err := dayone2023.CubeGamePartTwo("data2.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}
