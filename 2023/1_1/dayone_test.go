package dayone2023_test

import (
	"fmt"
	dayone2023 "github.com/sebastianring/aoc2022/2023/1_1"
	"testing"
)

func TestCalibrate(t *testing.T) {
	result, err := dayone2023.Calibrate("data2.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}

func TestCalibratePartTwo(t *testing.T) {
	result, err := dayone2023.CalibratePartTwo("data2.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}
