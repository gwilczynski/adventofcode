package guardgallivant

import (
	"errors"
	"fmt"
	"strings"
)

func Call(data []string) int {
	var acc int
	var guard *Guard
	var fields [][]*Field

	for j, d := range data {
		l, g := DecodeLine(d, j)

		fields = append(fields, l)

		if g != nil {
			guard = g
		}
	}
	if guard == nil {
		panic(errors.New("guard is nil"))
	}

	printScreen(fields, guard)

	return acc
}

func printScreen(fields [][]*Field, guard *Guard) {
	for j := 0; j < len(fields); j++ {
		fmt.Println("")

		for i := 0; i < len(fields[j]); i++ {
			if guard.P.I == i && guard.P.J == j {
				fmt.Print(guard)
			} else {
				fmt.Print(fields[j][i])
			}
		}
	}
}

func DecodeLine(line string, j int) ([]*Field, *Guard) {
	elements := strings.Split(line, "")
	fields := make([]*Field, len(elements))
	var guard *Guard

	for i, element := range elements {
		if element == "." {
			fields[i] = &Field{T: Free}
		}
		if element == "#" {
			fields[i] = &Field{T: Obstruction}
		}
		if element == "^" {
			fields[i] = &Field{T: Free}
			guard = &Guard{
				D: Up,
				P: Position{
					I: i,
					J: j,
				},
			}
		}
	}

	return fields, guard
}

type FieldType string

const (
	Free        FieldType = "."
	Obstruction FieldType = "#"
)

type Field struct {
	T FieldType
}

func (f Field) String() string {
	return string(f.T)
}

type Direction string

const (
	Up    Direction = "up"
	Down  Direction = "down"
	Right Direction = "right"
	Left  Direction = "left"
)

type Position struct {
	I, J int
}
type Guard struct {
	D Direction
	P Position
}

func (g Guard) String() string {
	var d string
	if g.D == Up {
		d = "^"
	}
	if g.D == Down {
		d = "v"
	}
	if g.D == Right {
		d = ">"
	}
	if g.D == Left {
		d = "<"
	}

	return d
	// return fmt.Sprintf("%s(%d,%d)", d, g.P.I, g.P.J)
}

func (g Guard) Flight() (newPosition Position) {
	return g.P
}
