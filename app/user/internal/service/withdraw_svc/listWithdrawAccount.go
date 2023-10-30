package withdraw_svc

import (
	"context"
	"game/app/user/api/user/withdraw"
	"game/db/dao"
	"github.com/gogf/gf/v2/text/gstr"
)

type ListWithdrawAccount struct {
	Uid  int64
	Page int
	Size int
}

func (m *ListWithdrawAccount) Exec(ctx context.Context) (*withdraw.ListWithdrawAccountRes, error) {
	var res = &withdraw.ListWithdrawAccountRes{}
	res.List = make([]*withdraw.WithdrawAccountItem, 0)
	db := dao.WithdrawAccount.Ctx(ctx).Where("uid", m.Uid)
	count, err := db.Count()
	if err != nil {
		return nil, err
	}
	res.Total = int64(count)
	page := m.Page
	size := m.Size
	if size <= 0 || size >= 50 {
		size = 10
	}
	err = db.Page(page, size).Scan(&res.List)
	if len(res.List) > 0 {
		for _, i := range res.List {
			i.Address = gstr.HideStr(i.Address, 70, "*")
		}
	}
	return res, nil
}
