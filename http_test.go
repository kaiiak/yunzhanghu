package yunzhanghu

import "testing"

func Test_randomString(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "normal",
			args: args{10},
		},
		{
			name: "zero",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := randomString(tt.args.length)
			t.Logf("randomString() = %v", got)
		})
	}
}
