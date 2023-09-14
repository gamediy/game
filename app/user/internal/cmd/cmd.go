package cmd

import (
	"context"
	"game/app/user/internal/controller"
	"game/utility/xetcd"
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/net/gsvc"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start crontab job",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			xetcd.InitEtcd(ctx)
			gsvc.SetRegistry(etcd.NewWithClient(xetcd.Client))
			s := grpcx.Server.New()
			controller.RegisterWithdraw(s)
			s.Run()
			return
		},
	}
)
