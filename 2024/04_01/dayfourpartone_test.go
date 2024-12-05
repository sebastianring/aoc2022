package dayfourpartone

import "testing"

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
			want:    18,
			wantErr: false,
		},
		{
			name: "real",
			args: args{
				filename: "data.txt",
			},
			want:    2603,
			wantErr: false,
		},
		// {
		// 	name: "example2",
		// 	args: args{
		// 		filename: "data_example2.txt",
		// 	},
		// 	want:    0,
		// 	wantErr: false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DayFourPartOne(tt.args.filename)
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
