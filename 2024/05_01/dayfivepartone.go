package dayfivepartone

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DayFivePartOne(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	sum := 0
	pageRules := map[string][]string{}
	pageRuleSection := true
	updates := [][]string{}

	for {
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			return 0, err
		}

		lineString = strings.TrimSuffix(lineString, "\n")
		// section #1 - split on |, until it cant find any |
		// section #2 - split on , until eof
		fmt.Printf("current lineString: %s\n", lineString)
		if pageRuleSection {
			if len(lineString) == 0 {
				pageRuleSection = false
			} else {
				numbers := strings.Split(lineString, "|")
				pageRules[numbers[1]] = append(pageRules[numbers[1]], numbers[0])
			}
		} else {
			tempSlice := []string{}
			numbers := strings.Split(lineString, ",")
			tempSlice = append(tempSlice, numbers...)
			updates = append(updates, tempSlice)
		}
	}

	for k, v := range pageRules {
		fmt.Printf("k: %s, v: %s\n", k, v)
	}

	for i, update := range updates {
		fmt.Printf("#%d: %v\n", i, update)
		valid := checkUpdates(update, pageRules)
		if valid {
			middle := len(update) / 2
			sumValue, err := strconv.Atoi(update[middle])
			if err != nil {
				fmt.Printf("error converting: %s\n", err.Error())
			}

			fmt.Printf("found a valid update, middle index: %d, sumValue: %d\n", middle, sumValue)
			sum += sumValue
		}
	}

	return sum, nil
}

func checkUpdates(update []string, rules map[string][]string) bool {
	invalidNumbers := map[string]int{}
	for i := 0; i < len(update); i++ {
		_, exist := invalidNumbers[update[i]]
		if exist {
			return false
		}

		_, ruleExist := rules[update[i]]
		if ruleExist {
			for _, invalidNumber := range rules[update[i]] {
				invalidNumbers[invalidNumber] += 1
			}
		}
	}

	return true
}
