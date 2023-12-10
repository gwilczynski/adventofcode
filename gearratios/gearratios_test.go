package gearratios

import (
	"reflect"
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

func Test_scanItems(t *testing.T) {
	tests := []struct {
		name      string
		index     int
		line      string
		wantItems []Item
	}{
		{
			name:  "first line",
			index: 0,
			line:  "467..114..",

			wantItems: []Item{
				{
					line:       0,
					startIndex: 0,
					endIndex:   2,
					itemType:   Number,
					value:      467,
				},
				{
					line:       0,
					startIndex: 5,
					endIndex:   7,
					itemType:   Number,
					value:      114,
				},
			},
		},
		{
			name:  "second line",
			index: 1,
			line:  "...*......",

			wantItems: []Item{
				{
					line:       1,
					startIndex: 3,
					endIndex:   3,
					value:      0,
					itemType:   Special,
				},
			},
		},
		{
			name:  "third line",
			index: 2,
			line:  "....++.58.",

			wantItems: []Item{
				{
					line:       2,
					startIndex: 4,
					endIndex:   4,
					value:      0,
					itemType:   Special,
				},
				{
					line:       2,
					startIndex: 5,
					endIndex:   5,
					value:      0,
					itemType:   Special,
				},
				{
					line:       2,
					startIndex: 7,
					endIndex:   8,
					value:      58,
					itemType:   Number,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotItems := scanItems(tt.index, tt.line); !reflect.DeepEqual(gotItems, tt.wantItems) {
				t.Errorf("scanItems() = %v, want %v", gotItems, tt.wantItems)
			}
		})
	}
}
