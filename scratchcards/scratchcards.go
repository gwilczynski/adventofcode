package scratchcards

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

// GetWorthPoints returns How many points they are worth in total?
func GetWorthPoints(data []string) int {
	var acc int
	cards := make([]Card, 0, len(data))
	for _, line := range data {
		card, _ := GetCard(line)
		cards = append(cards, *card)
	}

	for _, card := range cards {
		points := CollectPoints(&card)
		acc += points
	}

	return acc
}

func CollectPoints(card *Card) int {
	matchingNumbers := lo.Filter(card.numbersYouHave, func(x int, _ int) bool {
		return lo.Contains(card.winningNumbers, x)
	})

	return len(matchingNumbers) ^ 2
}

func GetCard(data string) (*Card, error) {
	pattern := `Card (\d+): ([\d\s]+)\| ([\d\s]+)`
	re := regexp.MustCompile(pattern)

	matches := re.FindStringSubmatch(data)
	if len(matches) < 4 {
		return nil, errors.New("no match found")
	}

	cardNumber := matches[1]
	winningNumbersStr := matches[2]
	numbersYouHaveStr := matches[3]

	winningNumbers := convertStringToIntSlice(winningNumbersStr)
	numbersYouHave := convertStringToIntSlice(numbersYouHaveStr)

	return &Card{
		number:         cardNumber,
		winningNumbers: winningNumbers,
		numbersYouHave: numbersYouHave,
	}, nil
}

func convertStringToIntSlice(str string) []int {
	var numbers []int
	for _, numStr := range strings.Fields(str) {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			continue
		}
		numbers = append(numbers, num)
	}
	return numbers
}

type Card struct {
	number         string
	winningNumbers []int
	numbersYouHave []int
}
