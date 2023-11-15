package daythree

import (
	"bufio"
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

			rowctr = 0
			group = []string{}
		}

		// fmt.Println(result)
	}

	return sum, nil
}

func compareBadgesInGroup(rucksacks []string, target int) byte {
	totalMap := make(map[byte]int)

	for _, s := range rucksacks {
		for _, v := range s {
			charByte := byte(v)
			_, ok := totalMap[charByte]

			if !ok {
				totalMap[charByte] = 1
			} else {
				if totalMap[charByte] == target-1 {
					return charByte
				}
			}
		}
	}

	return 0
}

//
// mid := len(lineString) / 2
// compartment1 := createRucksack(lineString[:mid])
// compartment2 := createRucksack(lineString[mid:])
//
// // fmt.Println(lineString)
// // fmt.Println(compartment1, compartment2)
// // fmt.Println("---------------")
//
// duplicate := compareCompartments(compartment1, compartment2)
// // fmt.Println("DUPLICATE:", string(duplicate))
//
