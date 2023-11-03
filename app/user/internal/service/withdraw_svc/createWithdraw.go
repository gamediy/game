package withdraw_svc

import (
	"context"
	"fmt"
	"game/consts"
	"game/core/sync"
	"game/core/wallet"
	"game/db/dao"
	"game/db/model/entity"
	"game/pure/get"
	"game/utility/utils/xtime"
	"game/utility/utils/xtrans"
	"game/utility/utils/xuuid"
	"github.com/gogf/gf/v2/database/gdb"
)

type CreateWithdraw struct {
	AmountItemId      int
	Amount            float64
	WithdrawAccountId int
	Lang              string
	Uid               int64
}

func (m *CreateWithdraw) Exec(ctx context.Context) error {
	u, err := get.UserById(ctx, m.Uid)
	if err != nil {
		return err
	}

	withdrawInfo := entity.AmountItem{}
	_ = dao.AmountItem.Ctx(ctx).Scan(&withdrawInfo, "id", m.AmountItemId)
	if withdrawInfo.Id == 0 || withdrawInfo.Status == -1 {
		return fmt.Errorf("the recharge channel has been closed")
	}

	if int64(m.Amount) < withdrawInfo.Min || int64(m.Amount) > withdrawInfo.Max {
		return fmt.Errorf("the recharge amount is incorrect")
	}
	withdrawAccount := entity.WithdrawAccount{}
	_ = dao.WithdrawAccount.Ctx(ctx).Scan(&withdrawAccount, "id", m.WithdrawAccountId)
	if withdrawAccount.Id == 0 {
		return fmt.Errorf("bind the withdrawal account first")
	}
	order := entity.Withdraw{}
	count, err := dao.Withdraw.Ctx(ctx).Count("uid=? and status=0", u.Id)
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("please do not resubmit")
	}
	order.Id = xuuid.GenSnowflakeUUID().Int64()
	order.Uid = u.Id
	order.Account = u.Account
	order.Status = consts.DepositStatusPending
	order.Net = withdrawAccount.Net
	order.Protocol = withdrawAccount.Protocol
	order.Currency = withdrawAccount.Currency
	order.Address = withdrawAccount.Address
	order.StatusRemark = xtrans.T(m.Lang, "processing")
	order.Amount = m.Amount
	order.Pid = u.Pid
	order.ParentPath = u.ParentPath
	order.FinishAt = xtime.Get1970Datetime()
	order.AmountItemId = withdrawInfo.Id
	order.Detail = fmt.Sprintf("%s %s", withdrawInfo.Title, withdrawInfo.Detail)
	order.ExchangeRate = withdrawInfo.ExchangeRate
	order.ExchangeMoney = m.Amount / withdrawInfo.ExchangeRate

	update := wallet.Balance{}
	update.Uid = u.Id
	update.Amount = order.Amount

	update.OrderNoRelation = order.Id
	update.Note = fmt.Sprintf("%s_%s", order.Protocol, order.Address)
	update.TransCode = wallet.Withdraw
	return update.Update(ctx, func(ctx context.Context, tx gdb.TX) error {
		if _, err = tx.Model(dao.Withdraw.Table()).Ctx(ctx).Insert(&order); err != nil {
			return err
		}
		message := sync.Message{}
		if err = message.Trigger(sync.ChannelAdmin, sync.EventWithdraw); err != nil {
			return err
		}
		return err
	})
}
