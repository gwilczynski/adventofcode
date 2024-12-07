package bridgerepair

import (
	"reflect"
	"testing"
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
					"190: 10 19",
					"3267: 81 40 27",
					"83: 17 5",
					"156: 15 6",
					"7290: 6 8 6 15",
					"161011: 16 10 13",
					"192: 17 8 14",
					"21037: 9 7 18 13",
					"292: 11 6 16 20",
				},
			},
			want: 3749,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Call(tt.args.data); got != tt.want {
				t.Errorf("Call() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want *Puzzle
	}{
		{
			name: "ok",
			args: args{
				line: "156: 15 6",
			},
			want: &Puzzle{
				Result:  156,
				Numbers: []int{15, 6},
			},
		},
		{
			name: "ok",
			args: args{
				line: "7290: 6 8 6 15",
			},
			want: &Puzzle{
				Result:  7290,
				Numbers: []int{6, 8, 6, 15},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseLine(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
		want []*Puzzle
	}{
		{
			name: "ok",
			args: args{
				data: []string{
					"190: 10 19",
					"3267: 81 40 27",
					"83: 17 5",
				},
			},
			want: []*Puzzle{
				{
					Result:  190,
					Numbers: []int{10, 19},
				},
				{
					Result:  3267,
					Numbers: []int{81, 40, 27},
				},
				{
					Result:  83,
					Numbers: []int{17, 5},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
