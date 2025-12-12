package daythree

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func SumRucksackPartTwo(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		log.Println("Error opening file: ", err)
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	priomap := generateAlphabetPriority()
	fmt.Println(priomap)
	sum := 0

	rowctr := 0
	group := []string{}

	for {
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			log.Println("Error reading line: ", err)
			return 0, err
		}

		group = append(group, lineString)
		rowctr++

		if rowctr == 3 {
			triplicate := compareBadgesInGroup(group, 3)
			result := priomap[triplicate]
			sum += result

			fmt.Println("Triplicate: ", triplicate, string(triplicate), "Result: ", result, "Sum: ", sum)

			rowctr = 0
			group = []string{}
		}
	}

	return sum, nil
}

func compareBadgesInGroup(rucksacks []string, target int) byte {
	totalMap := make(map[byte]int)

	for _, s := range rucksacks {
		tempMap := make(map[byte]bool)

		for _, v := range s {
			charByte := byte(v)

			_, ok := tempMap[charByte]

			if ok {
				continue
			} else {
				tempMap[charByte] = true
			}

			_, ok = totalMap[charByte]

			if !ok {
				totalMap[charByte] = 1
			} else {
				totalMap[charByte]++

				if totalMap[charByte] == target {
					return charByte
				}
			}
		}
	}

	return 0
}
