package ceressearch_test

import (
	"testing"

	"github.com/gwilczynski/adventofcode/pkg/ceressearch"
)

func TestCall(t *testing.T) {
	type args struct {
		scanX bool
		data  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ok",
			args: args{
				scanX: false,
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
		{
			name: "X",
			args: args{
				scanX: true,
				data: []string{
					".M.S......",
					"..A..MSMS.",
					".M.S.MAA..",
					"..A.ASMSM.",
					".M.S.M....",
					"..........",
					"S.S.S.S.S.",
					".A.A.A.A..",
					"M.M.M.M.M.",
					"..........",
				},
			},
			want: 9,
		},
		{
			name: "XX",
			args: args{
				scanX: true,
				data: []string{
					"..........",
					".M.S......",
					"..A.......",
					".M.S......",
					"..........",
					".S.S......",
					"..A.......",
					".M.M......",
					"..........",
					".S.M......",
					"..A.......",
					".S.M......",
					"..........",
					".M.M......",
					"..A.......",
					".S.S......",
					"..........",
				},
			},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ceressearch.Call(tt.args.data, tt.args.scanX); got != tt.want {
				t.Errorf("Call() = %v, want %v", got, tt.want)
			}
		})
	}
}
