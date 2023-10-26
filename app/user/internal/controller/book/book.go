package book

import (
	"context"
	"game/app/user/api/book/book"
	"game/app/user/api/book/english"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	book.UnimplementedBookServiceServer
	english.UnimplementedEnglishServiceServer
}

func Register(s *grpcx.GrpcServer) {
	book.RegisterBookServiceServer(s.Server, &Controller{})
}

func (*Controller) GetBook(ctx context.Context, req *book.TestReq) (res *book.TestRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
