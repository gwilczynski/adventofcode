package ceressearch

import (
	"fmt"
	"strings"
)

func Call(data []string) int {
	var counter int
	matrix := Matrix(data)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			result := matrix[i][j].Scan(matrix)
			if result > 0 {
				fmt.Println(result)
			}

			counter += result
		}
	}

	return counter
}

type CellType string

const (
	Up        CellType = "UP"
	UpRight   CellType = "UP-RIGHT"
	Right     CellType = "RIGHT"
	RightDown CellType = "RIGHT-DOWN"
	Down      CellType = "DOWN"
	DownLeft  CellType = "DOWN-LEFT"
	Left      CellType = "LEFT"
	LeftUp    CellType = "LEFT-UP"
)

type Cell struct {
	Value  string
	Row    int
	Column int
	Types  []CellType
}

func (c *Cell) Scan(matrix [][]*Cell) int {
	numberOfColumns := len(matrix)
	numberOfRows := len(matrix[0])
	var counter int

	if c.Value == "X" {

		// UP
		if c.Row-1 >= 0 && matrix[c.Column][c.Row-1].Value == "M" {
			if c.Row-2 >= 0 && matrix[c.Column][c.Row-2].Value == "A" {
				if c.Row-3 >= 0 && matrix[c.Column][c.Row-3].Value == "S" {
					c.Types = append(c.Types, Up)
					counter++
				}
			}
		}
		// UP-RIGHT
		if c.Column+1 < numberOfColumns && c.Row-1 >= 0 && matrix[c.Column+1][c.Row-1].Value == "M" {
			if c.Column+2 < numberOfColumns && c.Row-2 >= 0 && matrix[c.Column+2][c.Row-2].Value == "A" {
				if c.Column+3 < numberOfColumns && c.Row-3 >= 0 && matrix[c.Column+3][c.Row-3].Value == "S" {
					c.Types = append(c.Types, UpRight)
					counter++
				}
			}
		}

		// RIGHT
		if c.Column+1 < numberOfColumns && matrix[c.Column+1][c.Row].Value == "M" {
			if c.Column+2 < numberOfColumns && matrix[c.Column+2][c.Row].Value == "A" {
				if c.Column+3 < numberOfColumns && matrix[c.Column+3][c.Row].Value == "S" {
					c.Types = append(c.Types, Right)
					counter++
				}
			}
		}

		// RIGHT-DOWN
		if c.Column+1 < numberOfColumns && c.Row+1 < numberOfRows && matrix[c.Column+1][c.Row+1].Value == "M" {
			if c.Column+2 < numberOfColumns && c.Row+2 < numberOfRows && matrix[c.Column+2][c.Row+2].Value == "A" {
				if c.Column+3 < numberOfColumns && c.Row+3 < numberOfRows && matrix[c.Column+3][c.Row+3].Value == "S" {
					c.Types = append(c.Types, RightDown)
					counter++
				}
			}
		}

		// DOWN
		if c.Row+1 < numberOfRows && matrix[c.Column][c.Row+1].Value == "M" {
			if c.Row+2 < numberOfRows && matrix[c.Column][c.Row+2].Value == "A" {
				if c.Row+3 < numberOfRows && matrix[c.Column][c.Row+3].Value == "S" {
					c.Types = append(c.Types, Down)
					counter++
				}
			}
		}

		// DOWN-LEFT
		if c.Column-1 >= 0 && c.Row+1 < numberOfRows && matrix[c.Column-1][c.Row+1].Value == "M" {
			if c.Column-2 >= 0 && c.Row+2 < numberOfRows && matrix[c.Column-2][c.Row+2].Value == "A" {
				if c.Column-3 >= 0 && c.Row+3 < numberOfRows && matrix[c.Column-3][c.Row+3].Value == "S" {
					c.Types = append(c.Types, DownLeft)
					counter++
				}
			}
		}

		// LEFT
		if c.Column-1 >= 0 && matrix[c.Column-1][c.Row].Value == "M" {
			if c.Column-2 >= 0 && matrix[c.Column-2][c.Row].Value == "A" {
				if c.Column-3 >= 0 && matrix[c.Column-3][c.Row].Value == "S" {
					c.Types = append(c.Types, Left)
					counter++
				}
			}
		}

		// LEFT-UP
		if c.Column-1 >= 0 && c.Row-1 >= 0 && matrix[c.Column-1][c.Row-1].Value == "M" {
			if c.Column-2 >= 0 && c.Row-2 >= 0 && matrix[c.Column-2][c.Row-2].Value == "A" {
				if c.Column-3 >= 0 && c.Row-3 >= 0 && matrix[c.Column-3][c.Row-3].Value == "S" {
					c.Types = append(c.Types, LeftUp)
					counter++
				}
			}
		}
	}

	return counter
}

func Matrix(data []string) [][]*Cell {
	matrix := make([][]*Cell, len(data))

	for i := 0; i < len(data); i++ {
		letters := strings.Split(data[i], "")
		matrix[i] = make([]*Cell, len(letters))

		for j := 0; j < len(letters); j++ {
			matrix[i][j] = &Cell{
				Value:  letters[j],
				Row:    j,
				Column: i,
				Types:  make([]CellType, 0),
			}
		}
	}

	return matrix
}
