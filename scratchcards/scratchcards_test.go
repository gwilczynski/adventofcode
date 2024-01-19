package scratchcards

import "testing"

func TestGetWorthPoints(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "ok",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetWorthPoints(); got != tt.want {
				t.Errorf("GetWorthPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
