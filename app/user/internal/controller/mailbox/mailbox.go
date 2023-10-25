package mailbox

import (
	"context"
	"game/app/user/api/user/deposit"
	"game/app/user/api/user/mailbox"
	"game/app/user/internal/service/mailbox_svc"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	mailbox.UnimplementedMailBoxServiceServer
}

func Register(s *grpcx.GrpcServer) {
	mailbox.RegisterMailBoxServiceServer(s.Server, &Controller{})
}
func (*Controller) ListMailBox(ctx context.Context, req *mailbox.ListMailBoxReq) (res *mailbox.ListMailBoxRes, err error) {
	return mailbox_svc.ListMailBox(ctx, req)
}

func (*Controller) CountMailBoxTotalUnRead(ctx context.Context, req *mailbox.MailBoxTotalUnReadReq) (res *mailbox.MailBoxTotalUnReadRes, err error) {
	return mailbox_svc.CountMailBoxTotalUnRead(ctx, req)
}

func (*Controller) ListDepositAmountItems(ctx context.Context, req *deposit.DepositAmountItemsReq) (res *deposit.DepositAmountItemsRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
