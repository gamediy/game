package deposit

import (
	"context"
	"game/app/user/api/user/deposit"
	"game/app/user/internal/service/deposit_svc"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	deposit.UnimplementedDepositServiceServer
}

func Register(s *grpcx.GrpcServer) {
	deposit.RegisterDepositServiceServer(s.Server, &Controller{})
}

func (*Controller) ListDepositAmountItems(ctx context.Context, req *deposit.DepositAmountItemsReq) (res *deposit.DepositAmountItemsRes, err error) {
	items, _ := deposit_svc.ListDepositAmountItems(ctx, req.Uid)
	return items, nil
}
