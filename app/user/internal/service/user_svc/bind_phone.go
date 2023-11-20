package user_svc

import (
	"context"
	"fmt"
	"game/app/user/api/user/user"
	"game/db/dao"
	"game/db/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

func BindPhone(ctx context.Context, in *user.BindPhoneReq) (res *user.BindPhoneRes, err error) {
	code := &entity.SmsVerification{}
	_ = dao.SmsVerification.Ctx(ctx).Limit(1).
		Where("phone", in.Phone).
		WhereLTE("expiration", gtime.Now()).
		Scan(&code)
	if code.Id == 0 {
		return nil, fmt.Errorf("verification code has expired")
	}
	if code.Code != in.Code {
		return nil, fmt.Errorf("verification code error")
	}

	count, err := dao.User.Ctx(ctx).Count("phone", in.Phone)
	if err != nil {
		return nil, err
	}
	if count != 0 {
		return nil, fmt.Errorf("the phone is already in use")
	}

	if _, err = dao.User.Ctx(ctx).Where("id", in.Uid).
		Data(g.Map{"phone": in.Phone}).Update(); err != nil {
		return nil, err
	}
	return &user.BindPhoneRes{
		Phone: in.Phone,
	}, nil
}
