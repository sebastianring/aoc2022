package dayone_test

import (
	"fmt"
	"github.com/sebastianring/aoc2022/1_dayone"
	"testing"
)

func TestCalorieCounting(t *testing.T) {
	result, err := dayone.CalorieCounting("data.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}

func TestCalorieCountingTopThree(t *testing.T) {
	result, err := dayone.CalorieCountingTopThree("data.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}
