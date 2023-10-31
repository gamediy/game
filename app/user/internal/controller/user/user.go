package user

import (
	"context"
	"game/app/user/api/user/deposit"
	"game/app/user/api/user/mailbox"
	"game/app/user/api/user/user"
	"game/app/user/api/user/wallet"
	"game/app/user/api/user/withdraw"
	"game/app/user/internal/service/user_svc"
	"game/app/user/internal/service/wallet_svc"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	user.UnimplementedUserServiceServer
	deposit.UnimplementedDepositServiceServer
	mailbox.UnimplementedMailBoxServiceServer
	withdraw.UnimplementedWithdrawServiceServer
	wallet.UnimplementedWalletServiceServer
}

func Register(s *grpcx.GrpcServer) {
	user.RegisterUserServiceServer(s.Server, &Controller{})
	deposit.RegisterDepositServiceServer(s.Server, &Controller{})
	mailbox.RegisterMailBoxServiceServer(s.Server, &Controller{})
	withdraw.RegisterWithdrawServiceServer(s.Server, &Controller{})
	wallet.RegisterWalletServiceServer(s.Server, &Controller{})
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
