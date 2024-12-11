package seven

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct {
	sum    int
	values []int
}

func PartOne(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	sum := 0
	calcs := []calc{}

	for {
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			return 0, err
		}

		lineString = strings.TrimSuffix(lineString, "\n")

		newCalc := calc{}
		parts := strings.Split(lineString, ":")
		sumStr := parts[0]

		newCalc.sum, _ = strconv.Atoi(sumStr)

		parts[1] = strings.TrimPrefix(parts[1], " ")
		parts = strings.Split(parts[1], " ")

		for _, value := range parts {
			newVal, _ := strconv.Atoi(value)
			newCalc.values = append(newCalc.values, newVal)
		}

		calcs = append(calcs, newCalc)
	}

	for _, c := range calcs {
		fmt.Println("------")
		fmt.Println(c)
		sum += c.findSolution()
	}

	return sum, nil
}

func (c *calc) findSolution() int {
	operations := []string{"*", "+"}
	placements := len(c.values) - 1

	permutations := [][]string{}

	var generate func(current []string)
	generate = func(current []string) {
		if len(current) == placements {
			permutations = append(permutations, append([]string{}, current...))
			return
		}

		for _, o := range operations {
			current = append(current, o)
			generate(current)
			current = current[:len(current)-1]
		}
	}

	generate([]string{})

	for _, p := range permutations {
		fmt.Printf("perm: %s\n", p)
		result := calcResult(c.values, p)
		if result == c.sum {
			fmt.Printf(" ----- FOUND CORRECT!!! ------\n")
			fmt.Printf("result from calcResult: %d\n", result)
			return result
		} else {
			fmt.Printf(" ----- no corr!!! ------\n")
		}
	}

	return 0
}

func calcResult(values []int, p []string) int {
	fmt.Printf("calc result input, values: %v, p: %v\n", values, p)
	// allNumbers := []int{}

	prevVal := values[0]
	curValue := 0
	for i := 1; i < len(values); i++ {
		operator := p[i-1]
		if operator == "*" {
			curValue = prevVal * values[i]
		} else {
			// allNumbers = append(allNumbers, prevVal)
			curValue = prevVal + values[i]
		}

		prevVal = curValue
	}

	return curValue

	// allNumbers = append(allNumbers, prevVal)

	// return addition(allNumbers...)
}

// func multiply(v ...int) int {
// 	result := v[0]
//
// 	for i := 1; i < len(v); i++ {
// 		result *= v[i]
// 	}
//
// 	fmt.Printf("multiplied: %v result: %d\n", v, result)
//
// 	return result
// }
//
// func addition(v ...int) int {
// 	result := 0
// 	for _, val := range v {
// 		result += val
// 	}
//
// 	fmt.Printf("added: %v result: %d\n", v, result)
// 	return result
// }
//

func PartTwo(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	sum := 0
	calcs := []calc{}

	for {
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			return 0, err
		}

		lineString = strings.TrimSuffix(lineString, "\n")

		newCalc := calc{}
		parts := strings.Split(lineString, ":")
		sumStr := parts[0]

		newCalc.sum, _ = strconv.Atoi(sumStr)

		parts[1] = strings.TrimPrefix(parts[1], " ")
		parts = strings.Split(parts[1], " ")

		for _, value := range parts {
			newVal, _ := strconv.Atoi(value)
			newCalc.values = append(newCalc.values, newVal)
		}

		calcs = append(calcs, newCalc)
	}

	for _, c := range calcs {
		// fmt.Println("------")
		// fmt.Println(c)
		sum += c.findSolutionPartTwo()
	}

	return sum, nil
}

func (c *calc) findSolutionPartTwo() int {
	operations := []string{"*", "+", "||"}
	placements := len(c.values) - 1

	permutations := [][]string{}

	var generate func(current []string)
	generate = func(current []string) {
		if len(current) == placements {
			permutations = append(permutations, append([]string{}, current...))
			return
		}

		for _, o := range operations {
			current = append(current, o)
			generate(current)
			current = current[:len(current)-1]
		}
	}

	generate([]string{})

	for _, p := range permutations {
		// fmt.Printf("perm: %s\n", p)
		result := calcResultPartTwo(c.values, p)
		if result == c.sum {
			// fmt.Printf(" ----- FOUND CORRECT!!! ------\n")
			// fmt.Printf("result from calcResult: %d\n", result)
			return result
		} else {
			// fmt.Printf(" ----- no corr!!! ------\n")
		}
	}

	return 0
}

func calcResultPartTwo(values []int, p []string) int {
	// fmt.Printf("calc result input, values: %v, p: %v\n", values, p)
	// allNumbers := []int{}

	prevVal := values[0]
	curValue := 0
	for i := 1; i < len(values); i++ {
		operator := p[i-1]
		if operator == "*" {
			curValue = prevVal * values[i]
		} else if operator == "||" {
			curValue = combine(prevVal, values[i])
		} else {
			// allNumbers = append(allNumbers, prevVal)
			curValue = prevVal + values[i]
		}

		prevVal = curValue
	}

	return curValue

	// allNumbers = append(allNumbers, prevVal)

	// return addition(allNumbers...)
}

func combine(a, b int) int {
	c := strconv.Itoa(a) + strconv.Itoa(b)
	result, _ := strconv.Atoi(c)

	return result
}
