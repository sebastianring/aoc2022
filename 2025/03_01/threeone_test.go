package threeone

import (
	"fmt"
	"testing"
)

func TestDayThreeOne(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "example",
			args: args{
				filename: "data_example.txt",
			},
			want:    357,
			wantErr: false,
		},
		{
			name: "real data",
			args: args{
				filename: "data.txt",
			},
			want:    27180728081,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ThreeOne(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ThreeOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				fmt.Println(tt.want - got)
				t.Errorf("ThreeOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
