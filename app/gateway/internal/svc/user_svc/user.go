package user_svc

import (
	"context"
	"game/app/user/api/deposit/deposit"
	"game/app/user/api/user/user"
	"game/app/user/api/withdraw/withdraw"
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
	withdrawClient withdraw.WithdrawServiceClient
	depositClient  deposit.DepositSvcClient
)

func UserClientInit() {
	conn = grpcx.Client.MustNewGrpcClientConn("user_svc", grpcx.Balancer.WithRandom())
	userClient = user.NewUserServiceClient(conn)
	withdrawClient = withdraw.NewWithdrawServiceClient(conn)
	depositClient = deposit.NewDepositSvcClient(conn)
}

type userSvc struct {
}

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
func (userSvc) Wallet(ctx context.Context, req *user.WalletRequest) (*user.WalletReply, error) {
	return userClient.Wallet(ctx, req)
}
