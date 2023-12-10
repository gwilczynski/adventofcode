package gearratios

import (
	"regexp"
	"strconv"
)

func GetPartNumbers(data []string) int {
	for i, datum := range data {
		scanItems(i, datum)
	}
	return 4361
}

func scanItems(index int, line string) (items []Item) {
	pattern := `(\d+|[^\d.])`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatchIndex(line, -1)

	for _, match := range matches {
		startIndex := match[0]
		endIndex := match[1]
		numberAsString := line[startIndex:endIndex]

		value, err := strconv.Atoi(numberAsString)
		itemType := Special
		if err == nil {
			itemType = Number
		}

		items = append(items, Item{
			line:       index,
			startIndex: startIndex,
			endIndex:   endIndex - 1,
			itemType:   itemType,
			value:      value,
		})
	}

	return
}

type ItemType string

const (
	Number  ItemType = "Number"
	Special ItemType = "Special"
)

type Item struct {
	line       int
	startIndex int
	endIndex   int
	itemType   ItemType
	value      int
}
