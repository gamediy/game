package cmd

import (
	"context"
	"game/app/third/internal/controller/cq9"
	"game/app/third/internal/tools"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			Cq9()
			g.Wait()
			return nil
		},
	}
)

func Cq9() {
	s := g.Server("cq9")
	s.BindMiddlewareDefault(ghttp.MiddlewareCORS, tools.MiddlewareHandlerResponse)
	s.Group("/", func(g *ghttp.RouterGroup) {
		g.Bind(cq9.NewV1())
	})
	s.Start()
}
