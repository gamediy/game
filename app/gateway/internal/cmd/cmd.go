package cmd

import (
	"context"
	"game/app/gateway/internal/sync"
	"game/app/gateway/internal/ws"
	"game/utility/xetcd"
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"
	"math/rand"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start gateway",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			xetcd.InitEtcd(ctx)
			grpcx.Resolver.Register(etcd.NewWithClient(xetcd.Client))
			g.Log().Info(ctx, `gateway start`)
			s := g.Server()
			ws.StartWebSocket(ctx)
			go sync.Router(ctx)
			s.BindHandler("/socket", func(r *ghttp.Request) {
				token := r.GetQuery("token", "")
				if token == nil || token.Val() == nil {
					r.Response.Write("error token")
					return
				}
				if token.Val().(string) == "" {
					return
				}
				socket, err := r.WebSocket()
				if err != nil {
					glog.Error(ctx, err)
					r.Exit()
				}

				client := ws.NewClient(r.GetClientIp(), rand.Int(), "1", socket.Conn)
				go client.Read(ctx)
				go client.Write(ctx)

			})
			s.BindHandler("/login", func(r *ghttp.Request) {
			})
			s.BindHandler("/register", func(r *ghttp.Request) {})
			s.Run()
			return nil
		},
	}
)
