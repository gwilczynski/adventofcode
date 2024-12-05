package printqueue

import "testing"

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
					"47|53",
					"97|13",
					"97|61",
					"97|47",
					"75|29",
					"61|13",
					"75|53",
					"29|13",
					"97|29",
					"53|29",
					"61|53",
					"97|53",
					"61|29",
					"47|13",
					"75|47",
					"97|75",
					"47|61",
					"75|61",
					"47|29",
					"75|13",
					"53|13",
					"",
					"75,47,61,53,29",
					"97,61,53,29,13",
					"75,29,13",
					"75,97,47,61,53",
					"61,13,29",
					"97,13,75,29,47",
				},
			},
			want: 143,
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
