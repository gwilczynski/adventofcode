package cubeconundrum

import (
	"regexp"
	"strconv"
	"strings"
)

type Color uint

const (
	Red Color = iota
	Green
	Blue
)

func (s Color) String() string {
	switch s {
	case Red:
		return "red"
	case Green:
		return "green"
	case Blue:
		return "blue"
	}
	return "unknown"
}

func FromString(color string) Color {
	switch color {
	case "red":
		return Red
	case "green":
		return Green
	case "blue":
		return Blue
	default:
		return Color(0)
	}
}

type (
	Set  map[Color]int
	Bag  map[Color]int
	Game struct {
		id   int
		sets []Set
	}
)

func WhichGamesPossible(doc string, bag Bag) int {
	var accumulator int

	lines := split(doc)
	games := parse(lines)
	for _, game := range games {
		if isPossible(game, bag) {
			accumulator += game.id
		}
	}

	return accumulator
}

func isPossible(game Game, bag Bag) (possible bool) {
	colours := []Color{Green, Blue, Red}
	for _, set := range game.sets {
		for _, colour := range colours {
			value, exists := set[colour]

			possible = exists == false || bag[colour] >= value
			if !possible {
				return false
			}
		}
	}

	return
}

func split(doc string) []string {
	return strings.Split(doc, "\n")
}

func parse(games []string) []Game {
	gamesParsed := make([]Game, len(games))

	for i, game := range games {
		gamesParsed[i] = buildGame(game)
	}

	return gamesParsed
}

func buildGame(g string) Game {
	parts := strings.Split(g, ":")
	label := strings.Split(parts[0], " ")
	setsAsStrings := strings.Split(parts[1], ";")

	sets := make([]Set, len(setsAsStrings))
	for i := 0; i < len(sets); i++ {
		var set Set = make(Set)

		parts := strings.Split(setsAsStrings[i], ",")
		for _, part := range parts {
			color, number := parseSet(part)

			set[color] = number
		}

		sets[i] = set
	}

	id, _ := strconv.Atoi(label[1])
	return Game{id: id, sets: sets}
}

func parseSet(s string) (color Color, number int) {
	re := regexp.MustCompile(`(\d+)\s+(red|blue|green)`)
	matches := re.FindStringSubmatch(s)

	number, _ = strconv.Atoi(matches[1])
	color = FromString(matches[2])

	return
}
