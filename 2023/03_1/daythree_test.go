package daythree2023_test

import (
	"fmt"
	daythree2023 "github.com/sebastianring/aoc2022/2023/03_1"
	"testing"
)

func TestEngineSchematic(t *testing.T) {
	result, err := daythree2023.EngineSchematic("data2.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}

func TestEngineSchematic2(t *testing.T) {
	result, err := daythree2023.EngineSchematic("data3.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}

func TestEngineSchematicPartTwo(t *testing.T) {
	result, err := daythree2023.EngineSchematicPartTwo("data2.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}
