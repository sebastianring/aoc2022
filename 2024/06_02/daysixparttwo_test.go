package daysixparttwo

import (
	"fmt"
	"testing"
)

func TestDayOne(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// {
		// 	name: "example",
		// 	args: args{
		// 		filename: "data_example.txt",
		// 	},
		// 	want:    6,
		// 	wantErr: false,
		// },
		// {
		// 	name: "example2",
		// 	args: args{
		// 		filename: "data_example2.txt",
		// 	},
		// 	want:    1,
		// 	wantErr: false,
		// },
		// {
		// 	name: "example3",
		// 	args: args{
		// 		filename: "data_example3.txt",
		// 	},
		// 	want:    4,
		// 	wantErr: false,
		// },
		{
			name: "real",
			args: args{
				filename: "data.txt",
			},
			want:    917, //too low // 1998 too high // 1895 too high
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("--")
			got, err := DayOne(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("DayOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DayOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
