package user

import (
	"context"
	"game/app/user/api/user/withdraw"
	"game/app/user/internal/service/withdraw_svc"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
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
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) CreateWithdraw(ctx context.Context, req *withdraw.CreateWithdrawReq) (res *withdraw.CreateWithdrawRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
