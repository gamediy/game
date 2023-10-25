package user

import (
	"context"
	"game/app/gateway/internal/svc/user_svc"
	"game/app/gateway/internal/ws"
	"game/app/user/api/user/user"
	"game/consts/event/user_event/wallet_event"
	"game/model"
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
