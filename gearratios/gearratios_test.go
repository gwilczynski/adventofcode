package gearratios

import (
	"testing"
)

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

func Test_scanItems1(t *testing.T) {
	tests := []struct {
		name     string
		index    int
		line     string
		numbers  *[]NumberItem
		specials *[]SpecialItem
	}{
		{
			name:  "first line",
			index: 0,
			line:  "467..114..",
			numbers: &[]NumberItem{
				{
					line:       0,
					startIndex: 0,
					endIndex:   2,
					value:      467,
				},
				{
					line:       0,
					startIndex: 5,
					endIndex:   7,
					value:      114,
				},
			},
			specials: nil,
		},
		{
			name:  "second line",
			index: 1,
			line:  "...*......",

			numbers: nil,
			specials: &[]SpecialItem{
				{
					line:       1,
					startIndex: 3,
					endIndex:   3,
					value:      "*",
				},
			},
		},
		{
			name:  "third line",
			index: 2,
			line:  "....++.58.",

			numbers: &[]NumberItem{
				{
					line:       2,
					startIndex: 7,
					endIndex:   8,
					value:      58,
				},
			},
			specials: &[]SpecialItem{
				{
					line:       2,
					startIndex: 4,
					endIndex:   4,
					value:      "+",
				},
				{
					line:       2,
					startIndex: 5,
					endIndex:   5,
					value:      "+",
				},
			},
		},
		{
			name:  "fourth line",
			index: 3,
			line:  "617*......",

			numbers: &[]NumberItem{
				{
					line:       3,
					startIndex: 0,
					endIndex:   2,
					value:      617,
				},
			},
			specials: &[]SpecialItem{
				{
					line:       2,
					startIndex: 3,
					endIndex:   3,
					value:      "+",
				},
			},
		},
		{
			name:  "empty line",
			index: 4,
			line:  "..........",

			numbers:  nil,
			specials: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanItems(tt.index, tt.line, tt.numbers, tt.specials)
		})
	}
}
