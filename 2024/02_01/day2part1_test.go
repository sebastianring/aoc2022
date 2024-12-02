package daytwopartone

import "testing"

func TestDay2Part1(t *testing.T) {
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
			want:    2,
			wantErr: false,
		},
		{
			name: "real",
			args: args{
				filename: "data.txt",
			},
			want:    2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Day2Part1(tt.args.filename)
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
