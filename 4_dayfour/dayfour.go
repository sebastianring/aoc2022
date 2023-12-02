package dayfour

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Interval struct {
	min int
	max int
}

func OverlapSections(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		log.Println("Error opening file: ", err)
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	overlaps := 0

	for {
		lineString, err := reader.ReadString('\n')

		lineString = strings.Replace(lineString, "-", "", 0)
		sections := strings.Split(lineString, ",")

		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			log.Println("Error reading line: ", err)
			return 0, err
		}

		for i, v := range sections {
			fmt.Printf("%v. %s \n", i, v)
		}

		oneMin, err := strconv.Atoi(string(sections[0][0]))

		if err != nil {
			return 0, err
		}

		section1 := Interval{
			min: oneMin,
			max: int(sections[0][2]),
		}

		section2 := Interval{
			min: int(sections[1][0]),
			max: int(sections[1][2]),
		}

		if checkOverlap(section1, section2) {
			overlaps++
		}
	}
	return overlaps, nil
}

func checkOverlap(section1 Interval, section2 Interval) bool {
	fmt.Printf("%v >= %v \n", section1.min, section2.min)
	fmt.Printf("%v <= %v \n", section1.max, section2.max)

	fmt.Printf("%v >= %v \n", section2.min, section1.min)
	fmt.Printf("%v <= %v \n", section2.max, section1.max)

	if section1.min >= section2.min && section1.max <= section2.max ||
		section2.min >= section1.min && section2.max <= section1.max {
		fmt.Println("FOUND AN OVERLAP")
		fmt.Println(section1, section2)
		return true
	}

	return false
}

func checkOverlap2(section1 Interval, section2 Interval) bool {
	// 3 8
	// 2 6
	//
	// 3 8
	// 1 5
	//
	// 3 8
	// 2 9

	var forcedResult int
	var contained Interval
	var container Interval

	if section1.min == section2.min {
		forcedResult = 1
	} else if section1.min < section2.min {
		forcedResult = 0
		contained = section2
		container = section1
	} else {
		forcedResult = 0
		contained = section1
		container = section2
	}

	if container.max-contained.max >= forcedResult {
		return true
	}

	return false
}
