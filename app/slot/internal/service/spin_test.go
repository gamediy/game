package service

import (
	"context"
	"game/app/slot/api/slot/slot"
	"reflect"
	"testing"
)

func TestSpin(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *slot.SpinReq
	}
	tests := []struct {
		name    string
		args    args
		wantRes *slot.SpinRes
		wantErr bool
	}{
		{
			args: args{
				request: &slot.SpinReq{
					Uid:      136,
					Amount:   1,
					GameCode: 1000,
				},
				ctx: context.Background(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := Spin(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Spin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Spin() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
