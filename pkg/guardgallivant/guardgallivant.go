package guardgallivant

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/inancgumus/screen"
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
	for i := 0; i < 10; i++ {
		guard.Flight(fields)
		printScreen(fields, guard)
	}

	return acc
}

func printScreen(fields [][]*Field, guard *Guard) {
	screen.Clear()
	screen.MoveTopLeft()
	time.Sleep(time.Second)

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

func (g *Guard) String() string {
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
}

func (g *Guard) Rotate() {
	if g.D == Up {
		g.D = Right
	}

	if g.D == Right {
		g.D = Down
	}

	if g.D == Down {
		g.D = Left
	}

	if g.D == Left {
		g.D = Up
	}
}

func (g *Guard) Flight(fields [][]*Field) {
	if g.D == Up {
		if fields[g.P.J-1][g.P.I].T == Free {
			g.P.J--
		} else {
			g.Rotate()
		}
	}
	if g.D == Down {
		if fields[g.P.J+1][g.P.I].T == Free {
			g.P.J++
		} else {
			g.Rotate()
		}
	}
	if g.D == Right {
		if fields[g.P.J][g.P.I+1].T == Free {
			g.P.I++
		} else {
			g.Rotate()
		}
	}
	if g.D == Left {
		if fields[g.P.J][g.P.I-1].T == Free {
			g.P.I--
		} else {
			g.Rotate()
		}
	}
}
