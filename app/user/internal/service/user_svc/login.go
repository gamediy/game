package user_svc

import (
	"context"
	"fmt"
	"game/db/dao"
	"game/db/model/entity"
	"game/utility/utils/xpwd"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/guid"
)

type Login struct {
	Account  string `json:"account" v:"required|length:6,30#请输入账号|账号长度为:{min}到:{max}位"`
	Password string `dc:"密码" json:"password" v:"required|length:6,30#请输入密码|密码长度不够"`
	Ip       string
}

func (s *Login) Exec(ctx context.Context) (token string, err error) {
	var user entity.User
	_ = dao.User.Ctx(ctx).Scan(&user, "account", s.Account)
	if user.Id == 0 {
		return "", fmt.Errorf("用户不存在")
	}
	if !xpwd.ComparePassword(user.Password, s.Password) {
		return "", fmt.Errorf("密码错误")
	}
	d := entity.UserLoginLog{}
	d.Account = s.Account
	d.Uid = uint64(user.Id)
	d.Ip = s.Ip
	_, err = dao.UserLoginLog.Ctx(ctx).Insert(d)
	if err != nil {
		return "", err
	}
	user.Token = guid.S()
	dao.User.Ctx(ctx).Data(g.Map{
		"token": user.Token,
	}).Where("id", user.Id).Update()
	return user.Token, nil
}
