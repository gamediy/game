package mailbox_svc

import (
	"context"
	"game/app/user/api/user/mailbox"
	"game/db/dao"
)

func ListMailBox(ctx context.Context, req *mailbox.ListMailBoxReq) (*mailbox.ListMailBoxRes, error) {
	var (
		total int
		data  = make([]*mailbox.MailBox, 0)
	)
	db := dao.Mailbox.Ctx(ctx)
	if req.Type != "" {
		db = db.Where("type", req.Type)
	}
	if req.Receiver != "" {
		db = db.Where("receiver", req.Receiver)
	}
	if req.Read != "" {
		db = db.Where("read", req.Read)
	}
	err := db.Page(int(req.Page), int(req.Size)).ScanAndCount(&data, &total, false)
	if err != nil {
		return nil, err
	}
	return &mailbox.ListMailBoxRes{List: data, Total: int64(total)}, nil
}
