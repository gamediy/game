package main

import (
	"game/app/gateway/internal/cmd"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"runtime"
)

func main() {
	runtime.SetMutexProfileFraction(1) // (非必需)开启对锁调用的跟踪
	runtime.SetBlockProfileRate(1)
	go ghttp.StartPProfServer(8199)
	cmd.Main.Run(gctx.GetInitCtx())

}
