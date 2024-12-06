package printqueue

import (
	"reflect"
	"testing"
)

func TestCall(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name      string
		args      args
		want      int
		wantFixed int
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
			want:      143,
			wantFixed: 123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if correct, _ := Call(tt.args.data); correct != tt.want {
				t.Errorf("Call() = %v, want %v", correct, tt.want)
			}

			if _, fixed := Call(tt.args.data); fixed != tt.wantFixed {
				t.Errorf("Call() = %v, want %v", fixed, tt.wantFixed)
			}
		})
	}
}

func TestRulesAndUpdates(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name        string
		args        args
		wantRules   []Rule
		wantUpdates [][]int
	}{
		{
			name: "ok",
			args: args{
				data: []string{
					"47|53",
					"97|13",
					"",
					"75,47,61,53,29",
					"75,29,13",
				},
			},
			wantRules: []Rule{
				{
					before: 47,
					after:  53,
				},
				{
					before: 97,
					after:  13,
				},
			},
			wantUpdates: [][]int{
				{
					75, 47, 61, 53, 29,
				},
				{
					75, 29, 13,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rules, updates := RulesAndUpdates(tt.args.data)
			if !reflect.DeepEqual(rules, tt.wantRules) {
				t.Errorf("RulesAndUpdates() got = %v, want %v", rules, tt.wantRules)
			}
			if !reflect.DeepEqual(updates, tt.wantUpdates) {
				t.Errorf("RulesAndUpdates() got1 = %v, want %v", updates, tt.wantUpdates)
			}
		})
	}
}

func TestValid(t *testing.T) {
	input := []string{
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
	}
	rules, updates := RulesAndUpdates(input)

	type args struct {
		update []int
		rules  []Rule
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ok",
			args: args{
				update: updates[0],
				rules:  rules,
			},
			want: true,
		},
		{
			name: "ok",
			args: args{
				update: updates[1],
				rules:  rules,
			},
			want: true,
		},
		{
			name: "ok",
			args: args{
				update: updates[2],
				rules:  rules,
			},
			want: true,
		},
		{
			name: "not ok",
			args: args{
				update: updates[3],
				rules:  rules,
			},
			want: false,
		},
		{
			name: "not ok",
			args: args{
				update: updates[4],
				rules:  rules,
			},
			want: false,
		},
		{
			name: "not ok",
			args: args{
				update: updates[5],
				rules:  rules,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Valid(tt.args.update, tt.args.rules); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiddle(t *testing.T) {
	type args struct {
		update []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ok",
			args: args{
				update: []int{97, 61, 53, 29, 13, 75, 29},
			},
			want: 29,
		},
		{
			name: "ok",
			args: args{
				update: []int{97, 61, 53, 29, 13},
			},
			want: 53,
		},
		{
			name: "ok",
			args: args{
				update: []int{75, 29, 13},
			},
			want: 29,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Middle(tt.args.update); got != tt.want {
				t.Errorf("Middle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSort(t *testing.T) {
	input := []string{
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
	}
	rules, _ := RulesAndUpdates(input)

	type args struct {
		update []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "ok",
			args: args{
				update: []int{75, 97, 47, 61, 53},
			},
			want: []int{97, 75, 47, 61, 53},
		},
		{
			name: "ok",
			args: args{
				update: []int{61, 13, 29},
			},
			want: []int{61, 29, 13},
		},
		{
			name: "ok",
			args: args{
				update: []int{97, 13, 75, 29, 47},
			},
			want: []int{97, 75, 47, 29, 13},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort(tt.args.update, rules); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
