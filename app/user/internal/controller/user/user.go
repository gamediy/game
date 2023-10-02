package user

import (
	"context"
	"game/app/user/api/user/user"
	"game/app/user/internal/service/user_svc"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	user.UnimplementedUserServiceServer
}

func Register(s *grpcx.GrpcServer) {
	user.RegisterUserServiceServer(s.Server, &Controller{})
}

func (*Controller) Reg(ctx context.Context, req *user.RegRequest) (res *user.RegReply, err error) {
	res = &user.RegReply{}
	register := user_svc.Register{}
	err = gconv.Struct(req, &register)
	if err != nil {
		g.Log().Errorf(ctx, err.Error())
		return nil, err
	}
	err = register.Exec(ctx)
	if err != nil {
		res.Code = 50
		res.Msg = err.Error()
		return res, err
	}
	res.Code = 0
	return res, nil
}

func (*Controller) Login(ctx context.Context, req *user.LoginRequest) (res *user.LoginReply, err error) {

	res = &user.LoginReply{}

	login := user_svc.Login{}
	token, err := login.Exec(ctx)
	if err != nil {
		res.Code = 50
		res.Msg = err.Error()
		return nil, err
	}
	res.Msg = token
	return res, nil

}
