package user_svc

import (
	"context"
	"fmt"
	"game/app/user/api/user/user"
	"game/app/user/api/withdraw/withdraw"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"google.golang.org/grpc"
)

var (
	Service = &userSvc{}
)

var (
	conn           *grpc.ClientConn
	userClient     user.UserServiceClient
	withdrawClient withdraw.WithdrawServiceClient
)

func UserClientInit() {
	conn = grpcx.Client.MustNewGrpcClientConn("user_svc", grpcx.Balancer.WithRandom())
	userClient = user.NewUserServiceClient(conn)
	withdrawClient = withdraw.NewWithdrawServiceClient(conn)
}

type userSvc struct {
}

func (userSvc) Register(ctx context.Context, in *user.RegRequest) error {

	register, err := userClient.Reg(ctx, in)
	if err != nil {
		return err
	}
	if register.Code != 0 {
		return fmt.Errorf(register.Msg)
	}
	return err

}
