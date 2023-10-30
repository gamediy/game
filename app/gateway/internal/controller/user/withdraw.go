package user

import (
	"context"
	"game/app/gateway/internal/svc/user_svc"
	"game/app/gateway/internal/ws"
	"game/consts/event/user_event/withdraw_event"
	"game/model"
	"github.com/gogf/gf/v2/frame/g"
)

func payPassStatus(ctx context.Context, wsclient *ws.Client, query g.Map) (*model.WsMessage, error) {
	status, err := user_svc.Service.PayPassStatus(ctx, wsclient.UserInfo.Uid)
	return &model.WsMessage{
		Event: model.WrapEventResponse(withdraw_event.PayPassStatus),
		Body:  model.WrapMessage(status, err),
	}, nil
}
