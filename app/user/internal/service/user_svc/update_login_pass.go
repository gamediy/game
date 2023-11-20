package user_svc

import (
	"context"
	"game/app/user/api/user/user"
	"game/consts"
	"game/db/dao"
	"game/db/model/entity"
	"game/utility/utils/xpwd"
)

func UpdateLoginPass(ctx context.Context, in *user.UpdateLoginPassReq) (err error) {
	var u *entity.User
	if err = dao.User.Ctx(ctx).Scan(&u, "id", in.Uid); err != nil {
		return err
	}
	if u.Id == 0 {
		return consts.ErrDataNotFound
	}
	if !xpwd.ComparePassword(u.Password, in.OldPass) {
		return consts.ErrOldPass
	}
	if _, err = dao.User.Ctx(ctx).
		Where("id", in.Uid).
		OmitEmptyData().
		Update(&entity.User{Password: xpwd.GenPwd(in.NewPass)}); err != nil {
		return err
	}
	return
}
