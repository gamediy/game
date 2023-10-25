package user_svc

import (
	"context"
	"game/app/user/api/user/deposit"
)

func (userSvc) ListDepositAmountItems(ctx context.Context) (*deposit.DepositAmountItemsRes, error) {
	items, err := depositClient.ListDepositAmountItems(ctx, &deposit.DepositAmountItemsReq{})
	if err != nil {
		return nil, err
	}
	return items, nil
}
