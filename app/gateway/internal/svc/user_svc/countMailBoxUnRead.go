package user_svc

import (
	"context"
	"game/app/user/api/user/user"
)

func (userSvc) CountMailBoxUnRead(ctx context.Context, uid int64) (*user.MailBoxTotalUnReadRes, error) {
	return userClient.CountMailBoxTotalUnRead(ctx, &user.MailBoxTotalUnReadReq{Uid: uid})
}
