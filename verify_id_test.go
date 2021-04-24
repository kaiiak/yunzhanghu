package yunzhanghu_test

import (
	"context"
	"testing"

	"github.com/kaiiak/yunzhanghu"
)

func TestYunzhanghu_UserCardCheck(t *testing.T) {
	type args struct {
		ctx      context.Context
		realName string
		idCard   string
	}
	tests := []struct {
		name    string
		client  *yunzhanghu.Yunzhanghu
		args    args
		wantErr bool
	}{
		{
			name:   "ok",
			client: nil,
			args: args{
				ctx: context.Background(),
				// TODO
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.client.VerifyId(tt.args.ctx, tt.args.realName, tt.args.idCard); (err != nil) != tt.wantErr {
				t.Errorf("Yunzhanghu.UserCardCheck() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
