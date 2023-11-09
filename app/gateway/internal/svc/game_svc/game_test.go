package game_svc

import (
	"game/utility/utils/xetcd"
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestGameClientInit(t *testing.T) {
	ctx := gctx.New()
	xetcd.InitEtcd(ctx)
	grpcx.Resolver.Register(etcd.NewWithClient(xetcd.Client))
	GameClientInit()
	banner, err := ListBanner(ctx)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(banner)
}
