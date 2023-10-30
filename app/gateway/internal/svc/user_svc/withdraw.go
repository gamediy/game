package user_svc

import (
	"context"
	"game/app/user/api/user/withdraw"
)

func (userSvc) PayPassStatus(ctx context.Context, uid int64) (*withdraw.PayPassStatusRes, error) {
	return withdrawClient.PayPassStatus(ctx, &withdraw.PayPassStatusReq{Uid: uid})
}
func (userSvc) SetPayPass(ctx context.Context, in *withdraw.SetPayPassReq) (*withdraw.SetPayPassRes, error) {
	return withdrawClient.SetPayPass(ctx, in)
}
func (userSvc) BindWithdrawAccount(ctx context.Context, in *withdraw.BindWithdrawAccountReq) (*withdraw.BindWithdrawAccountRes, error) {
	return withdrawClient.BindWithdrawAccount(ctx, in)
}
