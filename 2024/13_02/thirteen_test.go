package thirteen

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
			want:    480,
			wantErr: false,
		},
		{
			name: "real data",
			args: args{
				filename: "data.txt",
			},
			// 35887 too low
			want:    35886,
			wantErr: false,
		},
		{
			name: "3",
			args: args{
				filename: "data_3.txt",
			},
			// 35887 too low
			want:    35886,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Thirteen(tt.args.filename)
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
