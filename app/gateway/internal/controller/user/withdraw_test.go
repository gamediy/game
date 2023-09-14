package user

import (
	"context"
	"game/utility/xetcd"
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"testing"
)

func Test_withdraw_Submit(t *testing.T) {
	xetcd.InitEtcd(context.TODO())
	grpcx.Resolver.Register(etcd.NewWithClient(xetcd.Client))
	wi := withdraw{}
	wi.Submit()

}
