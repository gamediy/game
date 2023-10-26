package user

import (
	"context"
	"game/app/user/api/user/deposit"
	"game/app/user/internal/service/deposit_svc"
)

func (*Controller) ListDepositAmountItems(ctx context.Context, req *deposit.DepositAmountItemsReq) (res *deposit.DepositAmountItemsRes, err error) {
	return deposit_svc.ListDepositAmountItems(ctx, req.Uid)
}

func (*Controller) CreateDeposit(ctx context.Context, req *deposit.CreateDepositReq) (res *deposit.CreateDepositRes, err error) {
	x := deposit_svc.CreateDeposit{}
	x.PayId = int(req.PayId)
	x.Amount = req.Amount
	x.Uid = req.Uid
	x.Lang = req.Lang
	x.TransferOrderNo = req.TransferOrderNo
	x.TransferImg = req.TransferImg
	orderNo, err := x.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &deposit.CreateDepositRes{OrderNo: orderNo}, nil
}
