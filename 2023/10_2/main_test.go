package daytentwo2023_test

import (
	"fmt"
	dayten2023 "github.com/sebastianring/aoc2022/2023/10_2"
	"testing"
)

func TestDayTenDataOne(t *testing.T) {
	result, err := dayten2023.DayTen("data.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}

func TestDayTenDataTwo(t *testing.T) {
	result, err := dayten2023.DayTen("data2.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}

func TestDayTenDataThree(t *testing.T) {
	result, err := dayten2023.DayTen("data3.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}

func TestDayTenDataFour(t *testing.T) {
	result, err := dayten2023.DayTen("data4.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}

func TestDayTenDataFive(t *testing.T) {
	result, err := dayten2023.DayTen("data5.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}
