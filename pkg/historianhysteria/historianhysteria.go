package historianhysteria

import (
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func TotalDistance(data []string) int {
	result := SplitLines(data, 2)
	result = SortLists(result)
	distances := GetDistances(result)

	return lo.Reduce(distances, func(aggregator, item int, _ int) int {
		return aggregator + item
	}, 0)
}

func SimilarityScore(data []string) int {
	result := SplitLines(data, 2)
	appears := Appearing(result)
	multiplied := Multiplied(result[0], appears)

	return lo.Reduce(multiplied, func(aggregator, item int, _ int) int {
		return aggregator + item
	}, 0)
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

func Appearing(lists [][]int) []int {
	buff := make([]int, len(lists[0]))

	for i := 0; i < len(lists[0]); i++ {
		for j := 0; j < len(lists[1]); j++ {
			if lists[0][i] == lists[1][j] {
				buff[i]++
			}
		}
	}

	return buff
}

func Multiplied(lists []int, appeared []int) []int {
	buff := make([]int, len(lists))

	for i := 0; i < len(lists); i++ {
		buff[i] = lists[i] * appeared[i]
	}

	return buff
}
