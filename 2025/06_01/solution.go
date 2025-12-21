package sixone

import (
	"log"
	"strconv"
	"strings"
)

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

	for i := range len(data) {
		sum += calcCol(data[i], ops[i])
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

func FormatData(lines []string) ([][]int, []operation) {
	cols := len(strings.Fields(lines[0]))
	rows := len(lines) - 1

	res := make([][]int, cols)
	ops := make([]operation, cols)

	for i := range len(res) {
		res[i] = make([]int, rows)
	}

	for y := range rows {
		columns := strings.Fields(lines[y])
		for x := range cols {
			val, _ := strconv.Atoi(columns[x])
			res[x][y] = val
		}
	}

	opStrings := strings.Fields(lines[rows])
	for x, op := range opStrings {
		ops[x] = operation(op)
	}

	return res, ops
}
