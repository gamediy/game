package user

import (
	"context"
	"game/app/user/api/user/withdraw"
	"game/app/user/internal/service/withdraw_svc"
)

func (*Controller) PayPassStatus(ctx context.Context, req *withdraw.PayPassStatusReq) (res *withdraw.PayPassStatusRes, err error) {
	return withdraw_svc.GetPayPassStatus(ctx, req)
}

func (*Controller) SetPayPass(ctx context.Context, req *withdraw.SetPayPassReq) (res *withdraw.SetPayPassRes, err error) {
	x := withdraw_svc.SetPayPass{
		Uid:  req.Uid,
		Pass: req.Pass,
	}
	if err = x.Exec(ctx); err != nil {
		return nil, err
	}
	return &withdraw.SetPayPassRes{}, nil
}

func (*Controller) BindWithdrawAccount(ctx context.Context, req *withdraw.BindWithdrawAccountReq) (res *withdraw.BindWithdrawAccountRes, err error) {
	x := withdraw_svc.BindWithdrawAccount{
		BankId:  req.BankId,
		Address: req.Address,
		Title:   req.Title,
		Pass:    req.Pass,
		Uid:     req.Uid,
	}
	if err = x.Exec(ctx); err != nil {
		return nil, err
	}
	return &withdraw.BindWithdrawAccountRes{}, nil
}

func (*Controller) DelWithdrawAccount(ctx context.Context, req *withdraw.DelWithdrawAccountReq) (res *withdraw.DelWithdrawAccountRes, err error) {
	x := withdraw_svc.DelWithdrawAccount{
		Id:      req.Id,
		Uid:     req.Uid,
		PayPass: req.PayPass,
	}
	if err = x.Exec(ctx); err != nil {
		return nil, err
	}
	return &withdraw.DelWithdrawAccountRes{}, nil
}

func (*Controller) ListWithdrawAccount(ctx context.Context, req *withdraw.ListWithdrawAccountReq) (res *withdraw.ListWithdrawAccountRes, err error) {
	x := withdraw_svc.ListWithdrawAccount{
		Uid:  req.Uid,
		Page: int(req.Page),
		Size: int(req.Size),
	}
	return x.Exec(ctx)
}

func (*Controller) CreateWithdraw(ctx context.Context, req *withdraw.CreateWithdrawReq) (res *withdraw.CreateWithdrawRes, err error) {
	x := withdraw_svc.CreateWithdraw{
		AmountItemId:      int(req.AmountItemId),
		Amount:            req.Amount,
		WithdrawAccountId: int(req.WithdrawAccountId),
		Lang:              req.Lang,
		Uid:               req.Uid,
	}
	if err = x.Exec(ctx); err != nil {
		return nil, err
	}
	return &withdraw.CreateWithdrawRes{}, nil
}

func (*Controller) ListWithdrawMethod(ctx context.Context, req *withdraw.ListWithdrawMethodReq) (res *withdraw.ListWithdrawMethodRes, err error) {
	x := withdraw_svc.ListWithdrawMethod{
		Uid:  req.Uid,
		Lang: req.Lang,
	}
	return x.Exec(ctx)
}
func (*Controller) ListWithdraw(ctx context.Context, req *withdraw.ListWithdrawReq) (res *withdraw.ListWithdrawRes, err error) {
	x := withdraw_svc.ListWithdraw{
		Uid:  req.Uid,
		Page: int(req.Page),
		Size: int(req.Size),
	}
	return x.Exec(ctx)
}
