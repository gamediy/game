package deposit

import (
	"context"
	"game/app/user/api/deposit/deposit"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	deposit.UnimplementedDepositSvcServer
}

func Register(s *grpcx.GrpcServer) {
	deposit.RegisterDepositSvcServer(s.Server, &Controller{})
}

func (*Controller) ListDepositAmountItems(ctx context.Context, req *deposit.DepositAmountItemsReq) (res *deposit.DepositAmountItemsRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
