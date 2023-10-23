package cmd

import (
	"context"
	"fmt"
	"game/app/gateway/internal/svc/slot_svc"
	"game/app/gateway/internal/svc/user_svc"
	"game/app/gateway/internal/sync"
	"game/app/gateway/internal/ws"
	"game/app/user/api/user/user"
	"github.com/gogf/gf/v2/util/gconv"

	"game/model"
	"game/utility/utils/xetcd"
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"
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
			slot_svc.SlotClientInit()

			g.Log().Info(ctx, `gateway start`)
			s := g.Server()
			ws.StartWebSocket(ctx)
			ws.MakeController()
			sync.Router(ctx)
			s.BindMiddlewareDefault(func(r *ghttp.Request) {
				r.Response.CORSDefault()
				r.Middleware.Next()
			})
			s.BindHandler("/socket", func(r *ghttp.Request) {
				token := r.GetQuery("token", "")
				if token == nil || token.Val() == nil {
					r.Response.WriteJsonExit(model.WrapMessage(nil, fmt.Errorf("token  error")))
				}
				if token.Val().(string) == "" {

					r.Response.WriteJsonExit(model.WrapMessage(nil, fmt.Errorf("token  error")))
				}
				info, err := user_svc.Service.UserInfo(ctx, &user.UserInfoRequest{
					Token: token.Val().(string),
				})
				userInfo := &model.UserInfo{}
				gconv.Struct(info, userInfo)
				if err != nil {
					r.Response.WriteJsonExit(model.WrapMessage(nil, fmt.Errorf("token  error")))
				}
				socket, err := r.WebSocket()
				if err != nil {
					glog.Error(ctx, err)
					r.Exit()
				}
				client := ws.NewClient(r.GetClientIp(), userInfo, socket.Conn)
				go client.Read(ctx)
				go client.Write(ctx)

			})
			s.BindHandler("/api/login", func(r *ghttp.Request) {
				request := user.LoginRequest{}
				err := r.Parse(&request)
				if err != nil {
					r.Response.WriteJsonExit(model.WrapMessage(nil, err))
				}
				request.Ip = r.GetClientIp()
				request.ClientAgent = r.GetHeader("User-Agent")
				token, err := user_svc.Service.Login(ctx, &request)
				r.Response.WriteJson(model.WrapMessage(token, err))
			})
			s.BindHandler("/api/register", func(r *ghttp.Request) {
				request := user.RegRequest{}
				err := r.Parse(&request)
				if err != nil {
					r.Response.WriteJsonExit(model.WrapMessage(nil, err))
				}
				request.Ip = r.GetClientIp()
				request.ClientAgent = r.GetHeader("User-Agent")
				token, err := user_svc.Service.Register(ctx, &request)
				message := model.WrapMessage(token, err)
				r.Response.Write(message)
			})
			//第三方游戏
			s.Group("/api/third", func(group *ghttp.RouterGroup) {
				group.GET("/check_account", func(r *ghttp.Request) {

				})
			})

			s.SetPort(5000)
			s.Run()
			return nil
		},
	}
)
