package user_svc

import (
	"context"
	"game/db/dao"
	"game/db/model/entity"
	"game/model"
)

type UserInfo struct {
	Token string `json:"token"`
}

func (s *UserInfo) Get(ctx context.Context) *model.UserInfo {
	user := entity.User{}
	mu := model.UserInfo{}
	dao.User.Ctx(ctx).Where("token=?", s.Token).Scan(&user)
	if user.Id > 0 {
		mu.Uid = user.Id
		mu.Account = user.Account
		mu.ParentPath = user.ParentPath
		mu.Pid = user.Pid
	}
	return &mu
}
