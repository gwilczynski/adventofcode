package scratchcards

import (
	"reflect"
	"testing"
)

func TestGetWorthPoints(t *testing.T) {
	tests := []struct {
		name string
		data []string
		want int
	}{
		{
			name: "ok",
			data: []string{
				"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
				"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
				"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
				"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
				"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
				"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetWorthPoints(tt.data); got != tt.want {
				t.Errorf("GetWorthPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCard(t *testing.T) {
	tests := []struct {
		name string
		data string
		want Card
	}{
		{
			name: "ok",
			data: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			want: Card{
				number:         "1",
				winningNumbers: []int{41, 48, 83, 86, 17},
				numbersYouHave: []int{83, 86, 6, 31, 17, 9, 48, 53},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCard(tt.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCard() = %v, want %v", got, tt.want)
			}
		})
	}
}
