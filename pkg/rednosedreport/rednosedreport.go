package rednosedreport

import (
	"strconv"
	"strings"
)

func HowManyReportsAreSafe(data []string, problemDampener bool) int {
	var safe int

	for _, line := range data {
		split := Split(line)

		if Safe(split) {
			safe++
		} else if problemDampener && SafeWithProblemDampener(split) {
			safe++
		}
	}

	return safe
}

func SafeWithProblemDampener(data []int) bool {
	for i := 0; i < len(data); i++ {
		slice := make([]int, 0, len(data)-1)
		slice = append(slice, data[:i]...)
		slice = append(slice, data[i+1:]...)

		if Safe(slice) {
			return true
		}
	}

	return false
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
