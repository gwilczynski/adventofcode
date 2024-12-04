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
