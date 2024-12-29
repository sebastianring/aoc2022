package twelvetwo

import "testing"

func TestTwelveTwo(t *testing.T) {
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
			want:    80,
			wantErr: false,
		},
		{
			name: "example2",
			args: args{
				filename: "data_example2.txt",
			},
			want:    436,
			wantErr: false,
		},
		{
			name: "example3",
			args: args{
				filename: "data_example3.txt",
			},
			want:    368,
			wantErr: false,
		},
		{
			name: "example4",
			args: args{
				filename: "data_example4.txt",
			},
			want:    1206,
			wantErr: false,
		},
		{
			name: "example5",
			args: args{
				filename: "data_example5.txt",
			},
			want:    160,
			wantErr: false,
		},
		{
			name: "example6",
			args: args{
				filename: "data_example6.txt",
			},
			want:    192,
			wantErr: false,
		},
		{
			name: "example7",
			args: args{
				filename: "data_example7.txt",
			},
			want:    946,
			wantErr: false,
		},
		{
			name: "real",
			args: args{
				filename: "data.txt",
			},
			want:    811148, // 801607 too low // 801856 too low // 802614 too low // 818743 .. wrong
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TwelveTwo(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("Twelve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Twelve() = %v, want %v", got, tt.want)
			}
		})
	}
}
