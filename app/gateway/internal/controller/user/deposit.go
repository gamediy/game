package user

import (
	"context"
	"game/app/gateway/internal/svc/user_svc"
	"game/app/gateway/internal/ws"
	"game/app/user/api/user/deposit"
	"game/consts/event/user_event/deposit_event"
	"game/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func depositAmountItems(ctx context.Context, wsclient *ws.Client, query g.Map) (*model.WsMessage, error) {
	items, err := user_svc.Service.ListDepositAmountItems(ctx, wsclient.UserInfo.Uid)
	return &model.WsMessage{
		Event: model.WrapEventResponse(deposit_event.DepositAmountItems),
		Body:  model.WrapMessage(items, err),
	}, nil
}

func createDeposit(ctx context.Context, wsclient *ws.Client, query g.Map) (*model.WsMessage, error) {
	in := &deposit.CreateDepositReq{}
	in.PayId = gconv.Int64(query["payId"])
	in.Amount = gconv.Float64(query["amount"])
	in.TransferOrderNo = gconv.String(query["transferOrderNo"])
	in.TransferImg = gconv.String(query["transferImg"])
	in.Uid = wsclient.UserInfo.Uid
	in.Lang = wsclient.UserInfo.Lang
	orderNo, err := user_svc.Service.CreateDeposit(ctx, in)
	return &model.WsMessage{
		Event: model.WrapEventResponse(deposit_event.CreateDeposit),
		Body:  model.WrapMessage(orderNo, err),
	}, nil
}
