package ceressearch_test

import (
	"reflect"
	"testing"

	"github.com/gwilczynski/adventofcode/pkg/ceressearch"
)

func TestCall(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ok",
			args: args{
				data: []string{
					"MMMSXXMASM",
					"MSAMXMSMSA",
					"AMXSXMAAMM",
					"MSAMASMSMX",
					"XMASAMXAMM",
					"XXAMMXXAMA",
					"SMSMSASXSS",
					"SAXAMASAAA",
					"MAMMMXMMMM",
					"MXMXAXMASX",
				},
			},
			want: 18,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ceressearch.Call(tt.args.data); got != tt.want {
				t.Errorf("Call() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
		want [][]ceressearch.Cell
	}{
		{
			name: "ok",
			args: args{
				data: []string{
					"MMMSXXMASM",
					"MSAMXMSMSA",
					"AMXSXMAAMM",
					"MSAMASMSMX",
					"XMASAMXAMM",
					"XXAMMXXAMA",
					"SMSMSASXSS",
					"SAXAMASAAA",
					"MAMMMXMMMM",
					"MXMXAXMASX",
				},
			},
			want: [][]ceressearch.Cell{
				{
					ceressearch.Cell{Value: "M"},
					ceressearch.Cell{Value: "M", Column: 1},
					ceressearch.Cell{Value: "M", Column: 2},
					ceressearch.Cell{Value: "S", Column: 3},
					ceressearch.Cell{Value: "X", Column: 4},
					ceressearch.Cell{Value: "X", Column: 5},
					ceressearch.Cell{Value: "M", Column: 6},
					ceressearch.Cell{Value: "A", Column: 7},
					ceressearch.Cell{Value: "S", Column: 8},
					ceressearch.Cell{Value: "M", Column: 9},
				},
				{
					ceressearch.Cell{Value: "M", Row: 1},
					ceressearch.Cell{Value: "S", Row: 1, Column: 1},
					ceressearch.Cell{Value: "A", Row: 1, Column: 2},
					ceressearch.Cell{Value: "M", Row: 1, Column: 3},
					ceressearch.Cell{Value: "X", Row: 1, Column: 4},
					ceressearch.Cell{Value: "M", Row: 1, Column: 5},
					ceressearch.Cell{Value: "S", Row: 1, Column: 6},
					ceressearch.Cell{Value: "M", Row: 1, Column: 7},
					ceressearch.Cell{Value: "S", Row: 1, Column: 8},
					ceressearch.Cell{Value: "A", Row: 1, Column: 9},
				},
				{
					ceressearch.Cell{Value: "A", Row: 2},
					ceressearch.Cell{Value: "M", Row: 2, Column: 1},
					ceressearch.Cell{Value: "X", Row: 2, Column: 2},
					ceressearch.Cell{Value: "S", Row: 2, Column: 3},
					ceressearch.Cell{Value: "X", Row: 2, Column: 4},
					ceressearch.Cell{Value: "M", Row: 2, Column: 5},
					ceressearch.Cell{Value: "A", Row: 2, Column: 6},
					ceressearch.Cell{Value: "A", Row: 2, Column: 7},
					ceressearch.Cell{Value: "M", Row: 2, Column: 8},
					ceressearch.Cell{Value: "M", Row: 2, Column: 9},
				},
				{
					ceressearch.Cell{Value: "M", Row: 3},
					ceressearch.Cell{Value: "S", Row: 3, Column: 1},
					ceressearch.Cell{Value: "A", Row: 3, Column: 2},
					ceressearch.Cell{Value: "M", Row: 3, Column: 3},
					ceressearch.Cell{Value: "A", Row: 3, Column: 4},
					ceressearch.Cell{Value: "S", Row: 3, Column: 5},
					ceressearch.Cell{Value: "M", Row: 3, Column: 6},
					ceressearch.Cell{Value: "S", Row: 3, Column: 7},
					ceressearch.Cell{Value: "M", Row: 3, Column: 8},
					ceressearch.Cell{Value: "X", Row: 3, Column: 9},
				},
				{
					ceressearch.Cell{Value: "X", Row: 4},
					ceressearch.Cell{Value: "M", Row: 4, Column: 1},
					ceressearch.Cell{Value: "A", Row: 4, Column: 2},
					ceressearch.Cell{Value: "S", Row: 4, Column: 3},
					ceressearch.Cell{Value: "A", Row: 4, Column: 4},
					ceressearch.Cell{Value: "M", Row: 4, Column: 5},
					ceressearch.Cell{Value: "X", Row: 4, Column: 6},
					ceressearch.Cell{Value: "A", Row: 4, Column: 7},
					ceressearch.Cell{Value: "M", Row: 4, Column: 8},
					ceressearch.Cell{Value: "M", Row: 4, Column: 9},
				},
				{
					ceressearch.Cell{Value: "X", Row: 5},
					ceressearch.Cell{Value: "X", Row: 5, Column: 1},
					ceressearch.Cell{Value: "A", Row: 5, Column: 2},
					ceressearch.Cell{Value: "M", Row: 5, Column: 3},
					ceressearch.Cell{Value: "M", Row: 5, Column: 4},
					ceressearch.Cell{Value: "X", Row: 5, Column: 5},
					ceressearch.Cell{Value: "X", Row: 5, Column: 6},
					ceressearch.Cell{Value: "A", Row: 5, Column: 7},
					ceressearch.Cell{Value: "M", Row: 5, Column: 8},
					ceressearch.Cell{Value: "A", Row: 5, Column: 9},
				},
				{
					ceressearch.Cell{Value: "S", Row: 6},
					ceressearch.Cell{Value: "M", Row: 6, Column: 1},
					ceressearch.Cell{Value: "S", Row: 6, Column: 2},
					ceressearch.Cell{Value: "M", Row: 6, Column: 3},
					ceressearch.Cell{Value: "S", Row: 6, Column: 4},
					ceressearch.Cell{Value: "A", Row: 6, Column: 5},
					ceressearch.Cell{Value: "S", Row: 6, Column: 6},
					ceressearch.Cell{Value: "X", Row: 6, Column: 7},
					ceressearch.Cell{Value: "S", Row: 6, Column: 8},
					ceressearch.Cell{Value: "S", Row: 6, Column: 9},
				},
				{
					ceressearch.Cell{Value: "S", Row: 7},
					ceressearch.Cell{Value: "A", Row: 7, Column: 1},
					ceressearch.Cell{Value: "X", Row: 7, Column: 2},
					ceressearch.Cell{Value: "A", Row: 7, Column: 3},
					ceressearch.Cell{Value: "M", Row: 7, Column: 4},
					ceressearch.Cell{Value: "A", Row: 7, Column: 5},
					ceressearch.Cell{Value: "S", Row: 7, Column: 6},
					ceressearch.Cell{Value: "A", Row: 7, Column: 7},
					ceressearch.Cell{Value: "A", Row: 7, Column: 8},
					ceressearch.Cell{Value: "A", Row: 7, Column: 9},
				},
				{
					ceressearch.Cell{Value: "M", Row: 8},
					ceressearch.Cell{Value: "A", Row: 8, Column: 1},
					ceressearch.Cell{Value: "M", Row: 8, Column: 2},
					ceressearch.Cell{Value: "M", Row: 8, Column: 3},
					ceressearch.Cell{Value: "M", Row: 8, Column: 4},
					ceressearch.Cell{Value: "X", Row: 8, Column: 5},
					ceressearch.Cell{Value: "M", Row: 8, Column: 6},
					ceressearch.Cell{Value: "M", Row: 8, Column: 7},
					ceressearch.Cell{Value: "M", Row: 8, Column: 8},
					ceressearch.Cell{Value: "M", Row: 8, Column: 9},
				},
				{
					ceressearch.Cell{Value: "M", Row: 9},
					ceressearch.Cell{Value: "X", Row: 9, Column: 1},
					ceressearch.Cell{Value: "M", Row: 9, Column: 2},
					ceressearch.Cell{Value: "X", Row: 9, Column: 3},
					ceressearch.Cell{Value: "A", Row: 9, Column: 4},
					ceressearch.Cell{Value: "X", Row: 9, Column: 5},
					ceressearch.Cell{Value: "M", Row: 9, Column: 6},
					ceressearch.Cell{Value: "A", Row: 9, Column: 7},
					ceressearch.Cell{Value: "S", Row: 9, Column: 8},
					ceressearch.Cell{Value: "X", Row: 9, Column: 9},
				},
			},
		},
		{
			name: "ok",
			args: args{
				data: []string{
					"AB",
					"CD",
				},
			},
			want: [][]ceressearch.Cell{
				{
					ceressearch.Cell{
						Value:  "A",
						Row:    0,
						Column: 0,
					},
					ceressearch.Cell{
						Value:  "B",
						Row:    0,
						Column: 1,
					},
				},
				{
					ceressearch.Cell{
						Value:  "C",
						Row:    1,
						Column: 0,
					},
					ceressearch.Cell{
						Value:  "D",
						Row:    1,
						Column: 1,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ceressearch.Matrix(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
