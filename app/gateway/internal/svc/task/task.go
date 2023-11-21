package task

import (
	"context"
	"game/app/task/api/task/task"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"google.golang.org/grpc"
)

var (
	conn       *grpc.ClientConn
	taskClient task.TaskServiceClient
)

func ClientInit() {
	conn = grpcx.Client.MustNewGrpcClientConn("task_svc", grpcx.Balancer.WithRandom())
	taskClient = task.NewTaskServiceClient(conn)
}
func ListTaskItem(ctx context.Context, in *task.ListTaskItemReq) (*task.ListTaskItemRes, error) {
	return taskClient.ListTaskItem(ctx, in)
}
