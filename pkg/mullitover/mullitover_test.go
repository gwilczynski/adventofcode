package mullitover

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
			name: "only the four highlighted sections are real mul instructions",
			args: args{
				data: []string{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"},
			},
			want: 161,
		},
		{
			name: "only the most recent do() or don't() instruction applies. at the beginning of the program, mul instructions are enabled",
			args: args{
				data: []string{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"},
			},
			want: 48,
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

func TestExtractInstructions(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "only the four highlighted sections are real mul instructions",
			args: args{
				input: "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			},
			want: []string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractInstructions(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractInstructions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractNumbers(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want Num
	}{
		{
			name: "mul(2,4)",
			args: args{
				input: "mul(2,4)",
			},
			want: Num{2, 4},
		},
		{
			name: "mul(11,3)",
			args: args{
				input: "mul(11,3)",
			},
			want: Num{11, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtractNumbers(tt.args.input)
			if got != tt.want {
				t.Errorf("ExtractNumbers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractInstructionsCombined(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "only the most recent do() or don't() instruction applies. at the beginning of the program, mul instructions are enabled",
			args: args{
				input: "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			},
			want: []string{"mul(2,4)", "don't()", "mul(5,5)", "mul(11,8)", "do()", "mul(8,5)"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractInstructionsCombined(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractInstructionsCombined() = %v, want %v", got, tt.want)
			}
		})
	}
}
