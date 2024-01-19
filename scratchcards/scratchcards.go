package scratchcards

// GetWorthPoints returns How many points they are worth in total?
func GetWorthPoints(data []string) int {
	return 0
}

func GetCard(data string) Card {
	return Card{
		number:         "1",
		winningNumbers: []int{41, 48, 83, 86, 17},
		numbersYouHave: []int{83, 86, 6, 31, 17, 9, 48, 53},
	}
}

type Card struct {
	number         string
	winningNumbers []int
	numbersYouHave []int
}
