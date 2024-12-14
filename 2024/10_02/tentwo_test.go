package tentwo

import "testing"

func TestTenOne(t *testing.T) {
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
			want:    3,
			wantErr: false,
		},
		{
			name: "example 2",
			args: args{
				filename: "data_example2.txt",
			},
			want:    13,
			wantErr: false,
		},
		{
			name: "example 3",
			args: args{
				filename: "data_example3.txt",
			},
			want:    227,
			wantErr: false,
		},
		{
			name: "example 4",
			args: args{
				filename: "data_example4.txt",
			},
			want:    81,
			wantErr: false,
		},
		{
			name: "real",
			args: args{
				filename: "data.txt",
			},
			want:    966,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TenTwo(tt.args.filename)
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
