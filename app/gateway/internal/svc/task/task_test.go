package task

import (
	"context"
	"game/app/task/api/task/task"
	"game/utility/utils/xetcd"
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestListTaskItem(t *testing.T) {
	ctx := context.TODO()
	xetcd.InitEtcd(ctx)
	grpcx.Resolver.Register(etcd.NewWithClient(xetcd.Client))
	ClientInit()
	type args struct {
		ctx context.Context
		in  *task.ListTaskItemReq
	}
	tests := []struct {
		name    string
		args    args
		want    *task.ListTaskItemRes
		wantErr bool
	}{
		{
			args: args{
				ctx: ctx,
				in: &task.ListTaskItemReq{
					TaskId: 1,
					Uid:    30,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListTaskItem(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListTaskItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			g.Dump(got)
		})
	}
}
