package ceressearch

import (
	"fmt"
	"strings"
)

func Call(data []string) int {
	fmt.Printf("Calling: %d\n", len(data))

	return 0
}

type Cell struct {
	Value  string
	Row    int
	Column int
}

func (c Cell) Scan(matrix [][]Cell) bool {
	if c.Value == "X" {

		// UP
		if matrix[c.Column][c.Row-1].Value == "M" {
			if matrix[c.Column][c.Row-2].Value == "A" {
				if matrix[c.Column][c.Row-3].Value == "S" {
					return true
				}
			}
		}

		// DOWN
		if matrix[c.Column][c.Row+1].Value == "M" {
			if matrix[c.Column][c.Row+2].Value == "A" {
				if matrix[c.Column][c.Row+3].Value == "S" {
					return true
				}
			}
		}

		// LEFT
		if matrix[c.Column+1][c.Row].Value == "M" {
			if matrix[c.Column+2][c.Row].Value == "A" {
				if matrix[c.Column+3][c.Row].Value == "S" {
					return true
				}
			}
		}

		// RIGHT
		if matrix[c.Column-1][c.Row].Value == "M" {
			if matrix[c.Column-2][c.Row].Value == "A" {
				if matrix[c.Column-3][c.Row].Value == "S" {
					return true
				}
			}
		}
	}

	return false
}

func Matrix(data []string) [][]Cell {
	matrix := make([][]Cell, len(data))

	for i := 0; i < len(data); i++ {
		letters := strings.Split(data[i], "")
		matrix[i] = make([]Cell, len(letters))

		for j := 0; j < len(letters); j++ {
			matrix[i][j].Value = letters[j]
			matrix[i][j].Row = i
			matrix[i][j].Column = j
		}
	}

	return matrix
}
