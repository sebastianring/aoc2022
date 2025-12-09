package onetwo

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
			want:    6,
			wantErr: false,
		},
		{
			name: "real data",
			args: args{
				filename: "data.txt",
			},
			want:    6364, // too large, 5805 too small
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OneTwo(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("OneTwo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("OneTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
