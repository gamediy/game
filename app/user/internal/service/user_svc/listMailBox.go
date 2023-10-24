package user_svc

import (
	"context"
	"game/app/user/api/user/user"
	"game/db/dao"
)

func ListMailBox(ctx context.Context, req *user.ListMailBoxReq) (*user.ListMailBoxRes, error) {
	var (
		total int
		data  = make([]*user.MailBox, 0)
	)
	db := dao.Mailbox.Ctx(ctx)
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
	return &user.ListMailBoxRes{List: data, Total: int64(total)}, nil
}
