package fivetwo

import (
	"testing"
)

func TestSolution(t *testing.T) {
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
			want:    3121910778619,
			wantErr: false,
		},
		{
			name: "real data",
			args: args{
				filename: "data.txt",
			},
			want: 354149806372819, // too low

			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Solution(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("Solution() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
