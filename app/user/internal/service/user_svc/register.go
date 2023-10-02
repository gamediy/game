package user_svc

import (
	"context"
	"fmt"
	"game/db/dao"
	"game/db/model/entity"
	"game/utility/utils/xip"
	"game/utility/utils/xpwd"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type Register struct {
	Account  string `dc:"账号" json:"account"  v:"required|length:6,30#请输入账号|账号长度为:{min}到:{max}位"`
	Password string `dc:"密码" json:"password" v:"required|length:6,30#请输入密码|密码长度不够"`
	Xid      string `dc:"邀请码" json:"xid" `
	Phone    string `json:"phone"`
	Email    string `json:"email"`

	Ip          string `json:"-"`
	RealName    string
	ClientAgent string
}

func (s *Register) Exec(ctx context.Context) error {
	//if err := g.Validator().Assoc(s).Data(&s).Run(ctx); err != nil {
	//	return err
	//}

	if err := s.checkAccountExist(ctx, s.Account); err != nil {
		return err
	}
	var d entity.User
	loc := xip.GetIpInfoIo(ctx, s.Ip)
	d.Account = s.Account
	d.Status = 1
	d.Ip = s.Ip
	d.City = loc.City
	d.Country = loc.Country
	d.Password = xpwd.GenPwd(s.Password)
	d.Phone = s.Phone
	d.Email = s.Email
	d.RealName = s.RealName
	d.ClientAgent = s.ClientAgent
	parent, err := s.getParent(ctx)
	if err != nil {
		return err
	}
	if parent != nil {
		d.ParentPath = fmt.Sprintf("%s%d/", parent.ParentPath, parent.Id)
	}
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		uid, err := tx.Model(dao.User.Table()).InsertAndGetId(&d)
		if err != nil {
			return err
		}
		if _, err = tx.Model(dao.Wallet.Table()).Ctx(ctx).Insert(&entity.Wallet{
			Uid:     uid,
			Balance: 0,
			Account: s.Account,
		}); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *Register) getParent(ctx context.Context) (*entity.User, error) {
	var d entity.User
	if s.Xid == "" {
		return nil, nil
	}
	one, err := dao.User.Ctx(ctx).One(&d, "xid", s.Xid)
	if one.IsEmpty() {
		return nil, nil
	}
	if err = one.Struct(&d); err != nil {
		return nil, err
	}
	return &d, err
}

func (s *Register) checkAccountExist(ctx context.Context, account string) error {
	count, err := dao.User.Ctx(ctx).Count("account", account)
	if err != nil {
		return err
	}
	if count != 0 {
		return fmt.Errorf("账户存在")
	}
	return nil
}
