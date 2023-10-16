package slot

import (
	"context"
	"game/app/slot/api/slot/slot"
	"game/app/slot/internal/service"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"

	_ "github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

func Register(s *grpcx.GrpcServer) {
	slot.RegisterSlotServiceServer(s.Server, &Controller{})
}

type Controller struct {
	slot.UnimplementedSlotServiceServer
}

func (*Controller) Spin(ctx context.Context, req *slot.SpinReq) (res *slot.SpinRes, err error) {
	return service.Spin(ctx, req)
}

func (*Controller) CheckWon(ctx context.Context, req *slot.CheckWonReq) (res *slot.CheckWonRes, err error) {
	return service.CheckWon(ctx, req)
}
