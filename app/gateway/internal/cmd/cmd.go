package cmd

import (
	"context"
	"game/app/gateway/internal/svc/user_svc"
	"game/app/gateway/internal/sync"
	"game/app/gateway/internal/ws"
	"game/app/user/api/user/user"

	"game/model"
	"game/utility/utils/xetcd"
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
			user_svc.UserClientInit()

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
			s.BindHandler("/api/login", func(r *ghttp.Request) {
			})
			s.BindHandler("/api/register", func(r *ghttp.Request) {
				// 设置一个超时上下文

				request := user.RegRequest{}
				err := r.Parse(&request)
				if err != nil {
					r.Response.WriteJsonExit(model.WrapHttpMessage(nil, err))
				}
				request.Ip = r.GetClientIp()
				request.ClientAgent = r.GetHeader("User-Agent")
				err = user_svc.Service.Register(ctx, &request)

				message := model.WrapHttpMessage(nil, err)
				r.Response.Write(message)
			})
			s.SetPort(5000)
			s.Run()
			return nil
		},
	}
)
