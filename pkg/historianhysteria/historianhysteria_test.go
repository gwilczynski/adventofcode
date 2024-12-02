package historianhysteria_test

import (
	"reflect"
	"testing"

	"github.com/gwilczynski/adventofcode/pkg/historianhysteria"
)

func TestTotalDistance(t *testing.T) {
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
				data: []string{"3   4", "4   3", "2   5", "1   3", "3   9", "3   3"},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := historianhysteria.TotalDistance(tt.args.data); got != tt.want {
				t.Errorf("TotalDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitLines(t *testing.T) {
	type args struct {
		numberOfColumns int
		lines           []string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "ok",
			args: args{
				numberOfColumns: 2,
				lines:           []string{"3   4", "4   3", "2   5", "1   3", "3   9", "3   3"},
			},
			want: [][]int{{3, 4, 2, 1, 3, 3}, {4, 3, 5, 3, 9, 3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := historianhysteria.SplitLines(tt.args.lines, tt.args.numberOfColumns)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitLines() got = %v, want %v", got, tt.want)
			}
		})
	}
}
