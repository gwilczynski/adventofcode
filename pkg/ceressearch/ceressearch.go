package ceressearch

import (
	"strings"
)

func Call(data []string) int {
	var counter int
	matrix := Matrix(data)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if ok := matrix[i][j].Scan(matrix); ok {
				counter++
			}
		}
	}

	return counter
}

type Cell struct {
	Value  string
	Row    int
	Column int
}

func (c Cell) Scan(matrix [][]Cell) bool {
	numberOfColumns := len(matrix)
	numberOfRows := len(matrix[0])

	if c.Value == "X" {

		// UP
		if c.Row-1 >= 0 && matrix[c.Column][c.Row-1].Value == "M" {
			if c.Row-2 >= 0 && matrix[c.Column][c.Row-2].Value == "A" {
				if c.Row-3 >= 0 && matrix[c.Column][c.Row-3].Value == "S" {
					return true
				}
			}
		}

		// DOWN
		if c.Row+1 < numberOfRows && matrix[c.Column][c.Row+1].Value == "M" {
			if c.Row+2 < numberOfRows && matrix[c.Column][c.Row+2].Value == "A" {
				if c.Row+3 < numberOfRows && matrix[c.Column][c.Row+3].Value == "S" {
					return true
				}
			}
		}

		// LEFT
		if c.Column+1 < numberOfColumns && matrix[c.Column+1][c.Row].Value == "M" {
			if c.Column+2 < numberOfColumns && matrix[c.Column+2][c.Row].Value == "A" {
				if c.Column+3 < numberOfColumns && matrix[c.Column+3][c.Row].Value == "S" {
					return true
				}
			}
		}

		// RIGHT
		if c.Column-1 >= 0 && matrix[c.Column-1][c.Row].Value == "M" {
			if c.Column-2 >= 0 && matrix[c.Column-2][c.Row].Value == "A" {
				if c.Column-3 >= 0 && matrix[c.Column-3][c.Row].Value == "S" {
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
			matrix[i][j] = Cell{
				Value:  letters[j],
				Row:    i,
				Column: j,
			}
		}
	}

	return matrix
}
