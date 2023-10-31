package wallet_svc

import (
	"context"
	"game/app/user/api/user/wallet"
	"game/db/dao"
)

type ListChangeLog struct {
	Uid  int64
	Page int
	Size int
}

func (m *ListChangeLog) Exec(ctx context.Context) (*wallet.ListChangeLogRes, error) {
	var res = &wallet.ListChangeLogRes{}
	res.List = make([]*wallet.ChangeLogItem, 0)
	db := dao.WalletChangeLog.Ctx(ctx).Where("uid", m.Uid)
	count, err := db.Count()
	if err != nil {
		return nil, err
	}
	res.Total = int64(count)
	if m.Size <= 0 || m.Size >= 100 {
		m.Size = 10
	}
	_ = db.Page(m.Page, m.Size).Order("id desc").Scan(&res.List)
	return res, nil
}
