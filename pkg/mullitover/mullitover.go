package mullitover

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Call(data []string) (sum int) {
	records := make([]Num, 0, 1000)

	for _, line := range data {
		instructions := ExtractInstructionsCombined(line)
		do := true

		for _, instruction := range instructions {
			if instruction == "do()" {
				do = true
				continue
			} else if instruction == "don't()" {
				do = false
				continue
			}

			if do {
				n := ExtractNumbers(instruction)
				records = append(records, n)
			}
		}
	}

	for _, record := range records {
		sum += record.one * record.two
	}

	return
}

const (
	patternMul  = `mul\(\d+,\d+\)`
	patternDo   = `do\(\)`
	patternDont = `don\'t\(\)`
)

func ExtractInstructions(input string) []string {
	r := regexp.MustCompile(patternMul)

	return r.FindAllString(input, -1)
}

func ExtractInstructionsCombined(input string) []string {
	combinedPattern := fmt.Sprintf(`%s|%s|%s`, patternMul, patternDo, patternDont)
	r := regexp.MustCompile(combinedPattern)

	return r.FindAllString(input, -1)
}

func ExtractNumbers(input string) Num {
	trimmed := strings.TrimPrefix(input, "mul(")
	trimmed = strings.TrimSuffix(trimmed, ")")
	parts := strings.Split(trimmed, ",")

	num1, _ := strconv.Atoi(parts[0])
	num2, _ := strconv.Atoi(parts[1])

	return Num{num1, num2}
}

type Num struct {
	one int
	two int
}
