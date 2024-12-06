package dayfiveparttwo

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

	for i, update := range updates {
		fmt.Printf("#%d: %v\n", i, update)
		ctr := 0
		sumValue := 0
		for {
			indexOne, indexTwo := checkUpdates(update, pageRules)
			if indexOne > 0 {
				tempVal := update[indexOne]
				update[indexOne] = update[indexTwo]
				update[indexTwo] = tempVal

				ctr++
			} else if indexOne == -1 && indexTwo == -1 {
				break
			}
		}

		if ctr > 0 {
			middle := len(update) / 2
			sumValue, err = strconv.Atoi(update[middle])
			if err != nil {
				fmt.Printf("error converting: %s\n", err.Error())
			}
		}

		sum += sumValue

	}

	return sum, nil
}

func checkUpdates(update []string, rules map[string][]string) (int, int) {
	invalidNumbers := map[string]int{}
	for i := 0; i < len(update); i++ {
		_, exist := invalidNumbers[update[i]]
		if exist {
			return i, invalidNumbers[update[i]]
		}

		_, ruleExist := rules[update[i]]
		if ruleExist {
			for _, invalidNumber := range rules[update[i]] {
				invalidNumbers[invalidNumber] = i
			}
		}
	}

	return -1, -1
}
