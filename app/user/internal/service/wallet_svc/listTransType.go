package wallet_svc

import (
	"context"
	"game/app/user/api/user/wallet"
	"game/db/dao"
)

type ListTransType struct {
}

func (ListTransType) Exec(ctx context.Context) (*wallet.ListTransTypeRes, error) {
	var res = &wallet.ListTransTypeRes{}
	res.List = make([]*wallet.TransTypeItem, 0)
	err := dao.TransactionType.Ctx(ctx).Scan(&res.List)
	return res, err
}
