package dayfive2023_test

import (
	"fmt"
	"testing"

	dayfive2023 "github.com/sebastianring/aoc2022/2023/5_1"
)

func TestSeedToLocation(t *testing.T) {
	result, err := dayfive2023.SeedToLocation("data.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}
