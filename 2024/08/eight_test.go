package eight

import (
	"fmt"
	"testing"
)

func TestDayEightPartOne(t *testing.T) {
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
			want:     259,
			wantErr:  false,
			expected: ``,
		},
		{
			name: "example4",
			args: args{
				filename: "data_example4.txt",
			},
			want:    259,
			wantErr: false,
			expected: `T....#....
...T......
.T....#...
.........#
..#.......
..........
...#......
..........
....#.....
..........`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DayEightPartOne(tt.args.filename)
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

func TestDayEightPartTwo(t *testing.T) {
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
			name: "example4",
			args: args{
				filename: "data_example4.txt",
			},
			want:    9,
			wantErr: false,
			expected: `T....#....
...T......
.T....#...
.........#
..#.......
..........
...#......
..........
....#.....
..........`,
		},
		{
			name: "example",
			args: args{
				filename: "data_example.txt",
			},
			want:    34,
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
			name: "real",
			args: args{
				filename: "data.txt",
			},
			want:     34,
			wantErr:  false,
			expected: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DayEightPartTwo(tt.args.filename)
			fmt.Println(tt.expected)
			if (err != nil) != tt.wantErr {
				t.Errorf("DayEightPartTwo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DayEightPartTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
