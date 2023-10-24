package report

import (
	"context"
	"game/db/dao"
	"game/db/model/entity"
	"game/utility/utils/xuuid"
	"github.com/avast/retry-go/v4"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type Report struct {
	User      entity.User
	Amount    float64
	TransType *entity.TransactionType
}

func (m *Report) Exec() error {
	return retry.Do(func() error {
		var (
			log   entity.WalletReportDay
			today = gtime.Now().Format("Y-m-d")
			ctx   = context.Background()
		)
		_ = dao.WalletReportDay.Ctx(ctx).Scan(&log, "date = ? and uid = ? and trans_code = ?", today, m.User.Id, m.TransType.Code)
		if log.Id == 0 {
			log.Uid = m.User.Id
			log.Account = m.User.Account
			log.TransCode = m.TransType.Code
			log.Date = gtime.NewFromStr(today)
			log.Count = 1
			log.Title = m.TransType.Title
			log.Amount = m.Amount
			log.Id = xuuid.GenSnowflakeUUID().Int64()
			if _, err := dao.WalletReportDay.Ctx(ctx).Insert(log); err != nil {
				g.Log().Error(ctx, err)
				return err
			}
		} else {
			log.Amount += m.Amount
			log.Count += 1
			if _, err := dao.WalletReportDay.Ctx(ctx).Update(&log, "id", log.Id); err != nil {
				return err
			}
		}
		return nil
	}, retry.Attempts(3))
}
