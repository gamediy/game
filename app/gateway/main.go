package main

import (
	"github.com/gogf/gf/v2/os/gctx"

	"game/app/gateway/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
