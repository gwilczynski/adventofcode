package bridgerepair

import (
	"strconv"
	"strings"
)

func Call(data []string) int {
	return len(data)
}

type Puzzle struct {
	Result  int
	Numbers []int
}

func Parse(data []string) []*Puzzle {
	puzzles := make([]*Puzzle, len(data))

	for i, d := range data {
		puzzles[i] = ParseLine(d)
	}

	return puzzles
}

func ParseLine(line string) *Puzzle {
	parts := strings.Split(line, ":")
	result, _ := strconv.Atoi(parts[0])

	rest := strings.Split(strings.TrimSpace(parts[1]), " ")
	numbers := make([]int, len(rest))

	for i, s := range rest {
		numbers[i], _ = strconv.Atoi(s)
	}

	return &Puzzle{
		Result:  result,
		Numbers: numbers,
	}
}
