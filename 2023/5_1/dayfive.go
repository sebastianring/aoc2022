package dayfive2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RangeValues struct {
	Dest      int
	Source    int
	Length    int
	SourceMin int
	SourceMax int
}

type SeedRange struct {
	Min int
	Max int
}

func SeedToLocation(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	sum := 999100000000

	var seeds []int
	convertMap := make(map[string][]RangeValues)
	var currentString string

	// sequenceMap := map[int]string{
	// 	1: "seed-to-soil",
	// 	2: "soil-to-fertilizer",
	// 	3: "fertilizer-to-water",
	// 	4: "water-to-light",
	// 	5: "light-to-temperature",
	// 	6: "temperature-to-humidity",
	// 	7: "humidity-to-location",
	// }

	for {
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			return 0, err
		}

		lineString = strings.TrimSuffix(lineString, "\n")
		titles := strings.Split(lineString, ":")

		if titles[0] == "" {
			_, ok := convertMap[currentString]

			if ok {
				for i := 0; i < len(seeds); i++ {
					dest, err := GetDestFromSource(seeds[i], convertMap[currentString])

					if err != nil {
						return 0, err
					}

					fmt.Printf("I: %v, Map: %v, Source: %v, Destination: %v \n", i, currentString, seeds[i], dest)

					seeds[i] = dest
				}
			}

			currentString = ""
			continue
		}

		values := strings.Split(lineString, " ")

		if currentString == "" {
			switch titles[0] {
			case "seeds":
				tempseeds := strings.Trim(titles[1], " ")
				stringSeeds := strings.Split(tempseeds, " ")

				for _, seed := range stringSeeds {
					intSeed, err := strconv.Atoi(seed)

					if err != nil {
						return 0, err
					}

					seeds = append(seeds, intSeed)
				}
			default:
				currentString = values[0]
			}
		} else {
			source, err := strconv.Atoi(values[1])

			if err != nil {
				return 0, err
			}

			dest, err := strconv.Atoi(values[0])

			if err != nil {
				return 0, err
			}

			length, err := strconv.Atoi(values[2])

			if err != nil {
				return 0, nil
			}

			rv := RangeValues{
				Source: source,
				Dest:   dest,
				Length: length,
			}

			convertMap[currentString] = append(convertMap[currentString], rv)

			// for _, seed := range seeds {
			// 	dest, err := GetDestFromSource(seed, rv)
			//
			// 	if err != nil {
			// 		return 0, err
			// 	}
			//
			// }
		}
	}

	fmt.Printf("Seeds: %v \n", seeds)
	// for k, v := range convertMap {
	// 	fmt.Printf("Key: %v, Value: %v \n", k, v)
	// }

	for _, seed := range seeds {
		if seed < sum {
			sum = seed
		}
	}

	return sum, nil
}

