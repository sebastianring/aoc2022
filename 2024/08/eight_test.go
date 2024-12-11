package eight

import (
	"fmt"
	"testing"
)

func TestDayOne(t *testing.T) {
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
			want:    14,
			wantErr: false,
			expected: `......#....#
...#....0...
....#0....#.
..#....0....
....0....#..
.#....A.....
...#........
#......#....
........A...
.........A..
..........#.
..........#.`,
		},
		{
			name: "example3",
			args: args{
				filename: "data_example3.txt",
			},
			want:    4,
			wantErr: false,
			expected: `..........
...#......
#.........
....a.....
........a.
.....a....
..#.......
......#...
..........
..........`,
		},
		{
			name: "real",
			args: args{
				filename: "data.txt",
			},
			want:     -1,
			wantErr:  false,
			expected: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DayEight(tt.args.filename)
			fmt.Println(tt.expected)
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
