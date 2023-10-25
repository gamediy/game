package user_svc

import (
	"context"
	"game/app/user/api/user/deposit"
)

func (userSvc) ListDepositAmountItems(ctx context.Context, uid int64) (*deposit.DepositAmountItemsRes, error) {
	items, err := depositClient.ListDepositAmountItems(ctx, &deposit.DepositAmountItemsReq{Uid: uid})
	if err != nil {
		return nil, err
	}
	return items, nil
}
