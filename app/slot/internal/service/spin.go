package service

import (
	"context"
	"fmt"
	"game/app/slot/api/slot/slot"
	"game/app/slot/internal/logic"
	"game/core/wallet"
	"game/db/dao"
	"game/db/model/entity"
	"game/utility/utils/xuuid"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"google.golang.org/grpc/status"
	"math/rand"
)

func Spin(ctx context.Context, request *slot.SpinReq) (res *slot.SpinRes, err error) {

	res = &slot.SpinRes{}
	game := entity.Game{}
	dao.Game.Ctx(ctx).Where("code", request.GameCode).Scan(&game)
	if game.Code == 0 {
		return res, status.Errorf(1, "没有此游戏")
	}
	play := []entity.Play{}
	dao.Play.Ctx(ctx).Where("play_type", game.PlayType).OrderAsc("probabilities").Scan(&play)
	if len(play) == 0 {
		return nil, status.Errorf(1, "没有玩法")
	}
	randNum := rand.Intn(game.RandomRange)
	var playOne entity.Play
	for _, e := range play {
		if randNum < gconv.Int(e.Probabilities) {
			playOne = e
			break
		}
	}
	if playOne.Code == 0 {
		playOne = play[len(play)-1]
	}
	balance := wallet.Balance{}
	balance.Uid = request.Uid
	balance.TransCode = wallet.Bet
	balance.Amount = request.Amount
	balance.OrderNoRelation = xuuid.GenSnowflakeUUID().Int64()
	balance.Note = fmt.Sprintf("Spin %d", game.Code)
	err = balance.Update(ctx, func(ctx context.Context, tx gdb.TX) error {
		f, ok := logic.GameSpin[int32(game.PlayType)]
		if ok {
			spin, bet, err := f(ctx, playOne, game)
			if err != nil {
				return err
			}

			res.OrderNo = balance.OrderNoRelation
			res.OpenResult = spin
			order := entity.SlotOrder{}
			order.Uid = request.Uid
			order.OrderNo = balance.OrderNoRelation
			order.Amount = float64(request.Amount)
			order.Odds = playOne.Odds
			order.BetResult = bet
			order.Profit = (playOne.Odds * request.Amount) - request.Amount
			order.Status = 1
			user := entity.User{}
			dao.User.Ctx(ctx).Where("id", request.Uid).Scan(&user)
			order.Account = user.Account
			order.ParentPath = user.ParentPath
			order.Pid = user.Pid
			order.PlayCode = playOne.Code
			order.GameCode = game.Code

			_, err = tx.Insert(dao.SlotOrder.Table(), order)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return res, status.Errorf(1, err.Error())
	}

	return res, nil

}

func CheckWon(ctx context.Context, req *slot.CheckWonReq) (res *slot.CheckWonRes, err error) {
	order := entity.SlotOrder{}
	dao.SlotOrder.Ctx(ctx).Where("order_no", req.OrderNo).Scan(&order)
	if order.OrderNo == 0 {
		return nil, status.Errorf(1, "没有此订单")
	}
	if order.Status != 1 {
		return nil, status.Errorf(1, "订单已结算")
	}
	res = &slot.CheckWonRes{}
	res.Profit = order.Profit
	order.Status = 3
	if res.Profit > 0 {
		order.Status = 2

		balance := wallet.Balance{}
		balance.Uid = req.Uid
		balance.TransCode = wallet.Profit
		balance.Note = "Spin won"
		balance.Amount = order.Profit + order.Amount
		err := balance.Update(ctx, func(ctx context.Context, tx gdb.TX) error {
			_, err = tx.Ctx(ctx).Update(dao.SlotOrder.Table(), &order, "order_no", order.OrderNo)
			return err
		})
		if err != nil {
			return nil, status.Errorf(1, err.Error())
		}
	} else {
		_, err = g.DB().Ctx(ctx).Update(ctx, dao.SlotOrder.Table(), &order, "order_no", order.OrderNo)
	}
	res.Status = int32(order.Status)
	return res, nil
}
