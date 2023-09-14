package user

import (
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	ctx  = gctx.New()
	conn = grpcx.Client.MustNewGrpcClientConn("user")
)
