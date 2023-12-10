package gearratios

import "testing"

func TestGetPartNumbers(t *testing.T) {
	tests := []struct {
		name string
		data []string
		want int
	}{
		{
			name: "ok",
			data: []string{
				"467..114..",
				"...*......",
				"..35..633.",
				"......#...",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			},
			want: 4361,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPartNumbers(tt.data); got != tt.want {
				t.Errorf("GetPartNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
