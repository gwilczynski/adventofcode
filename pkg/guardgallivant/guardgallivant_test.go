package guardgallivant_test

import (
	"reflect"
	"testing"

	"github.com/gwilczynski/adventofcode/pkg/guardgallivant"
)

func TestDecodeLine(t *testing.T) {
	type args struct {
		line string
		j    int
	}
	tests := []struct {
		name   string
		args   args
		fields []*guardgallivant.Field
		guard  *guardgallivant.Guard
	}{
		{
			name: "ok",
			args: args{
				line: "....#.....",
				j:    0,
			},
			fields: []*guardgallivant.Field{
				{
					T: guardgallivant.Free,
				},
				{
					T: guardgallivant.Free,
				},
				{
					T: guardgallivant.Free,
				},
				{
					T: guardgallivant.Free,
				},
				{
					T: guardgallivant.Obstruction,
				},
				{
					T: guardgallivant.Free,
				},
				{
					T: guardgallivant.Free,
				},
				{
					T: guardgallivant.Free,
				},
				{
					T: guardgallivant.Free,
				},
				{
					T: guardgallivant.Free,
				},
			},
			guard: nil,
		},
		{
			name: "ok",
			args: args{
				line: ".#..^.....",
				j:    6,
			},
			fields: []*guardgallivant.Field{
				{
					T: guardgallivant.Free,
				},
				{
					T: guardgallivant.Obstruction,
				},
				{
					T: guardgallivant.Free,
				},
				{
					T: guardgallivant.Free,
				},
				{
					T: guardgallivant.Free,
				},
				{
					T: guardgallivant.Free,
				},
				{
					T: guardgallivant.Free,
				},
				{
					T: guardgallivant.Free,
				},
				{
					T: guardgallivant.Free,
				},
				{
					T: guardgallivant.Free,
				},
			},
			guard: &guardgallivant.Guard{
				D: guardgallivant.Up,
				P: guardgallivant.Position{
					I: 4,
					J: 6,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields, guard := guardgallivant.DecodeLine(tt.args.line, tt.args.j)
			if !reflect.DeepEqual(fields, tt.fields) {
				t.Errorf("DecodeLine() fields = %v, want %v", fields, tt.fields)
			}
			if !reflect.DeepEqual(guard, tt.guard) {
				t.Errorf("DecodeLine() guard = %v, want %v", guard, tt.guard)
			}
		})
	}
}

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
					"....#.....",
					".........#",
					"..........",
					"..#.......",
					".......#..",
					"..........",
					".#..^.....",
					"........#.",
					"#.........",
					"......#...",
				},
			},
			want: 41,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := guardgallivant.Call(tt.args.data); got != tt.want {
				t.Errorf("Call() = %v, want %v", got, tt.want)
			}
		})
	}
}
