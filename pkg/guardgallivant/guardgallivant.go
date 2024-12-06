package guardgallivant

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func Call(data []string) int {
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

	// printScreen(fields, guard)
	for i := 0; i < 1_000_000; i++ {
		guard.Flight(fields)
		// printScreen(fields, guard)

		if guard.Stop {
			return guard.Visited
		}
	}

	panic(errors.New("guard did not leave"))
}

func printScreen(fields [][]*Field, guard *Guard) {
	time.Sleep(500 * time.Millisecond)

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	for j := guard.P.J - 10; j < guard.P.J+10; j++ {
		fmt.Println("")

		for i := guard.P.I - 20; i < guard.P.I+20; i++ {
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
	T       FieldType
	Visited bool
}

func (f Field) String() string {
	if f.Visited {
		return "*"
	}

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
	D       Direction
	P       Position
	Visited int
	Stop    bool
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
		return
	}

	if g.D == Right {
		g.D = Down
		return
	}

	if g.D == Down {
		g.D = Left
		return
	}

	if g.D == Left {
		g.D = Up
		return
	}
}

func (g *Guard) Flight(fields [][]*Field) {
	if g.P.J-1 < 0 || g.P.I-1 < 0 || g.P.I+1 >= len(fields[0]) || g.P.J+1 >= len(fields) {
		g.Stop = true
		return
	}

	if g.D == Up {
		f := fields[g.P.J-1][g.P.I]

		if f.T == Free {
			if !f.Visited {
				g.Visited++
			}
			f.Visited = true
			g.P.J--
		} else {
			g.Rotate()
		}
	}
	if g.D == Down {
		f := fields[g.P.J+1][g.P.I]

		if f.T == Free {
			if !f.Visited {
				g.Visited++
			}
			f.Visited = true
			g.P.J++
		} else {
			g.Rotate()
		}
	}
	if g.D == Right {
		f := fields[g.P.J][g.P.I+1]

		if f.T == Free {
			if !f.Visited {
				g.Visited++
			}
			f.Visited = true
			g.P.I++
		} else {
			g.Rotate()
		}
	}
	if g.D == Left {
		f := fields[g.P.J][g.P.I-1]

		if f.T == Free {
			if !f.Visited {
				g.Visited++
			}
			f.Visited = true
			g.P.I--
		} else {
			g.Rotate()
		}
	}
}
