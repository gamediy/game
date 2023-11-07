package withdraw_svc

import (
	"context"
	"fmt"
	"game/app/user/api/user/withdraw"
	"game/db/dao"
	"game/db/model/entity"
	"game/pure/get"
	"game/utility/utils/xtrans"
	"github.com/gogf/gf/v2/text/gstr"
	"strconv"
)

type ListWithdrawMethod struct {
	Uid  int64
	Lang string
}

func (m *ListWithdrawMethod) Exec(ctx context.Context) (*withdraw.ListWithdrawMethodRes, error) {
	_, err := get.UserById(ctx, m.Uid)
	if err != nil {
		return nil, err
	}
	var (
		wc         = make([]*entity.WithdrawAccount, 0)
		amountItem = make([]*entity.AmountItem, 0)
		res        = &withdraw.ListWithdrawMethodRes{
			Tips: xtrans.T(m.Lang, "withdrawal prompt get the configuration file"),
			List: make([]*withdraw.AmountItem, 0),
		}
	)
	_ = dao.WithdrawAccount.Ctx(ctx).Where("uid = ? and status = 1", m.Uid).Scan(&wc)
	_ = dao.AmountItem.Ctx(ctx).Where("status = 1 and type = 'Withdraw'").Scan(&amountItem)
	for _, item := range amountItem {
		bind := &withdraw.WithdrawBindInfo{}
		for _, account := range wc {
			if account.Protocol == item.Protocol {
				bind.Address = account.Address
				bind.Protocol = account.Protocol
				bind.WithdrawAccountId = fmt.Sprint(account.Id)
			}
		}
		imgPrefix := get.ImgPrefix()
		res.List = append(res.List, &withdraw.AmountItem{
			Id:           int64(item.Id),
			Title:        item.Title,
			Protocol:     item.Protocol,
			Logo:         imgPrefix + item.Logo,
			Currency:     item.Currency,
			Address:      gstr.HideStr(item.Address, 60, "*"),
			Detail:       fmt.Sprintf("%s %d-%d", item.Detail, item.Min, item.Max),
			ExchangeRate: item.ExchangeRate,
			SelectMoney:  item.SelectMoney,
			Bind:         bind,
			BankId:       strconv.Itoa(item.BankId),
		})
	}
	return res, nil
}
