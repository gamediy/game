package controller

import (
	"context"
	"game/app/user/pb/withdraw"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Withdraw struct {
	withdraw_pb.UnimplementedWithdrawServiceServer
}

func RegisterWithdraw(s *grpcx.GrpcServer) {
	withdraw_pb.RegisterWithdrawServiceServer(s.Server, &Withdraw{})
}

// SayHello implements helloworld.GreeterServer
func (s *Withdraw) Submit(ctx context.Context, in *withdraw_pb.SubmitRequest) (reply *withdraw_pb.SubmitReply, err error) {
	reply = &withdraw_pb.SubmitReply{}
	reply.Code = 100
	return reply, nil
}