func SeedToLocationPartTwo(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	sum := 999100000000

	// var seeds []int
	var seedRanges []SeedRange
	convertMap := make(map[string][]RangeValues)
	var currentString string

	// sequenceMap := map[int]string{
	// 	1: "seed-to-soil",
	// 	2: "soil-to-fertilizer",
	// 	3: "fertilizer-to-water",
	// 	4: "water-to-light",
	// 	5: "light-to-temperature",
	// 	6: "temperature-to-humidity",
	// 	7: "humidity-to-location",
	// }

	for {
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			return 0, err
		}

		lineString = strings.TrimSuffix(lineString, "\n")
		titles := strings.Split(lineString, ":")

		if titles[0] == "" && currentString != "" {
			_, ok := convertMap[currentString]

			if ok {
				fmt.Println("Found this convertMap: ", convertMap[currentString])
				seedRangeLength := len(seedRanges)

				for i := 0; i < seedRangeLength; i++ {
					dests, err := GetDestRangesFromSourceRange(seedRanges[i], convertMap[currentString])

					if err != nil {
						return 0, err
					}

					seedRanges = append(seedRanges, dests...)
					// fmt.Printf("I: %v, Map: %v, Source: %v, Destination: %v Length Seedranges: %v \n", i, currentString, seedRanges[i], dests, len(seedRanges))
				}
			}

			currentString = ""
			continue
		}

		values := strings.Split(lineString, " ")

		if currentString == "" {
			switch titles[0] {
			case "seeds":
				tempseeds := strings.Trim(titles[1], " ")
				stringSeeds := strings.Split(tempseeds, " ")

				seedRanges, err = GetAllSeedRanges(stringSeeds)

				if err != nil {
					return 0, err
				}

			default:
				currentString = values[0]
			}
		} else {
			source, err := strconv.Atoi(values[1])

			if err != nil {
				return 0, err
			}

			dest, err := strconv.Atoi(values[0])

			if err != nil {
				return 0, err
			}

			length, err := strconv.Atoi(values[2])

			if err != nil {
				return 0, nil
			}

			rv := RangeValues{
				Source:    source,
				Dest:      dest,
				Length:    length,
				SourceMin: source,
				SourceMax: source + length,
			}

			convertMap[currentString] = append(convertMap[currentString], rv)

		}
	}

	// fmt.Printf("SeedRanges: %v \n", seedRanges)
	for k, v := range convertMap {
		fmt.Printf("Key: %v, Value: %v \n", k, v)
	}
	//
	// for _, seed := range seedRanges {
	// 	if seed < sum {
	// 		sum = seed
	// 	}
	// }

	return sum, nil
}

func GetDestFromSource(source int, rvs []RangeValues) (int, error) {
	var dest int

	for _, rv := range rvs {
		if source >= rv.Source && source <= rv.Source+rv.Length {
			fmt.Printf("Using these values: %v", rv)
			diff := source - rv.Source
			fmt.Println(diff)

			dest = rv.Dest + diff
			return dest, nil
		}
	}

	return source, nil
}

func GetDestRangesFromSourceRange(sourceRange SeedRange, rvs []RangeValues) ([]SeedRange, error) {
	var dests []SeedRange

	for _, rv := range rvs {
		overlap, err := GetOverlapRanges(sourceRange, rv)

		if err != nil {
			return dests, err
		}

		dests = append(dests, overlap)
	}

	// fmt.Println("DESTS: ", dests)

	return dests, nil
}

func GetOverlapRanges(input SeedRange, comparable RangeValues) (SeedRange, error) {
	result := SeedRange{}
	// fmt.Printf("Input: %v Comparable: %v\n", input, comparable)

	result.Min = max(input.Min, comparable.SourceMin)
	result.Max = min(input.Max, comparable.SourceMax)

	// fmt.Println("Found an overlap: ", result.Min, result.Max)

	return result, nil
}

func GetAllSeeds(input []string) ([]int, error) {
	var result []int

	for i := 0; i < len(input); i += 2 {
		// fmt.Printf("%v ", i)
		base, err := strconv.Atoi(input[i])

		if err != nil {
			return result, err
		}

		end, err := strconv.Atoi(input[i+1])

		if err != nil {
			return result, nil
		}

		for j := 0; j < end; j++ {
			// fmt.Printf("%v ", j)
			result = append(result, base+j)
		}
	}

	// fmt.Println("All seeds: ", result)
	return result, nil
}

func GetAllSeedRanges(input []string) ([]SeedRange, error) {
	var result []SeedRange

	for i := 0; i < len(input); i += 2 {
		// fmt.Printf("%v ", i)
		mini, err := strconv.Atoi(input[i])

		if err != nil {
			return result, err
		}

		length, err := strconv.Atoi(input[i+1])

		if err != nil {
			return result, nil
		}

		sr := SeedRange{
			Min: mini,
			Max: mini + length,
		}

		result = append(result, sr)
	}

	fmt.Println("Found all seed ranges: ", result)

	return result, nil
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
