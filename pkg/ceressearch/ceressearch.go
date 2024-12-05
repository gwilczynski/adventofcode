package ceressearch

import (
	"strings"
)

func Call(data []string, scanX bool) int {
	var counter int
	matrix := Matrix(data)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			var result int
			if scanX {
				result = matrix[i][j].ScanX(matrix)
			} else {
				result = matrix[i][j].Scan(matrix)
			}

			counter += result
		}
	}

	return counter
}

type Cell struct {
	Value  string
	Row    int
	Column int
}

func (c *Cell) Scan(matrix [][]*Cell) int {
	numberOfColumns := len(matrix)
	numberOfRows := len(matrix[0])
	var counter int

	directions := []struct {
		dCol, dRow int
	}{
		{0, -1},  // UP
		{1, -1},  // UP-RIGHT
		{1, 0},   // RIGHT
		{1, 1},   // RIGHT-DOWN
		{0, 1},   // DOWN
		{-1, 1},  // DOWN-LEFT
		{-1, 0},  // LEFT
		{-1, -1}, // LEFT-UP
	}

	if c.Value == "X" {
		for _, dir := range directions {
			if matchesPattern(c, dir, matrix, numberOfColumns, numberOfRows) {
				counter++
			}
		}
	}

	return counter
}

func (c *Cell) ScanX(matrix [][]*Cell) (counter int) {
	numberOfColumns := len(matrix)
	numberOfRows := len(matrix[0])

	if !isValidPosition(c.Column-1, c.Row-1, numberOfColumns, numberOfRows) || !isValidPosition(c.Column+1, c.Row+1, numberOfColumns, numberOfRows) {
		return 0
	}

	if c.Value == "A" {
		var ok bool

		sequence := []struct{ start, end string }{
			{
				start: "M", end: "S",
			},
			{
				start: "S", end: "M",
			},
		}

		for _, s := range sequence {
			if matrix[c.Column-1][c.Row-1].Value == s.start && matrix[c.Column+1][c.Row+1].Value == s.end {
				if matrix[c.Column+1][c.Row-1].Value == s.start && matrix[c.Column-1][c.Row+1].Value == s.end {
					ok = true
				}
				if matrix[c.Column+1][c.Row-1].Value == s.end && matrix[c.Column-1][c.Row+1].Value == s.start {
					ok = true
				}
			}
		}

		if ok {
			counter++
		}
	}

	return counter
}

func matchesPattern(c *Cell, dir struct{ dCol, dRow int }, matrix [][]*Cell, numberOfColumns, numberOfRows int) bool {
	col, row := c.Column, c.Row
	sequence := []string{"M", "A", "S"}

	for _, value := range sequence {
		col += dir.dCol
		row += dir.dRow

		if !isValidPosition(col, row, numberOfColumns, numberOfRows) || matrix[col][row].Value != value {
			return false
		}
	}

	return true
}

func isValidPosition(col, row, numberOfColumns, numberOfRows int) bool {
	return col >= 0 && col < numberOfColumns && row >= 0 && row < numberOfRows
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
			}
		}
	}

	return matrix
}
