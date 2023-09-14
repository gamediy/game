package main

import (
	_ "game/app/svc-template/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"game/app/svc-template/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
