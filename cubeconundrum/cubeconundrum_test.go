package cubeconundrum

import (
	"reflect"
	"testing"
)

func Test_buildGame(t *testing.T) {
	tests := []struct {
		name string
		game string
		want Game
	}{
		{
			name: "Game 1",
			game: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want: Game{
				id: 1,
				sets: []Set{
					{
						Blue: 3,
						Red:  4,
					},
					{
						Red:   1,
						Green: 2,
						Blue:  6,
					},
					{
						Green: 2,
					},
				},
			},
		},
		{
			name: "Game 2",
			game: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			want: Game{
				id: 2,
				sets: []Set{
					{
						Blue:  1,
						Green: 2,
					},
					{
						Green: 3,
						Blue:  4,
						Red:   1,
					},
					{
						Green: 1,
						Blue:  1,
					},
				},
			},
		},
		{
			name: "Game 3",
			game: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			want: Game{
				id: 3,
				sets: []Set{
					{
						Green: 8,
						Blue:  6,
						Red:   20,
					},
					{
						Blue:  5,
						Red:   4,
						Green: 13,
					},
					{
						Green: 5,
						Red:   1,
					},
				},
			},
		},
		{
			name: "Game 4",
			game: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			want: Game{
				id: 4,
				sets: []Set{
					{
						Green: 1,
						Red:   3,
						Blue:  6,
					},
					{
						Green: 3,
						Red:   6,
					},
					{
						Green: 3,
						Blue:  15,
						Red:   14,
					},
				},
			},
		},
		{
			name: "Game 5",
			game: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			want: Game{
				id: 5,
				sets: []Set{
					{
						Red:   6,
						Blue:  1,
						Green: 3,
					},
					{
						Blue:  2,
						Red:   1,
						Green: 2,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildGame(tt.game); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseSet(t *testing.T) {
	tests := []struct {
		name       string
		s          string
		wantColor  Color
		wantNumber int
	}{
		{
			name:       "red",
			s:          " 6 red",
			wantColor:  Red,
			wantNumber: 6,
		},
		{
			name:       "blue",
			s:          " 1 blue",
			wantColor:  Blue,
			wantNumber: 1,
		},
		{
			name:       "green",
			s:          " 3 green",
			wantColor:  Green,
			wantNumber: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotColor, gotNumber := parseSet(tt.s)
			if gotColor != tt.wantColor {
				t.Errorf("parseSet() gotColor = %v, want %v", gotColor, tt.wantColor)
			}
			if gotNumber != tt.wantNumber {
				t.Errorf("parseSet() gotNumber = %v, want %v", gotNumber, tt.wantNumber)
			}
		})
	}
}

func Test_parse(t *testing.T) {
	tests := []struct {
		name  string
		games []string
		want  []Game
	}{
		{
			name: "games",
			games: []string{
				"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
				"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
				"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
				"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
				"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			},
			want: []Game{
				{
					id: 1,
					sets: []Set{
						{
							Blue: 3,
							Red:  4,
						},
						{
							Red:   1,
							Green: 2,
							Blue:  6,
						},
						{
							Green: 2,
						},
					},
				},
				{
					id: 2,
					sets: []Set{
						{
							Blue:  1,
							Green: 2,
						},
						{
							Green: 3,
							Blue:  4,
							Red:   1,
						},
						{
							Green: 1,
							Blue:  1,
						},
					},
				},
				{
					id: 3,
					sets: []Set{
						{
							Green: 8,
							Blue:  6,
							Red:   20,
						},
						{
							Blue:  5,
							Red:   4,
							Green: 13,
						},
						{
							Green: 5,
							Red:   1,
						},
					},
				},
				{
					id: 4,
					sets: []Set{
						{
							Green: 1,
							Red:   3,
							Blue:  6,
						},
						{
							Green: 3,
							Red:   6,
						},
						{
							Green: 3,
							Blue:  15,
							Red:   14,
						},
					},
				},
				{
					id: 5,
					sets: []Set{
						{
							Red:   6,
							Blue:  1,
							Green: 3,
						},
						{
							Blue:  2,
							Red:   1,
							Green: 2,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse(tt.games); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_split(t *testing.T) {
	tests := []struct {
		name string
		doc  string
		want []string
	}{
		{
			name: "doc",
			doc:  "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\nGame 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\nGame 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\nGame 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			want: []string{
				"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
				"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
				"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
				"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
				"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := split(tt.doc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("split() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPossible(t *testing.T) {
	bag := Bag{
		Red:   12,
		Green: 13,
		Blue:  14,
	}
	tests := []struct {
		name string
		game Game
		bag  Bag
		want bool
	}{
		{
			name: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			game: Game{
				id: 1,
				sets: []Set{
					{
						Blue: 3,
						Red:  4,
					},
					{
						Red:   1,
						Green: 2,
						Blue:  6,
					},
					{
						Green: 2,
					},
				},
			},
			bag:  bag,
			want: true,
		},
		{
			name: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			game: Game{
				id: 2,
				sets: []Set{
					{
						Blue:  1,
						Green: 2,
					},
					{
						Green: 3,
						Blue:  4,
						Red:   1,
					},
					{
						Green: 1,
						Blue:  1,
					},
				},
			},
			bag:  bag,
			want: true,
		},
		{
			name: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			game: Game{
				id: 3,
				sets: []Set{
					{
						Green: 8,
						Blue:  6,
						Red:   20,
					},
					{
						Blue:  5,
						Red:   4,
						Green: 13,
					},
					{
						Green: 5,
						Red:   1,
					},
				},
			},
			bag:  bag,
			want: false,
		},
		{
			name: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			game: Game{
				id: 4,
				sets: []Set{
					{
						Green: 1,
						Red:   3,
						Blue:  6,
					},
					{
						Green: 3,
						Red:   6,
					},
					{
						Green: 3,
						Blue:  15,
						Red:   14,
					},
				},
			},
			bag:  bag,
			want: false,
		},
		{
			name: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			game: Game{
				id: 5,
				sets: []Set{
					{
						Red:   6,
						Blue:  1,
						Green: 3,
					},
					{
						Blue:  2,
						Red:   1,
						Green: 2,
					},
				},
			},
			bag:  bag,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPossible(tt.game, tt.bag); got != tt.want {
				t.Errorf("ispossible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWhichGamesPossible(t *testing.T) {
	tests := []struct {
		name string
		doc  string
		bag  Bag
		want int
	}{
		{
			name: "game",
			want: 8,
			doc:  "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\nGame 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\nGame 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\nGame 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			bag: Bag{
				Red:   12,
				Green: 13,
				Blue:  14,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WhichGamesPossible(tt.doc, tt.bag); got != tt.want {
				t.Errorf("WhichGamesPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
