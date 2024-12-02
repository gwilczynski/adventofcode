package historianhysteria

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func TotalDistance(data []string) int {
	var distance int

	result := SplitLines(data, 2)
	result = SortLists(result)

	fmt.Println(result)

	// TODO, calculate distance between each pair

	// TODO, sum all distances

	return distance
}

func SplitLines(lines []string, numberOfColumns int) [][]int {
	result := make([][]int, numberOfColumns)

	for _, line := range lines {
		parts := strings.Fields(line)

		if len(parts) == 2 {
			first, _ := strconv.Atoi(parts[0])
			second, _ := strconv.Atoi(parts[1])

			result[0] = append(result[0], first)
			result[1] = append(result[1], second)
		}
	}

	return result
}

func SortLists(lists [][]int) [][]int {
	for _, list := range lists {
		slices.Sort(list)
	}

	return lists
}

func GetDistances(lists [][]int) []int {
	buff := make([]int, len(lists[0]))

	for _, list := range lists {
		for i := 0; i < len(list); i++ {
			buff[i] = int(math.Abs(float64(buff[i] - list[i])))
		}
	}

	return buff
}
