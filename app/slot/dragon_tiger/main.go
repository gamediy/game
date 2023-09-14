package main

import (
	_ "game/app/gateway/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
