package rednosedreport

import (
	"strconv"
	"strings"
)

func HowManyReportsAreSafe(data []string) int {
	var safe int

	for _, line := range data {
		split := Split(line)

		if Safe(split) {
			safe++
		}
	}

	return safe
}

func Safe(data []int) bool {
	increasing := 1
	decreasing := 1

	for i := 0; i < len(data)-1; i++ {
		diff := data[i] - data[i+1]

		if diff >= -3 && diff <= -1 {
			increasing++
		} else if diff >= 1 && diff <= 3 {
			decreasing++
		} else {
			return false
		}
	}

	return len(data) == increasing || len(data) == decreasing
}

func Split(data string) []int {
	fields := strings.Fields(data)
	numbers := make([]int, len(fields))

	for i, field := range fields {
		numbers[i], _ = strconv.Atoi(field)
	}

	return numbers
}
