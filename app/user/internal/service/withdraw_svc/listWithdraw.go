package withdraw_svc

import (
	"context"
	"game/app/user/api/user/withdraw"
	"game/db/dao"
)

type ListWithdraw struct {
	Uid  int64
	Page int
	Size int
	Lang string
}

func (m *ListWithdraw) Exec(ctx context.Context) (*withdraw.ListWithdrawRes, error) {
	var res = &withdraw.ListWithdrawRes{}
	res.List = make([]*withdraw.WithdrawItem, 0)
	db := dao.Withdraw.Ctx(ctx).Where("uid", m.Uid)
	count, err := db.Count()
	if err != nil {
		return nil, err
	}
	res.Total = int64(count)
	if m.Size <= 0 || m.Size > 100 {
		m.Size = 10
	}
	err = db.Page(m.Page, m.Size).Scan(&res.List)
	return res, nil
}
