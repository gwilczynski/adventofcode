package mullitover

import (
	"regexp"
	"strconv"
	"strings"
)

func Call(data []string) (sum int) {
	records := make([]Num, 0, 1000)

	for _, line := range data {
		instructions := ExtractInstructions(line)
		for _, instruction := range instructions {
			n := ExtractNumbers(instruction)

			records = append(records, n)
		}
	}

	for _, record := range records {
		sum += record.one * record.two
	}

	return
}

const pattern = `mul\(\d+,\d+\)`

func ExtractInstructions(input string) []string {
	r := regexp.MustCompile(pattern)

	return r.FindAllString(input, -1)
}

func ExtractNumbers(input string) Num {
	trimmed := strings.TrimPrefix(input, "mul(")
	trimmed = strings.TrimSuffix(trimmed, ")")
	parts := strings.Split(trimmed, ",")

	num1, _ := strconv.Atoi(parts[0])
	num2, _ := strconv.Atoi(parts[1])

	return Num{num1, num2}
}

type Num struct {
	one int
	two int
}
