package daythree

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func SumPrioritiesInRucksack(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		log.Println("Error opening file: ", err)
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	priomap := generateAlphabetPriority()
	// log.Println(priomap)
	//

	sum := 0

	for {
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			log.Println("Error reading line: ", err)
			return 0, err
		}

		mid := len(lineString) / 2
		compartment1 := createRucksack(lineString[:mid])
		compartment2 := createRucksack(lineString[mid:])

		// fmt.Println(lineString)
		// fmt.Println(compartment1, compartment2)
		// fmt.Println("---------------")

		duplicate := compareCompartments(compartment1, compartment2)
		// fmt.Println("DUPLICATE:", string(duplicate))
		result := priomap[duplicate]
		// fmt.Println(result)

		sum += result
	}

	return sum, nil
}

func compareCompartments(compartments ...map[byte]int) byte {
	totalMap := make(map[byte]int)

	for _, m := range compartments {
		for k := range m {
			_, ok := totalMap[k]

			if !ok {
				totalMap[k] = 1
			} else {
				return k
			}
		}
	}

	return 0
}

func compareCompartmentsTotal(compartments ...map[byte]int) map[byte]int {
	totalMap := make(map[byte]int)

	for _, m := range compartments {
		for k, v := range m {
			_, ok := totalMap[k]

			if !ok {
				totalMap[k] = v
			} else {
				totalMap[k] += v
			}
		}
	}

	return totalMap
}

func createRucksack(items string) map[byte]int {
	rucksack := make(map[byte]int)
	// log.Println("Items to be created: ", items)

	for _, char := range items {
		bytechar := byte(char)
		_, ok := rucksack[bytechar]

		if !ok {
			rucksack[bytechar] = 1
		} else {
			rucksack[bytechar] += 1
		}

		// log.Println("Val: ", val)
	}

	// log.Println("Created rucksack")

	return rucksack
}

func generateAlphabetPriority() map[byte]int {
	prioMap := make(map[byte]int)

	prio := 1

	for i := byte('a'); i <= byte('z'); i++ {
		fmt.Printf(string(i))
		prioMap[i] = prio
		prio++
	}

	for i := byte('A'); i <= byte('Z'); i++ {
		fmt.Printf(string(i))
		prioMap[i] = prio
		prio++
	}

	fmt.Printf("\n")

	return prioMap
}
