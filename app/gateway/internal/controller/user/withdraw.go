package user

import (
	"context"
	"game/app/gateway/internal/svc/user_svc"
	"game/app/gateway/internal/ws"
	"game/app/user/api/user/withdraw"
	"game/consts/event/user_event/withdraw_event"
	"game/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func payPassStatus(ctx context.Context, wsclient *ws.Client, query g.Map) (*model.WsMessage, error) {
	status, err := user_svc.Service.PayPassStatus(ctx, wsclient.UserInfo.Uid)
	return &model.WsMessage{
			Event: model.WrapEventResponse(withdraw_event.PayPassStatus),
			Body:  model.WrapMessage(status, err)},
		nil
}

func setPayPass(ctx context.Context, wsclient *ws.Client, query g.Map) (*model.WsMessage, error) {
	res, err := user_svc.Service.SetPayPass(ctx, &withdraw.SetPayPassReq{
		Uid:  wsclient.UserInfo.Uid,
		Pass: gconv.String(query["pass"]),
	})
	return &model.WsMessage{
			Event: model.WrapEventResponse(withdraw_event.SetPayPass),
			Body:  model.WrapMessage(res, err)},
		nil
}

func bindWithdrawAccount(ctx context.Context, wsclient *ws.Client, query g.Map) (*model.WsMessage, error) {
	res, err := user_svc.Service.BindWithdrawAccount(ctx, &withdraw.BindWithdrawAccountReq{
		Uid:     wsclient.UserInfo.Uid,
		BankId:  gconv.Int64(query["bankId"]),
		Address: gconv.String(query["address"]),
		Title:   gconv.String(query["title"]),
		Pass:    gconv.String(query["pass"]),
	})
	return &model.WsMessage{
			Event: model.WrapEventResponse(withdraw_event.BindWithdrawAccount),
			Body:  model.WrapMessage(res, err)},
		nil
}

func delWithdrawAccountById(ctx context.Context, wsclient *ws.Client, query g.Map) (*model.WsMessage, error) {
	res, err := user_svc.Service.DelWithdrawAccount(ctx, &withdraw.DelWithdrawAccountReq{
		Uid:     wsclient.UserInfo.Uid,
		Id:      gconv.Int64(query["id"]),
		PayPass: gconv.String(query["payPass"]),
	})
	return &model.WsMessage{
			Event: model.WrapEventResponse(withdraw_event.DelWithdrawAccount),
			Body:  model.WrapMessage(res, err)},
		nil
}

func listWithdrawAccount(ctx context.Context, wsclient *ws.Client, query g.Map) (*model.WsMessage, error) {
	res, err := user_svc.Service.ListWithdrawAccount(ctx, &withdraw.ListWithdrawAccountReq{
		Uid:  wsclient.UserInfo.Uid,
		Page: gconv.Int64(query["page"]),
		Size: gconv.Int64(query["size"]),
	})
	return &model.WsMessage{
		Event: model.WrapEventResponse(withdraw_event.ListWithdrawAccount),
		Body:  model.WrapMessage(res, err),
	}, nil
}

func createWithdraw(ctx context.Context, wsclient *ws.Client, query g.Map) (*model.WsMessage, error) {
	res, err := user_svc.Service.CreateWithdraw(ctx, &withdraw.CreateWithdrawReq{
		AmountItemId:      gconv.Int64(query["amountItemId"]),
		Amount:            gconv.Float64(query["amount"]),
		WithdrawAccountId: gconv.Int64(query["withdrawAccountId"]),
		Lang:              wsclient.UserInfo.Lang,
		Uid:               wsclient.UserInfo.Uid,
	})
	return &model.WsMessage{
		Event: model.WrapEventResponse(withdraw_event.CreateWithdraw),
		Body:  model.WrapMessage(res, err),
	}, nil
}

func listWithdrawMethod(ctx context.Context, wsclient *ws.Client, query g.Map) (*model.WsMessage, error) {
	res, err := user_svc.Service.ListWithdrawMethod(ctx, &withdraw.ListWithdrawMethodReq{
		Uid:  wsclient.UserInfo.Uid,
		Lang: wsclient.UserInfo.Lang,
	})
	return &model.WsMessage{
		Event: model.WrapEventResponse(withdraw_event.ListWithdrawMethod),
		Body:  model.WrapMessage(res, err),
	}, nil
}

func listWithdraw(ctx context.Context, wsclient *ws.Client, query g.Map) (*model.WsMessage, error) {
	res, err := user_svc.Service.ListWithdraw(ctx, &withdraw.ListWithdrawReq{
		Uid:  wsclient.UserInfo.Uid,
		Page: gconv.Int64(query["page"]),
		Size: gconv.Int64(query["size"]),
	})
	return &model.WsMessage{
		Event: model.WrapEventResponse(withdraw_event.ListWithdraw),
		Body:  model.WrapMessage(res, err),
	}, nil
}
