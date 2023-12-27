package dayX2023

import (
	"bufio"
	"os"
	"strings"
)

func DayX(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	sum := 0

	for {
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}
		}

		lineString = strings.TrimSuffix(lineString, "\n")
	}

	return sum, nil
}
