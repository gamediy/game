package deposit_svc

import (
	"context"
	"game/app/user/api/user/deposit"

	"game/db/dao"
)

func ListDepositAmountItems(ctx context.Context) *deposit.DepositAmountItemsRes {
	var (
		d    deposit.DepositAmountItemsRes
		list = make([]*deposit.UAmountItemEntry, 0)
	)
	_ = dao.AmountItem.Ctx(ctx).Scan(&list, "status = 1")
	d.List = list
	return &d
}
