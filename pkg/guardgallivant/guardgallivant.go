package guardgallivant

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func Call(data []string) int {
	fieldsInit, _ := build(data)
	var wasThere int

	for j := 0; j < len(fieldsInit); j++ {
		for i := 0; i < len(fieldsInit[0]); i++ {
			fields, guard := build(data)
			if fields[j][i].T == Obstruction {
				continue
			}
			if guard.P.I == i && guard.P.J == j {
				continue
			}

			fields[j][i].T = NewObstruction

			for iter := 0; iter < 100_000_000; iter++ {
				guard.Flight(fields)

				if guard.Stop {
					if guard.Loop > 1 {
						printScreen(fields, guard)
						wasThere++
					}

					break
				}
			}
		}
	}

	return wasThere
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
	time.Sleep(50 * time.Millisecond)

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
	fmt.Println("")
	fmt.Println("(", guard.P.I, ",", guard.P.J, ")")
	fmt.Println("")
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
	Free           FieldType = "."
	Obstruction    FieldType = "#"
	NewObstruction FieldType = "O"
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
	Hit     int
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
	Loop    int
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
		if f.T == NewObstruction {
			if g.Loop > 1 {
				g.Stop = true
			}
			g.Loop++
		}

		if f.T == Free {
			if !f.Visited {
				g.Visited++
			}

			f.Visited = true
			f.Hit++
			f.Path = PathUp
			g.P.J--
		} else {
			fields[g.P.J][g.P.I].Path = PathTurn
			g.Rotate()
		}
	}

	if g.D == Down {
		f := fields[g.P.J+1][g.P.I]
		if f.T == NewObstruction {
			if g.Loop > 1 {
				g.Stop = true
			}
			g.Loop++
		}

		if f.T == Free {
			if !f.Visited {
				g.Visited++
			}

			f.Visited = true
			f.Hit++
			f.Path = PathDown
			g.P.J++
		} else {
			fields[g.P.J][g.P.I].Path = PathTurn
			g.Rotate()
		}
	}

	if g.D == Right {
		f := fields[g.P.J][g.P.I+1]
		if f.T == NewObstruction {
			if g.Loop > 1 {
				g.Stop = true
			}
			g.Loop++
		}

		if f.T == Free {
			if !f.Visited {
				g.Visited++
			}

			f.Visited = true
			f.Hit++
			f.Path = PathRight
			g.P.I++
		} else {
			fields[g.P.J][g.P.I].Path = PathTurn
			g.Rotate()
		}
	}

	if g.D == Left {
		f := fields[g.P.J][g.P.I-1]
		if f.T == NewObstruction {
			if g.Loop > 1 {
				g.Stop = true
			}
			g.Loop++
		}

		if f.T == Free {
			if !f.Visited {
				g.Visited++
			}

			f.Visited = true
			f.Hit++
			f.Path = PathLeft
			g.P.I--
		} else {
			fields[g.P.J][g.P.I].Path = PathTurn
			g.Rotate()
		}
	}
}
