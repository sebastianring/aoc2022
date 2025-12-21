package sixtwo

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type set struct {
	strings []string
	ints    []int

	operation
}

type operation string

const (
	multiply operation = "*"
	addition operation = "+"
	// subtraction operation = "-"
	// division    operation = "/"
)

func Template(lines []string) int {
	sum := 0
	data, ops := FormatData(lines)

	for _, d := range data {
		fmt.Println(d)
	}

	for i := range len(data) {
		sum += calcCol(data[i].ints, ops[i])
	}

	return sum
}

func calcCol(data []int, op operation) int {
	switch op {
	case addition:
		return add(data)
	case multiply:
		return mult(data)
	default:
		log.Fatal("operation not supported")
	}

	return 0
}

func add(data []int) int {
	sum := 0

	for _, v := range data {
		sum += v
	}

	return sum
}

func mult(data []int) int {
	sum := data[0]

	for i := 1; i < len(data); i++ {
		sum *= data[i]
	}

	return sum
}

func FormatData(lines []string) ([]set, []operation) {
	res := []set{}
	cols := len(lines[0])
	rows := len(lines) - 1
	ops := make([]operation, cols)
	cSet := set{}

	for x := range cols {
		hit := false
		colStr := ""
		for y := range rows {
			if string(lines[y][x]) == " " {
				continue
			}

			hit = true
			colStr += string(lines[y][x])
		}

		if hit {
			cSet.strings = append(cSet.strings, colStr)
		} else {
			res = append(res, cSet)
			cSet = set{}
		}
	}
	res = append(res, cSet)

	for i, r := range res {
		for _, s := range r.strings {
			v, _ := strconv.Atoi(s)
			res[i].ints = append(res[i].ints, v)
		}
	}

	opStrings := strings.Fields(lines[rows])
	for x, op := range opStrings {
		ops[x] = operation(op)
	}

	return res, ops
}
