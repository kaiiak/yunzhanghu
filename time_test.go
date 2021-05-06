package yunzhanghu_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/kaiiak/yunzhanghu"
)

func TestCustomTime_UnmarshalJson(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				data: []byte("2017-10-16 20:58:29"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := yunzhanghu.Time{}
			if err := tr.UnmarshalJson(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Time.UnmarshalJson() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				t.Log(tr.Time)
			}
		})
	}
}

func TestTime_MarshalJSON(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Time: time.Date(2017, time.October, 16, 20, 58, 29, 0, yunzhanghu.ShangHaiTimeLocation),
			},
			want: []byte("\"2017-10-16 20:58:29\""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := yunzhanghu.Time{
				Time: tt.fields.Time,
			}
			got, err := tr.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Time.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Time.MarshalJSON() = %s, want %s", got, tt.want)
			}
		})
	}
}
