package cmd

import (
	"context"
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
	s.SetPort(9001)
	s.BindHandler("//player/check/:account", func(r *ghttp.Request) {

		r.Response.Write(("cq9"))
	})
	s.Start()
}
