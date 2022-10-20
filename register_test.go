package yunzhanghu_test

import (
	"context"
	"fmt"
	"github.com/kaiiak/yunzhanghu"
	"testing"
)

func TestRegister(t *testing.T) {
	type args struct {
		ctx      context.Context
		DealerUserId   string
		PhoneNo   string
		IdCard string
		RealName   string
		IdCardAddress string
		IdCardAgency string
		IdCardNation string
		IdCardValidityStart string
		IdCardValidityEnd string

	}
	tests := []struct {
		name    string
		client  *yunzhanghu.Yunzhanghu
		args    args
		wantErr bool
	}{
		{
			name:   "ok",
			client: &yunzhanghu.Yunzhanghu{},
			args: args{
				ctx: context.Background(),
				// TODO
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dealerId,err := tt.client.Register(tt.args.ctx, tt.args.DealerUserId, tt.args.PhoneNo, tt.args.IdCard, tt.args.RealName,tt.args.IdCardAddress,tt.args.IdCardAgency,tt.args.IdCardNation,tt.args.IdCardValidityStart,tt.args.IdCardValidityEnd)
			if  (err != nil) != tt.wantErr {
				t.Errorf("Yunzhanghu.VerifyBankcardFourFactor() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Println("dealerId:",dealerId)
		})
	}
}