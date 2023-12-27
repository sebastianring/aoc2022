package dayfive2023_test

import (
	"fmt"
	"testing"

	dayfive2023 "github.com/sebastianring/aoc2022/2023/05_1"
)

func TestSeedToLocationDataOne(t *testing.T) {
	result, err := dayfive2023.SeedToLocation("data.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}

func TestSeedToLocationDataTwo(t *testing.T) {
	result, err := dayfive2023.SeedToLocation("data2.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}

func TestSeedToLocationPartTwoDataOne(t *testing.T) {
	result, err := dayfive2023.SeedToLocationPartTwo("data.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}

func TestSeedToLocationPartTwoDataTwo(t *testing.T) {
	result, err := dayfive2023.SeedToLocationPartTwo("data2.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}
