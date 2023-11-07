package withdraw_svc

import (
	"context"
	"fmt"
	"game/db/dao"
	"game/db/model/entity"
	"game/pure/get"
	"game/utility/utils/xpwd"
	"game/utility/utils/xuuid"
)

type BindWithdrawAccount struct {
	BankId  int64  `v:"required#银行ID不能为空" dc:"银行ID"`
	Address string `v:"required#收款地址不能为空"`
	Title   string `json:"title" dc:"持卡人" v:"required#持卡人不能为空"`
	Uid     int64
}

func (s BindWithdrawAccount) Exec(ctx context.Context) error {
	bank, err := get.Bank(ctx, s.BankId)
	if err != nil {
		return err
	}
	user, err := get.UserById(ctx, s.Uid)
	if err != nil {
		return err
	}
	// check already bind
	count, _ := dao.WithdrawAccount.Ctx(ctx).Count("address = ? and uid = ?", s.Address, user.Id)
	if count != 0 {
		return fmt.Errorf("this card has been bound")
	}

	d := entity.WithdrawAccount{}
	d.Id = xuuid.GenSnowflakeUUID().Int64()
	d.Net = bank.Net
	d.Protocol = bank.Protocol
	d.Uid = user.Id
	d.Account = user.Account
	d.Currency = bank.Currency
	d.Address = s.Address
	d.Status = 1
	d.Default = -1
	d.Title = s.Title
	if _, err = dao.WithdrawAccount.Ctx(ctx).Insert(d); err != nil {
		return err
	}
	return nil
}

func (s BindWithdrawAccount) CheckPayPass(ctx context.Context, pass string, uid int64) error {
	var u entity.User
	_ = dao.User.Ctx(ctx).Scan(&u, "id", uid)
	if u.Id == 0 {
		return fmt.Errorf("用户不存在")
	}
	if u.PayPass == "" {
		return fmt.Errorf("请先设置交易密码")
	}
	if !xpwd.ComparePassword(u.PayPass, pass) {
		return fmt.Errorf("交易密码错误")
	}
	return nil
}
