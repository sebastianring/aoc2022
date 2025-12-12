package fourtwo

import (
	"bufio"
	"os"
	"strings"
)

func Solution(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	lines := GetLines(reader)
	return FourOneSolution(lines), nil
}

func GetLines(reader *bufio.Reader) []string {
	lines := []string{}

	for {
		lineString, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
		}

		lineString = strings.TrimSuffix(lineString, "\n")
		lines = append(lines, lineString)
	}

	return lines
}
