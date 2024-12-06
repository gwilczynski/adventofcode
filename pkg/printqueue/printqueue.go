package printqueue

import (
	"slices"
	"strconv"
	"strings"
)

func Call(data []string) (correct int, fixed int) {
	rules, updates := RulesAndUpdates(data)

	for _, update := range updates {
		if Valid(update, rules) {
			m := Middle(update)
			correct += m
		} else {
			s := Sort(update, rules)
			m := Middle(s)
			fixed += m
		}
	}

	return correct, fixed
}

func Sort(update []int, rules []Rule) []int {
	for i := 0; i < len(update); i++ {
		for i, u := range update {
			if i == 0 || i == len(update)-1 {
				continue
			}

			for _, rule := range rules {
				if rule.before == u && slices.Contains(update[:i], rule.after) {
					replace(update, i, Before)
				}
				if rule.after == u && slices.Contains(update[i+1:], rule.before) {
					replace(update, i, After)
				}
			}
		}
	}

	return update
}

type Action string

const (
	Before Action = "before"
	After  Action = "after"
)

func replace(slice []int, i int, action Action) []int {
	buffer := slice[i]

	if action == Before {
		slice[i] = slice[i-1]
		slice[i-1] = buffer
	}
	if action == After {
		slice[i] = slice[i+1]
		slice[i+1] = buffer
	}

	return slice
}

func Middle(update []int) int {
	if len(update) == 0 {
		return 0
	}

	i := len(update) / 2

	return update[i]
}

func Valid(update []int, rules []Rule) bool {
	var counter int
	for i, u := range update {
		if i == 0 || i == len(update)-1 {
			continue
		}

		for _, rule := range rules {
			if rule.before == u && slices.Contains(update[:i], rule.after) {
				counter++
			}
			if rule.after == u && slices.Contains(update[i+1:], rule.before) {
				counter++
			}
		}
	}

	return counter == 0
}

type Rule struct {
	before, after int
}

const (
	ModeRule = iota
	ModeUpdate
)

func RulesAndUpdates(data []string) ([]Rule, [][]int) {
	// page ordering rules
	var rules []Rule

	// pages to produce in each update
	var updates [][]int

	mode := ModeRule

	for _, line := range data {
		if line == "" {
			mode = ModeUpdate
			continue
		}

		if mode == ModeRule {
			rule := strings.Split(line, "|")
			rules = append(rules, Rule{before: toInt(rule[0]), after: toInt(rule[1])})
		}

		if mode == ModeUpdate {
			updatesInLine := strings.Split(line, ",")
			updateInt := make([]int, len(updatesInLine))

			for i := 0; i < len(updatesInLine); i++ {
				updateInt[i] = toInt(updatesInLine[i])
			}

			updates = append(updates, updateInt)
		}
	}

	return rules, updates
}

func toInt(value string) (i int) {
	i, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	return i
}
