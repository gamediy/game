package user_svc

import (
	"context"
	"game/app/user/api/user/user"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestSendMsgCode(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *user.SendMsgCodeReq
	}
	tests := []struct {
		name    string
		args    args
		wantRes *user.SendMsgCodeRes
		wantErr bool
	}{
		{
			args: args{
				in: &user.SendMsgCodeReq{
					Phone: "123456",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := SendMsgCode(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendMsgCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			g.Dump(gotRes)
		})
	}
}
