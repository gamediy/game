package user_svc

import (
	"context"
	"game/app/user/api/user/user"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestCountMailBoxTotalUnRead(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *user.MailBoxTotalUnReadReq
	}
	tests := []struct {
		name    string
		args    args
		want    *user.MailBoxTotalUnReadRes
		wantErr bool
	}{
		{args: args{in: &user.MailBoxTotalUnReadReq{Uid: 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CountMailBoxTotalUnRead(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountMailBoxTotalUnRead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			g.Dump(got)
		})
	}
}
