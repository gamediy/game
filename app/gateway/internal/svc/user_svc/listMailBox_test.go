package user_svc

import (
	"context"
	"game/app/user/api/user/user"
	"game/utility/utils/xetcd"
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func Test_userSvc_ListMailBox(t *testing.T) {

	xetcd.InitEtcd(context.TODO())
	grpcx.Resolver.Register(etcd.NewWithClient(xetcd.Client))
	UserClientInit()
	type args struct {
		ctx context.Context
		req *user.ListMailBoxReq
	}
	tests := []struct {
		name    string
		args    args
		want    *user.ListMailBoxRes
		wantErr bool
	}{
		{
			args: args{ctx: context.TODO(), req: &user.ListMailBoxReq{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := userSvc{}
			got, err := us.ListMailBox(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListMailBox() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			g.Dump(got)
		})
	}
}
