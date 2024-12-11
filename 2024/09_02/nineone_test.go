package nineone

import (
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
			want:     1928,
			wantErr:  false,
			expected: "0099811188827773336446555566..............",
		},
		{
			name: "real",
			args: args{
				filename: "data.txt",
			},
			want:     6382875730645,
			wantErr:  false,
			expected: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := nineOne(tt.args.filename)
			// fmt.Printf(" ")
			// for i := 0; i < len(tt.expected); i++ {
			// 	fmt.Printf("%s ", tt.expected[i:i+1])
			// }
			// fmt.Println()
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
