package gearratios

import (
	"regexp"
	"strconv"
)

func GetPartNumbers(data []string) int {
	var numbers []NumberItem
	var specials []SpecialItem

	for i, datum := range data {
		scanItems(i, datum, &numbers, &specials)
	}

	return 4361
}

func findIndexedAround(number NumberItem) map[int][]int {
	coordinates := map[int][]int{}

	for i := 0; i < 3; i++ {
		var indexes []int

		indexes = append(indexes, number.startIndex-1)
		for j := number.startIndex; j <= number.endIndex; j++ {
			indexes = append(indexes, j)
		}
		indexes = append(indexes, number.endIndex+1)

		coordinates[number.line-1] = indexes
		coordinates[number.line] = indexes
		coordinates[number.line+1] = indexes
	}

	return coordinates
}

func scanItems(index int, line string, numbers *[]NumberItem, specials *[]SpecialItem) {
	pattern := `(\d+|[^\d.])`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatchIndex(line, -1)

	for _, match := range matches {
		startIndex := match[0]
		endIndex := match[1]
		valueAsString := line[startIndex:endIndex]

		value, err := strconv.Atoi(valueAsString)
		if err == nil {
			*numbers = append(*numbers, NumberItem{
				line:       index,
				startIndex: startIndex,
				endIndex:   endIndex - 1,
				value:      value,
			})
		} else {
			*specials = append(*specials, SpecialItem{
				line:       index,
				startIndex: startIndex,
				endIndex:   endIndex - 1,
				value:      valueAsString,
			})
		}
	}
}

type NumberItem struct {
	line       int
	startIndex int
	endIndex   int
	value      int
}

type SpecialItem struct {
	line       int
	startIndex int
	endIndex   int
	value      string
}
