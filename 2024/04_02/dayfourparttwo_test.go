package dayfourparttwo

import (
	"reflect"
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
		{
			name: "example",
			args: args{
				filename: "data_example.txt",
			},
			want:    9,
			wantErr: false,
		},
		{
			name: "real",
			args: args{
				filename: "data.txt",
			},
			want:    2007, //too high
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DayFourPartTwo(tt.args.filename)
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

func Test_removeAtIndex(t *testing.T) {
	type args struct {
		s     []string
		index int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "happy",
			args: args{
				s: []string{
					"A", "B",
				},
				index: 1,
			},
			want: []string{
				"A",
			},
		},
		{
			name: "happy",
			args: args{
				s: []string{
					"A", "B",
				},
				index: 0,
			},
			want: []string{
				"B",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeAtIndex(tt.args.s, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeAtIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
