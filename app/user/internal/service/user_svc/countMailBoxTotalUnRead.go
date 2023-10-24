package user_svc

import (
	"context"
	"game/app/user/api/user/user"
	"game/db/dao"
)

func CountMailBoxTotalUnRead(ctx context.Context, in *user.MailBoxTotalUnReadReq) (*user.MailBoxTotalUnReadRes, error) {
	array, err := dao.MailboxAlreadyRead.Ctx(ctx).Array("mail_box_id", "uid", in.Uid)
	if err != nil {
		return nil, err
	}
	db := dao.Mailbox.Ctx(ctx).Where("receiver in (0,?) and `read` = 0", in.Uid)
	if len(array) != 0 {
		db = db.WhereNotIn("id", array)
	}
	count, err := db.Count()
	if err != nil {
		return nil, err
	}
	var res = &user.MailBoxTotalUnReadRes{
		Num: int64(count),
	}
	return res, nil
}
