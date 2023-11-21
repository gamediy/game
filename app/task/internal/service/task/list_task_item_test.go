package task

import (
	"context"
	"game/app/task/api/task/task"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestListTaskItem(t *testing.T) {
	type args struct {
		ctx context.Context
		req *task.ListTaskItemReq
	}
	tests := []struct {
		name    string
		args    args
		want    *task.ListTaskItemRes
		wantErr bool
	}{
		{
			args: args{req: &task.ListTaskItemReq{
				TaskId: 1,
				Uid:    30,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListTaskItem(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListTaskItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			g.Dump(got)
		})
	}
}
