package trebuchet

import (
	"regexp"
	"strconv"
	"strings"
)

func Decode(doc string) int {
	var accumulator int
	lines := split(doc)
	for _, line := range lines {
		first, last := calibration(line)
		n, _ := number(decode(first), decode(last))

		accumulator += n
	}

	return accumulator
}

func split(doc string) []string {
	return strings.Split(doc, "\n")
}

func calibration(line string) (first string, last string) {
	r := regexp.MustCompile("[0-9]|one|two|three|four|five|six|seven|eight|nine")
	numbers := r.FindAllString(line, -1)

	first = numbers[0]
	last = numbers[len(numbers)-1]

	return
}

func number(first, second string) (int, error) {
	return strconv.Atoi(first + second)
}

func decode(number string) string {
	mapping := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	s := mapping[number]
	if s != "" {
		return s
	}

	return number
}
