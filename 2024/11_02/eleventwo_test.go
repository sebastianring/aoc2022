package eleventwo

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
		// {
		// 	name: "example",
		// 	args: args{
		// 		filename: "data_example.txt",
		// 	},
		// 	want:    -1,
		// 	wantErr: false,
		// },
		// {
		// 	name: "example2",
		// 	args: args{
		// 		filename: "data_example2.txt",
		// 	},
		// 	want:    55312,
		// 	wantErr: false,
		// },
		// {
		// 	name: "example3",
		// 	args: args{
		// 		filename: "data_example3.txt",
		// 	},
		// 	want:    55312,
		// 	wantErr: false,
		// },
		{
			name: "real",
			args: args{
				filename: "data.txt",
			},
			want:    55312,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := eleventwo(tt.args.filename)
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

func Test_split(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{
				s: "1010",
			},
			want: []string{"10", "10"},
		},
		{
			name: "test1",
			args: args{
				s: "2222",
			},
			want: []string{"22", "22"},
		},
		{
			name: "test1",
			args: args{
				s: "2200",
			},
			want: []string{"22", "0"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := split(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("split() = %v, want %v", got, tt.want)
			}
		})
	}
}
