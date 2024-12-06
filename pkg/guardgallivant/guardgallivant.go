package guardgallivant

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func Call(data []string) int {
	fieldsInit, _ := build(data)

	for k := 0; k < len(fieldsInit[0]); k++ {
		for l := 0; l < len(fieldsInit); l++ {
			fields, guard := build(data)
			if fields[k][l].T == Obstruction {
				continue
			}

			fields[k][l].T = Obstruction

			printScreen(fields, guard)
			for i := 0; i < 1_000; i++ {
				guard.Flight(fields)
				printScreen(fields, guard)

				if guard.Stop {
					// return guard.Visited
					break
				}
			}
		}
	}

	panic(errors.New("guard did not leave"))
}

func build(data []string) (fields [][]*Field, guard *Guard) {
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

	return fields, guard
}

func printScreen(fields [][]*Field, guard *Guard) {
	time.Sleep(100 * time.Millisecond)

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	for j := 0; j < len(fields[0]); j++ {
		fmt.Println("")

		for i := 0; i < len(fields); i++ {
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

type PathDirection string

const (
	PathUp    PathDirection = "|"
	PathDown  PathDirection = "|"
	PathLeft  PathDirection = "-"
	PathRight PathDirection = "-"
	PathTurn  PathDirection = "+"
)

type Field struct {
	T       FieldType
	Visited bool
	Path    PathDirection
}

func (f Field) String() string {
	if f.Visited {
		return string(f.Path)
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
	switch g.D {
	case Up:
		return "^"
	case Down:
		return "v"
	case Right:
		return ">"
	case Left:
		return "<"
	default:
		return ""
	}
}

func (g *Guard) Rotate() {
	switch g.D {
	case Up:
		g.D = Right
	case Right:
		g.D = Down
	case Down:
		g.D = Left
	case Left:
		g.D = Up
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
			f.Path = PathUp
			g.P.J--
		} else {
			fields[g.P.J][g.P.I].Path = PathTurn
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
			f.Path = PathDown
			g.P.J++
		} else {
			fields[g.P.J][g.P.I].Path = PathTurn
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
			f.Path = PathRight
			g.P.I++
		} else {
			fields[g.P.J][g.P.I].Path = PathTurn
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
			f.Path = PathLeft
			g.P.I--
		} else {
			fields[g.P.J][g.P.I].Path = PathTurn
			g.Rotate()
		}
	}
}
