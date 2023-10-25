package user

import (
	"context"
	"game/app/gateway/internal/svc/user_svc"
	"game/app/gateway/internal/ws"
	"game/consts/event/user_event/deposit_event"
	"game/model"
	"github.com/gogf/gf/v2/frame/g"
)

func depositAmountItems(ctx context.Context, wsclient *ws.Client, query g.Map) (*model.WsMessage, error) {
	items, err := user_svc.Service.ListDepositAmountItems(ctx, wsclient.UserInfo.Uid)
	if err != nil {
		return nil, err
	}

	return &model.WsMessage{
		Event: model.WrapEventResponse(deposit_event.DepositAmountItems),
		Body:  model.WrapMessage(items, nil),
	}, nil
}
