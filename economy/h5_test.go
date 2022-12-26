package economy

import (
	"context"
	"reflect"
	"testing"

	"github.com/kaiiak/yunzhanghu/core"
)

func TestEconomy_GetEconomyStatusWithIdCard(t *testing.T) {
	type fields struct {
		Config *core.Config
	}
	type args struct {
		ctx          context.Context
		readName     string
		idCard       string
		dealerUserId string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantStatus EconomyStatus
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Economy{
				Config: tt.fields.Config,
			}
			gotStatus, err := e.GetEconomyStatusWithIdCard(tt.args.ctx, tt.args.readName, tt.args.idCard, tt.args.dealerUserId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Economy.GetEconomyStatusWithIdCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStatus, tt.wantStatus) {
				t.Errorf("Economy.GetEconomyStatusWithIdCard() = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}
