package rednosedreport

import (
	"reflect"
	"testing"
)

func TestSafe(t *testing.T) {
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
			if got := Safe(tt.args.data); got != tt.want {
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
		want []string
	}{
		{
			name: "regular",
			args: args{
				data: "7 6 4 2 11 12",
			},
			want: []string{"7", "6", "4", "2", "11", "12"},
		},
		{
			name: "with additional spaces",
			args: args{
				data: "7  6   4    2      11 12     ",
			},
			want: []string{"7", "6", "4", "2", "11", "12"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Split(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Split() = %v, want %v", got, tt.want)
			}
		})
	}
}
