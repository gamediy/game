package withdraw

import (
	"context"
	"game/app/user/api/user/withdraw"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	withdraw.UnimplementedWithdrawServiceServer
}

func Register(s *grpcx.GrpcServer) {
	withdraw.RegisterWithdrawServiceServer(s.Server, &Controller{})
}

func (*Controller) Submit(ctx context.Context, req *withdraw.SubmitRequest) (res *withdraw.SubmitReply, err error) {
	res = &withdraw.SubmitReply{}
	res.Code = 100

	return res, nil
}

func (*Controller) Bind(ctx context.Context, req *withdraw.BindRequest) (res *withdraw.SubmitReply, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
