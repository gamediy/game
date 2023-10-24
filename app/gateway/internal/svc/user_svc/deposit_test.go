package user_svc

import (
	"context"
	"game/utility/utils/xetcd"
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"testing"
)

func Test_userSvc_List(t *testing.T) {
	ctx := context.TODO()
	xetcd.InitEtcd(ctx)
	grpcx.Resolver.Register(etcd.NewWithClient(xetcd.Client))
	UserClientInit()
	_, err := Service.ListDepositAmountItems(ctx)
	if err != nil {
		t.Fatal(err)
	}

}
