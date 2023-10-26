package deposit_svc

import (
	"context"
	"fmt"
	"game/consts"
	"game/db/dao"
	"game/db/model/entity"
	"game/pure/get"
	"game/utility/utils/xtime"
	"game/utility/utils/xtrans"
	"game/utility/utils/xuuid"
)

type CreateDeposit struct {
	PayId           int
	Amount          float64
	TransferOrderNo string
	TransferImg     string
	Lang            string
	Uid             int64
}

func (input CreateDeposit) Exec(ctx context.Context) (int64, error) {

	payInfo := entity.AmountItem{}
	_ = dao.AmountItem.Ctx(ctx).Scan(&payInfo, "id", input.PayId)
	if payInfo.Id == 0 || payInfo.Status == 0 {
		return 0, xtrans.TError(input.Lang, "the recharge channel has been closed")
	}

	if int64(input.Amount) < payInfo.Min || int64(input.Amount) > payInfo.Max {
		return 0, xtrans.TError(input.Lang, "")
	}
	userInfo, err := get.UserById(ctx, input.Uid)
	if err != nil {
		return 0, err
	}

	order := entity.Deposit{}
	order.OrderNo = xuuid.GenSnowflakeUUID().Int64()
	order.Uid = int(userInfo.Id)
	order.Account = userInfo.Account
	order.Status = consts.DepositStatusPending
	order.Net = payInfo.Net
	order.Protocol = payInfo.Protocol
	order.Currency = payInfo.Currency
	order.Address = payInfo.Address
	order.StatusRemark = xtrans.T(input.Lang, "处理中")
	order.Amount = input.Amount
	order.Pid = int(userInfo.Pid)
	order.ParentPath = userInfo.ParentPath
	order.FinishAt = xtime.Get1970Datetime()
	order.AmountItemId = payInfo.Id
	order.Detail = fmt.Sprintf("%s %s", payInfo.Title, payInfo.Detail)
	order.TransferOrderNo = input.TransferOrderNo
	order.TransferImg = input.TransferImg
	if order.Address == "" {
		if order.Protocol == "trc20" { //在获取充值列表的时候要给用户生成一个充值地址 银行卡 address如果为空可能需要调用支付三方API
			daccount := entity.DigitalAccount{}
			_ = dao.DigitalAccount.Ctx(ctx).Where("uid", userInfo.Id).Scan(&daccount)
			order.Address = daccount.Address
		}
	}

	insert, err := dao.Deposit.Ctx(ctx).Data(&order).Insert()
	if err != nil {
		return 0, err
	}
	affected, err := insert.RowsAffected()
	if err != nil {
		return 0, err
	}
	if affected < 1 {
		return 0, xtrans.TError(input.Lang, "失败请重试")
	}
	return order.OrderNo, nil
}
