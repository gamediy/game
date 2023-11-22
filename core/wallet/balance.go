package wallet

import (
	"context"
	"fmt"
	"game/consts"
	"game/core/report"
	"game/db/dao"
	"game/db/model/entity"
	"game/pure/get"
	"game/utility/utils/xuuid"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"math"
	"time"
)

const (
	Unfreeze     = 800
	DepositBank  = 501
	Deposit      = 500
	SignIn       = 401
	Gift         = 400
	Profit       = 201
	WithdrawFail = 100
	Withdraw     = -100
	Bet          = -700
	Freeze       = -800
)

type Balance struct {
	Uid             int64
	Amount          float64
	TransCode       int
	OrderNoRelation int64
	Note            string
}

func (m *Balance) Update(ctx context.Context, fc ...func(ctx context.Context, tx gdb.TX) error) error {
	timeout, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	var err error
	for {
		err = m.exec(ctx, fc)
		if err != nil {
			g.Log().Error(ctx, err)
		}
		if err == nil {
			return nil
		}
		select {
		case <-timeout.Done():
			return err
		case <-time.After(1 * time.Second):
			g.Log().Debug(ctx, "try again...")
		}
	}
}

func (m *Balance) exec(ctx context.Context, fc []func(ctx context.Context, tx gdb.TX) error) (err error) {
	m.Amount = math.Abs(m.Amount)
	if m.Amount == 0 {
		return fmt.Errorf("error amount")
	}
	var (
		user   entity.User
		wallet entity.Wallet
	)
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if err = tx.Model(dao.User.Table()).LockUpdate().Scan(&user, "id", m.Uid); err != nil {
			return err
		}
		if err = tx.Model(dao.Wallet.Table()).Scan(&wallet, "uid", m.Uid); err != nil {
			return err
		}
		if user.Id == 0 || wallet.Uid == 0 {
			return consts.ErrAccountNotFound
		}
		transCode, err := get.TransactionTypeFromCache(ctx, m.TransCode)
		if err != nil {
			return err
		}
		log := entity.WalletChangeLog{}
		uuid := xuuid.GenSnowflakeUUID()
		log.TransId = uuid.Int64()
		log.OrderNoRelation = m.OrderNoRelation
		log.Uid = user.Id
		log.Pid = user.Pid
		log.Account = user.Account
		log.ParentPath = user.ParentPath
		log.BalanceBefore = wallet.Balance
		log.Title = transCode.Title
		log.Note = m.Note
		log.TransCode = transCode.Code
		if transCode.Code < 0 {
			if wallet.Balance < m.Amount {
				return fmt.Errorf("insufficient balance")
			}
			log.Amount = -m.Amount
			log.BalanceAfter = wallet.Balance - m.Amount
			if log.BalanceAfter < 0 {
				return fmt.Errorf("error amount")
			}
		} else {
			log.Amount = m.Amount
			log.BalanceAfter = wallet.Balance + m.Amount
		}

		switch m.TransCode {
		case Bet:
			wallet.TotalBet += m.Amount
			wallet.TotalProfit -= m.Amount
		case Deposit:
			wallet.TotalDeposit += m.Amount
		case DepositBank:
			wallet.TotalDeposit += m.Amount
		case Withdraw:
			wallet.TotalWithdraw += m.Amount
		case Profit:
			wallet.TotalProfit += m.Amount
		case Gift:
			wallet.TotalGift += m.Amount
		case Unfreeze:
			wallet.Freeze -= m.Amount
		case Freeze:
			wallet.Freeze += m.Amount
			//wallet.TotalFreeze += this.Amount
		}
		// update wallet
		last, err := tx.Model(dao.Wallet.Table()).Data(g.Map{
			"balance":        log.BalanceAfter,
			"freeze":         wallet.Freeze,
			"total_bet":      wallet.TotalBet,
			"total_deposit":  wallet.TotalDeposit,
			"total_profit":   wallet.TotalProfit,
			"total_withdraw": wallet.TotalWithdraw,
			"total_gift":     wallet.TotalGift,
			//"total_freeze":   wallet.TotalFreeze,
		}).Where("uid = ? and balance != ?", user.Id, log.BalanceAfter).Update()
		if err != nil {
			return err
		}
		affected, err := last.RowsAffected()
		if err != nil || affected == 0 {
			return err
		}

		// insert log
		if _, err = tx.Model(dao.WalletChangeLog.Table()).Insert(&log); err != nil {
			return err
		}
		// other action
		if len(fc) > 0 {
			if err = fc[0](ctx, tx); err != nil {
				return err
			}
		}
		putReport := report.Report{
			Amount:    m.Amount,
			TransType: transCode,
			User:      user,
		}
		go putReport.Exec()
		return nil
	})
}
