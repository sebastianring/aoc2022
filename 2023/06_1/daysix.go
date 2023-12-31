package daysix2023

import (
	"bufio"
	"fmt"
	"strconv"

	// "fmt"
	"os"
	// "sort"
	// "strconv"
	"strings"
)

type Race struct {
	Time     int
	Distance int
}

func ButtonToWin(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	sum := 0

	times := []int{}
	distances := []int{}

	for {
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			return 0, err
		}

		lineString = strings.TrimSuffix(lineString, "\n")
		split := strings.Split(lineString, ":")

		if split[0] == "Time" {
			inputTimes := strings.Split(split[1], " ")
			for _, time := range inputTimes {
				tempTime := strings.TrimSpace(time)

				if len(tempTime) > 0 {
					tempTime, err := strconv.Atoi(tempTime)

					if err != nil {
						return 0, err
					}

					times = append(times, tempTime)
				}
			}
		} else if split[0] == "Distance" {
			inputDists := strings.Split(split[1], " ")
			for _, dist := range inputDists {
				tempDist := strings.TrimSpace(dist)

				if len(tempDist) > 0 {
					tempDist, err := strconv.Atoi(tempDist)

					if err != nil {
						return 0, err
					}

					distances = append(distances, tempDist)
				}
			}
		}
	}

	races := []Race{}

	for i := 0; i < len(times); i++ {
		race := Race{
			Time:     times[i],
			Distance: distances[i],
		}

		races = append(races, race)
	}

	fmt.Println(races)

	winnables := []int{}

	for _, race := range races {
		waysToWinRace := 0
		for i := 0; i < race.Time; i++ {
			distanceTraveled := DistanceTraveled(race, i)

			if distanceTraveled > race.Distance {
				waysToWinRace++
			}
		}

		if waysToWinRace > 0 {
			winnables = append(winnables, waysToWinRace)
		}
	}

	sum = Multiple(winnables)

	return sum, nil
}

func Multiple(input []int) int {
	var sum int

	for i, val := range input {
		if i == 0 {
			sum = val
			continue
		}

		sum *= val
	}

	return sum
}

func DistanceTraveled(race Race, secondsHeld int) int {
	timeLeft := race.Time - secondsHeld

	return timeLeft * secondsHeld
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}
