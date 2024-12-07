package bridgerepair

import (
	"slices"
	"strconv"
	"strings"
)

func Call(data []string) int {
	puzzles := Parse(data)

	return len(puzzles)
}

func Eval(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	results := make(map[string]int)

	var helper func(expression string, index int, currentResult int, lastOperand int, nums []int)
	helper = func(expression string, index int, currentResult int, lastOperand int, nums []int) {
		if index == len(nums) {
			results[expression] = currentResult
			return
		}

		helper(
			expression+"+"+strconv.Itoa(nums[index]),
			index+1,
			currentResult+nums[index],
			nums[index],
			nums,
		)

		helper(
			expression+"*"+strconv.Itoa(nums[index]),
			index+1,
			currentResult-lastOperand+lastOperand*nums[index],
			lastOperand*nums[index],
			nums,
		)
	}

	helper(strconv.Itoa(nums[0]), 1, nums[0], nums[0], nums)

	var rr []int
	for _, r := range results {
		rr = append(rr, r)
	}

	slices.Sort(rr)
	return rr
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
