package user

import (
	"context"
	"game/app/user/api/user/mailbox"
	"game/app/user/internal/service/mailbox_svc"
)

func (*Controller) ListMailBox(ctx context.Context, req *mailbox.ListMailBoxReq) (res *mailbox.ListMailBoxRes, err error) {
	return mailbox_svc.ListMailBox(ctx, req)
}

func (*Controller) CountMailBoxTotalUnRead(ctx context.Context, req *mailbox.MailBoxTotalUnReadReq) (res *mailbox.MailBoxTotalUnReadRes, err error) {
	return mailbox_svc.CountMailBoxTotalUnRead(ctx, req)
}
