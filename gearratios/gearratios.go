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

	var acc int
	for _, number := range numbers {
		if isPartNumber(number, specials) {
			acc += number.value
		}
	}

	return acc
}

func GetGearRatiosSum(data []string) int {
	var numbers []NumberItem
	var specials []SpecialItem

	for i, datum := range data {
		scanItems(i, datum, &numbers, &specials)
	}

	var acc int
	for _, special := range specials {
		if special.value == "*" {
			gearRatios := findGearRatios(special, numbers)
			if len(gearRatios) == 2 {
				acc += gearRatios[0] * gearRatios[1]
			}
		}
	}

	return acc
}

func findGearRatios(special SpecialItem, numbers []NumberItem) []int {
	var gearRatios []int
	for _, number := range numbers {
		if isRelevantNumber(special, number) {
			gearRatios = append(gearRatios, number.value)
		}
	}
	return gearRatios
}

func isRelevantNumber(special SpecialItem, number NumberItem) bool {
	return isLineRelevant(special, number) && isIndexRelevant(special, number)
}

func isLineRelevant(special SpecialItem, number NumberItem) bool {
	return number.line >= special.line-1 && number.line <= special.line+1
}

func isIndexRelevant(special SpecialItem, number NumberItem) bool {
	for offset := -1; offset <= 1; offset++ {
		if indexInRange(special.startIndex+offset, number) {
			return true
		}
	}
	return false
}

func indexInRange(index int, number NumberItem) bool {
	return index >= number.startIndex && index <= number.endIndex
}

func isPartNumber(number NumberItem, specials []SpecialItem) bool {
	for _, special := range specials {
		if isSpecialAdjacentOrOverlapping(special, number) {
			return true
		}
	}
	return false
}

func isSpecialAdjacentOrOverlapping(special SpecialItem, number NumberItem) bool {
	if special.line < number.line-1 || special.line > number.line+1 {
		return false
	}

	return special.startIndex >= number.startIndex-1 && special.startIndex <= number.endIndex+1
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
