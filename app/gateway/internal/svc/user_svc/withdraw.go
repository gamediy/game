package user_svc

import (
	"context"
	"game/app/user/api/user/withdraw"
)

func (userSvc) PayPassStatus(ctx context.Context, uid int64) (*withdraw.PayPassStatusRes, error) {
	return withdrawClient.PayPassStatus(ctx, &withdraw.PayPassStatusReq{Uid: uid})
}
