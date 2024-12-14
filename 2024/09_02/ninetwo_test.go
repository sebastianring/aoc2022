package ninetwo

import (
	"fmt"
	"testing"
)

func TestNineOne(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name     string
		args     args
		want     int
		wantErr  bool
		expected string
	}{
		{
			name: "example",
			args: args{
				filename: "data_example.txt",
			},
			want:     2858,
			wantErr:  false,
			expected: "00992111777.44.333....5555.6666.....8888..",
		},
		// {
		// 	name: "real",
		// 	args: args{
		// 		filename: "data.txt",
		// 	},
		// 	want:     6382875730645,
		// 	wantErr:  false,
		// 	expected: "", //8582381894860 too high
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := nineTwo(tt.args.filename)
			fmt.Printf(" ")
			for i := 0; i < len(tt.expected); i++ {
				fmt.Printf("%s ", tt.expected[i:i+1])
			}
			fmt.Println()
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
