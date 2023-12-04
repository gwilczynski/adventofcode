package trebuchet

import (
	"reflect"
	"testing"
)

func Test_Decode(t *testing.T) {
	tests := []struct {
		name   string
		value  string
		result int
	}{
		{
			name: "example",
			value: "1abc2\n" +
				"pqr3stu8vwx\n" +
				"a1b2c3d4e5f\n" +
				"treb7uchet",
			result: 142,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Decode(tt.value)
			if result != tt.result {
				t.Errorf("got %d, want %d", result, tt.result)
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
			name: "ok",
			doc: "1abc2\n" +
				"pqr3stu8vwx\n" +
				"a1b2c3d4e5f\n" +
				"treb7uchet",
			want: []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"},
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

func Test_calibration(t *testing.T) {
	tests := []struct {
		name  string
		line  string
		first string
		last  string
	}{
		{
			name:  "ok - 1",
			line:  "1abc2",
			first: "1",
			last:  "2",
		},
		{
			name:  "ok - 2",
			line:  "pqr3stu8vwx",
			first: "3",
			last:  "8",
		},
		{
			name:  "ok - 3",
			line:  "a1b2c3d4e5f",
			first: "1",
			last:  "5",
		},
		{
			name:  "ok - 4",
			line:  "treb7uchet",
			first: "7",
			last:  "7",
		},
		{
			name:  "ok - 5",
			line:  "357",
			first: "3",
			last:  "7",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := calibration(tt.line)
			if got != tt.first {
				t.Errorf("calibrationValue() got = %v, want %v", got, tt.first)
			}
			if got1 != tt.last {
				t.Errorf("calibrationValue() got1 = %v, want %v", got1, tt.last)
			}
		})
	}
}

func Test_number(t *testing.T) {
	tests := []struct {
		name    string
		first   string
		second  string
		want    int
		wantErr bool
	}{
		{
			name:   "ok - 1",
			first:  "1",
			second: "2",
			want:   12,
		},
		{
			name:   "ok - 2",
			first:  "3",
			second: "8",
			want:   38,
		},
		{
			name:   "ok - 3",
			first:  "1",
			second: "5",
			want:   15,
		},
		{
			name:   "ok - 4",
			first:  "7",
			second: "7",
			want:   77,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := number(tt.first, tt.second)
			if (err != nil) != tt.wantErr {
				t.Errorf("number() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("number() got = %v, want %v", got, tt.want)
			}
		})
	}
}
