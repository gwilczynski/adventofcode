package rednosedreport_test

import (
	"reflect"
	"testing"

	"github.com/gwilczynski/adventofcode/pkg/rednosedreport"
)

func TestHowManyReportsAreSafe(t *testing.T) {
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
				data: []string{"7 6 4 2 1", "1 2 7 8 9", "9 7 6 2 1", "1 3 2 4 5", "8 6 4 4 1", "1 3 6 7 9"},
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rednosedreport.HowManyReportsAreSafe(tt.args.data); got != tt.want {
				t.Errorf("Safe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "regular",
			args: args{
				data: "7 6 4 2 11 12",
			},
			want: []int{7, 6, 4, 2, 11, 12},
		},
		{
			name: "with additional spaces",
			args: args{
				data: "7  6   4    2      11 12     ",
			},
			want: []int{7, 6, 4, 2, 11, 12},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rednosedreport.Split(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Split() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafe(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "safe because the levels are all decreasing by 1 or 2",
			args: args{
				data: []int{7, 6, 4, 2, 1},
			},
			want: true,
		},
		{
			name: "unsafe because 2 7 is an increase of 5",
			args: args{
				data: []int{1, 2, 7, 8, 9},
			},
			want: false,
		},
		{
			name: "unsafe because 6 2 is a decrease of 4",
			args: args{
				data: []int{9, 7, 6, 2, 1},
			},
			want: false,
		},
		{
			name: "unsafe because 1 3 is increasing but 3 2 is decreasing",
			args: args{
				data: []int{1, 3, 2, 4, 5},
			},
			want: false,
		},
		{
			name: "unsafe because 4 4 is neither an increase or a decrease",
			args: args{
				data: []int{8, 6, 4, 4, 1},
			},
			want: false,
		},
		{
			name: "safe because the levels are all increasing by 1, 2, or 3",
			args: args{
				data: []int{1, 3, 6, 7, 9},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rednosedreport.Safe(tt.args.data); got != tt.want {
				t.Errorf("Safe() = %v, want %v", got, tt.want)
			}
		})
	}
}
