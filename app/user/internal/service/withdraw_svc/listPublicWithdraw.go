package withdraw_svc

import (
	"context"
	"game/app/user/api/user/withdraw"
	"game/consts"
	"game/db/dao"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/text/gstr"
	"time"
)

func ListPublicWithdraw(ctx context.Context) (*withdraw.ListPublicWithdrawRes, error) {
	var res = &withdraw.ListPublicWithdrawRes{}
	res.List = make([]*withdraw.PublicWithdrawItem, 0)
	all, err := gcache.GetOrSetFunc(ctx, "ListPublicWithdraw", func(ctx context.Context) (value interface{}, err error) {
		return dao.Withdraw.Ctx(ctx).Limit(13).All("status", consts.DepositStatusSuccess)
	}, time.Minute)
	if err != nil {
		return nil, err
	}
	if err = all.Scan(&res.List); err != nil {
		return nil, err
	}
	for _, i := range res.List {
		i.Account = gstr.HideStr(i.Account, 70, "*")
	}
	return res, nil
}
