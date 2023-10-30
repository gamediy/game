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
		Body:  model.WrapMessage(status, err),
	}, nil
}

func setPayPass(ctx context.Context, wsclient *ws.Client, query g.Map) (*model.WsMessage, error) {
	res, err := user_svc.Service.SetPayPass(ctx, &withdraw.SetPayPassReq{
		Uid:  wsclient.UserInfo.Uid,
		Pass: gconv.String(query["pass"]),
	})
	return &model.WsMessage{
		Event: model.WrapEventResponse(withdraw_event.SetPayPass),
		Body:  model.WrapMessage(res, err),
	}, nil
}
