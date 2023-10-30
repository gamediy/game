package withdraw_svc

import (
	"context"
	"fmt"
	"game/db/dao"
	"game/pure/get"
	"game/utility/utils/xpwd"
)

type DelWithdrawAccount struct {
	Uid     int64
	Id      int64
	PayPass string
}

func (m *DelWithdrawAccount) Exec(ctx context.Context) error {
	u, err := get.UserById(ctx, m.Uid)
	if err != nil {
		return err
	}
	if u.PayPass == "" {
		return fmt.Errorf("set pay password first")
	}
	if !xpwd.ComparePassword(u.PayPass, m.PayPass) {
		return fmt.Errorf("pay password is error")
	}
	count, _ := dao.WithdrawAccount.Ctx(ctx).Count("id = ? and uid = ?", m.Id, m.Uid)
	if count == 0 {
		return fmt.Errorf("deleted")
	}
	_, _ = dao.WithdrawAccount.Ctx(ctx).Delete("id", m.Id)
	return nil
}
