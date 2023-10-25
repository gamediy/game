package user

import (
	"context"
	"game/app/user/api/user/deposit"
	"game/app/user/api/user/mailbox"
	"game/app/user/api/user/user"
	"game/app/user/api/user/withdraw"
	"game/app/user/internal/service/deposit_svc"
	"game/app/user/internal/service/mailbox_svc"

	"game/app/user/internal/service/user_svc"
	"game/app/user/internal/service/wallet_svc"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	user.UnimplementedUserServiceServer
	withdraw.UnimplementedWithdrawServiceServer
	deposit.UnimplementedDepositServiceServer
	mailbox.UnimplementedMailBoxServiceServer
}

func Register(s *grpcx.GrpcServer) {
	user.RegisterUserServiceServer(s.Server, &Controller{})
	withdraw.RegisterWithdrawServiceServer(s.Server, &Controller{})
	deposit.RegisterDepositServiceServer(s.Server, &Controller{})
	mailbox.RegisterMailBoxServiceServer(s.Server, &Controller{})
}

func (*Controller) Reg(ctx context.Context, req *user.RegRequest) (res *user.RegReply, err error) {
	res = &user.RegReply{}
	register := user_svc.Register{}
	err = gconv.Struct(req, &register)
	if err != nil {
		g.Log().Errorf(ctx, err.Error())
		return nil, err
	}
	token, err := register.Exec(ctx)
	if err != nil {
		return res, err
	}
	res.Token = token
	return res, nil
}

func (*Controller) Login(ctx context.Context, req *user.LoginRequest) (res *user.LoginReply, err error) {

	res = &user.LoginReply{}
	login := user_svc.Login{}
	gconv.Struct(req, &login)
	token, err := login.Exec(ctx)
	if err != nil {
		return nil, err
	}
	res.Token = token
	return res, nil

}

func (*Controller) UserInfo(ctx context.Context, req *user.UserInfoRequest) (res *user.UserInfoReply, err error) {
	info := user_svc.UserInfo{}
	info.Token = req.Token
	get := info.Get(ctx)
	res = &user.UserInfoReply{}
	gconv.Struct(get, res)
	return res, nil
}

func (*Controller) Wallet(ctx context.Context, req *user.WalletRequest) (res *user.WalletReply, err error) {
	return wallet_svc.Wallet(ctx, req)
}

func (*Controller) ListDepositAmountItems(ctx context.Context, req *deposit.DepositAmountItemsReq) (res *deposit.DepositAmountItemsRes, err error) {
	return deposit_svc.ListDepositAmountItems(ctx, req.Uid)
}

func (*Controller) CreateDeposit(ctx context.Context, req *deposit.CreateDepositReq) (res *deposit.CreateDepositRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ListMailBox(ctx context.Context, req *mailbox.ListMailBoxReq) (res *mailbox.ListMailBoxRes, err error) {
	return mailbox_svc.ListMailBox(ctx, req)
}

func (*Controller) CountMailBoxTotalUnRead(ctx context.Context, req *mailbox.MailBoxTotalUnReadReq) (res *mailbox.MailBoxTotalUnReadRes, err error) {
	return mailbox_svc.CountMailBoxTotalUnRead(ctx, req)
}

func (*Controller) Submit(ctx context.Context, req *withdraw.SubmitRequest) (res *withdraw.SubmitReply, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) Bind(ctx context.Context, req *withdraw.BindRequest) (res *withdraw.SubmitReply, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
