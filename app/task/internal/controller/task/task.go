package task

import (
	"context"
	"game/app/task/api/task/task"
	tasksvc "game/app/task/internal/service/task"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	task.UnimplementedTaskServiceServer
}

func Register(s *grpcx.GrpcServer) {
	task.RegisterTaskServiceServer(s.Server, &Controller{})
}

func (*Controller) ListTaskItem(ctx context.Context, req *task.ListTaskItemReq) (res *task.ListTaskItemRes, err error) {
	return tasksvc.ListTaskItem(ctx, req)
}
