package user

import (
	"context"
	"game/app/gateway/internal/svc/user_svc"
	"game/app/gateway/internal/ws"
	"game/app/user/api/user/user"
	"game/consts/event/user_event/wallet_event"
	"game/model"
	"github.com/gogf/gf/v2/frame/g"
)

func SendWallet(ctx context.Context, client *ws.Client) {
	wallet, err := user_svc.Service.Wallet(ctx, &user.WalletRequest{
		Uid: client.UserInfo.Uid,
	})
	msg := model.WsMessage{
		Event: model.WrapEventResponse(wallet_event.Wallet),
		Body:  model.WrapMessage(wallet, err),
	}
	client.WriteChn <- &msg
}

func DepositAmountItems(ctx context.Context, wsclient *ws.Client, query g.Map) (*model.WsMessage, error) {
	items, err := user_svc.Service.ListDepositAmountItems(ctx)
	if err != nil {
		return nil, err
	}

	return &model.WsMessage{
		Event: model.WrapEventResponse(wallet_event.DepositAmountItems),
		Body:  model.WrapMessage(items, nil),
	}, nil
}
