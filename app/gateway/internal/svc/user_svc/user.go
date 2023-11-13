package user_svc

import (
	"context"
	"game/app/user/api/user/deposit"
	"game/app/user/api/user/mailbox"
	"game/app/user/api/user/sys"
	"game/app/user/api/user/wallet"
	"game/app/user/api/user/withdraw"

	"game/app/user/api/user/user"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/util/gconv"
	"google.golang.org/grpc"
)

var (
	Service = &userSvc{}
)

var (
	conn           *grpc.ClientConn
	userClient     user.UserServiceClient
	walletClient   wallet.WalletServiceClient
	depositClient  deposit.DepositServiceClient
	mailBoxClient  mailbox.MailBoxServiceClient
	withdrawClient withdraw.WithdrawServiceClient
	sysClient      sys.SysServiceClient
)

func UserClientInit() {
	conn = grpcx.Client.MustNewGrpcClientConn("user_svc", grpcx.Balancer.WithRandom())
	userClient = user.NewUserServiceClient(conn)
	walletClient = wallet.NewWalletServiceClient(conn)
	depositClient = deposit.NewDepositServiceClient(conn)
	mailBoxClient = mailbox.NewMailBoxServiceClient(conn)
	withdrawClient = withdraw.NewWithdrawServiceClient(conn)
	sysClient = sys.NewSysServiceClient(conn)
}

type userSvc struct{}

func (userSvc) Register(ctx context.Context, in *user.RegRequest) (string, error) {
	token, err := userClient.Reg(ctx, in)
	return token.Token, err
}

func (userSvc) Login(ctx context.Context, in *user.LoginRequest) (string, error) {
	login, err := userClient.Login(ctx, in)
	if err != nil {
		return "", err
	}
	return login.Token, err
}
func (userSvc) UserInfo(ctx context.Context, in *user.UserInfoRequest) (*user.UserInfoReply, error) {
	userInfo, err := userClient.UserInfo(ctx, in)
	if err != nil {
		return nil, err
	}
	reply := user.UserInfoReply{}
	gconv.Struct(userInfo, &reply)
	return &reply, err
}
