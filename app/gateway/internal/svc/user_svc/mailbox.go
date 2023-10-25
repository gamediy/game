package user_svc

import (
	"context"
	"game/app/user/api/user/mailbox"
)

func (userSvc) CountMailBoxUnRead(ctx context.Context, uid int64) (*mailbox.MailBoxTotalUnReadRes, error) {
	return mailBoxClient.CountMailBoxTotalUnRead(ctx, &mailbox.MailBoxTotalUnReadReq{Uid: uid})
}
func (userSvc) ListMailBox(ctx context.Context, req *mailbox.ListMailBoxReq) (*mailbox.ListMailBoxRes, error) {
	return mailBoxClient.ListMailBox(ctx, req)
}
