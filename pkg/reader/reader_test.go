package reader_test

import (
	"reflect"
	"testing"

	"github.com/gwilczynski/adventofcode/pkg/reader"
)

func TestReadData(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				filePath: "./data_test.txt",
			},
			want:    []string{"3   4", "4   3", "2   5", "1   3", "3   9", "3   3"},
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				filePath: "./does_not_exists.txt",
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := reader.ReadData(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadData() got = %v, want %v", got, tt.want)
			}
		})
	}
}
