package user_svc

import (
	"context"
	"game/app/user/api/user/user"
	"game/consts"
	"game/db/dao"
	"game/db/model/entity"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	"time"
)

func SendMsgCode(ctx context.Context, in *user.SendMsgCodeReq) (res *user.SendMsgCodeRes, err error) {
	code := &entity.SmsVerification{}
	_ = dao.SmsVerification.Ctx(ctx).
		Where("phone", in.Phone).
		Limit(1).
		Order("id desc").
		WhereGTE("expiration", gtime.Now()).
		Scan(&code)
	if code.Id != 0 {
		return nil, consts.ErrWaitAMinute
	}

	code.Code = grand.Str("1234567890", 5)
	code.Phone = in.Phone
	code.Expiration = gtime.Now().Add(time.Second * 60)
	code.Times = 0
	if _, err = dao.SmsVerification.Ctx(ctx).Insert(code); err != nil {
		return nil, err
	}
	return &user.SendMsgCodeRes{}, nil
}
