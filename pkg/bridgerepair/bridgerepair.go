package bridgerepair

import (
	"strconv"
	"strings"
)

func Call(data []string) int {
	var acc int

	for _, puzzle := range Parse(data) {
		if isValid(puzzle.Result, puzzle.Numbers) {
			acc += puzzle.Result
		}
	}

	return acc
}

func isValid(testValue int, numbers []int) bool {
	n := len(numbers)
	if n == 0 {
		return false
	}

	var dfs func(index int, currentValue int) bool
	dfs = func(index int, currentValue int) bool {
		if index == n {
			return currentValue == testValue
		}

		if dfs(index+1, currentValue+numbers[index]) {
			return true
		}

		if dfs(index+1, currentValue*numbers[index]) {
			return true
		}

		return false
	}

	return dfs(1, numbers[0])
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
