package template

func Template(lines []string) int {
	sum := 0
	data := FormatData(lines)

	for range data {
	}

	return sum
}

func FormatData(lines []string) []string {
	return lines
}
