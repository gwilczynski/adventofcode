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
					ceressearch.Cell{"M", 0, 0},
					ceressearch.Cell{"M", 0, 1},
					ceressearch.Cell{"M", 0, 2},
					ceressearch.Cell{"S", 0, 3},
					ceressearch.Cell{"X", 0, 4},
					ceressearch.Cell{"X", 0, 5},
					ceressearch.Cell{"M", 0, 6},
					ceressearch.Cell{"A", 0, 7},
					ceressearch.Cell{"S", 0, 8},
					ceressearch.Cell{"M", 0, 9},
				},
				{
					ceressearch.Cell{"M", 1, 0},
					ceressearch.Cell{"S", 1, 1},
					ceressearch.Cell{"A", 1, 2},
					ceressearch.Cell{"M", 1, 3},
					ceressearch.Cell{"X", 1, 4},
					ceressearch.Cell{"M", 1, 5},
					ceressearch.Cell{"S", 1, 6},
					ceressearch.Cell{"M", 1, 7},
					ceressearch.Cell{"S", 1, 8},
					ceressearch.Cell{"A", 1, 9},
				},
				{
					ceressearch.Cell{"A", 2, 0},
					ceressearch.Cell{"M", 2, 1},
					ceressearch.Cell{"X", 2, 2},
					ceressearch.Cell{"S", 2, 3},
					ceressearch.Cell{"X", 2, 4},
					ceressearch.Cell{"M", 2, 5},
					ceressearch.Cell{"A", 2, 6},
					ceressearch.Cell{"A", 2, 7},
					ceressearch.Cell{"M", 2, 8},
					ceressearch.Cell{"M", 2, 9},
				},
				{
					ceressearch.Cell{"M", 3, 0},
					ceressearch.Cell{"S", 3, 1},
					ceressearch.Cell{"A", 3, 2},
					ceressearch.Cell{"M", 3, 3},
					ceressearch.Cell{"A", 3, 4},
					ceressearch.Cell{"S", 3, 5},
					ceressearch.Cell{"M", 3, 6},
					ceressearch.Cell{"S", 3, 7},
					ceressearch.Cell{"M", 3, 8},
					ceressearch.Cell{"X", 3, 9},
				},
				{
					ceressearch.Cell{"X", 4, 0},
					ceressearch.Cell{"M", 4, 1},
					ceressearch.Cell{"A", 4, 2},
					ceressearch.Cell{"S", 4, 3},
					ceressearch.Cell{"A", 4, 4},
					ceressearch.Cell{"M", 4, 5},
					ceressearch.Cell{"X", 4, 6},
					ceressearch.Cell{"A", 4, 7},
					ceressearch.Cell{"M", 4, 8},
					ceressearch.Cell{"M", 4, 9},
				},
				{
					ceressearch.Cell{"X", 5, 0},
					ceressearch.Cell{"X", 5, 1},
					ceressearch.Cell{"A", 5, 2},
					ceressearch.Cell{"M", 5, 3},
					ceressearch.Cell{"M", 5, 4},
					ceressearch.Cell{"X", 5, 5},
					ceressearch.Cell{"X", 5, 6},
					ceressearch.Cell{"A", 5, 7},
					ceressearch.Cell{"M", 5, 8},
					ceressearch.Cell{"A", 5, 9},
				},
				{
					ceressearch.Cell{"S", 6, 0},
					ceressearch.Cell{"M", 6, 1},
					ceressearch.Cell{"S", 6, 2},
					ceressearch.Cell{"M", 6, 3},
					ceressearch.Cell{"S", 6, 4},
					ceressearch.Cell{"A", 6, 5},
					ceressearch.Cell{"S", 6, 6},
					ceressearch.Cell{"X", 6, 7},
					ceressearch.Cell{"S", 6, 8},
					ceressearch.Cell{"S", 6, 9},
				},
				{
					ceressearch.Cell{"S", 7, 0},
					ceressearch.Cell{"A", 7, 1},
					ceressearch.Cell{"X", 7, 2},
					ceressearch.Cell{"A", 7, 3},
					ceressearch.Cell{"M", 7, 4},
					ceressearch.Cell{"A", 7, 5},
					ceressearch.Cell{"S", 7, 6},
					ceressearch.Cell{"A", 7, 7},
					ceressearch.Cell{"A", 7, 8},
					ceressearch.Cell{"A", 7, 9},
				},
				{
					ceressearch.Cell{"M", 8, 0},
					ceressearch.Cell{"A", 8, 1},
					ceressearch.Cell{"M", 8, 2},
					ceressearch.Cell{"M", 8, 3},
					ceressearch.Cell{"M", 8, 4},
					ceressearch.Cell{"X", 8, 5},
					ceressearch.Cell{"M", 8, 6},
					ceressearch.Cell{"M", 8, 7},
					ceressearch.Cell{"M", 8, 8},
					ceressearch.Cell{"M", 8, 9},
				},
				{
					ceressearch.Cell{"M", 9, 0},
					ceressearch.Cell{"X", 9, 1},
					ceressearch.Cell{"M", 9, 2},
					ceressearch.Cell{"X", 9, 3},
					ceressearch.Cell{"A", 9, 4},
					ceressearch.Cell{"X", 9, 5},
					ceressearch.Cell{"M", 9, 6},
					ceressearch.Cell{"A", 9, 7},
					ceressearch.Cell{"S", 9, 8},
					ceressearch.Cell{"X", 9, 9},
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
