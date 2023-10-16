package ws

import (
	"context"
	"game/app/gateway/internal/svc/user_svc"
	"game/app/user/api/user/user"
	"game/consts/ws_consts"
	"game/model"
)

func SendWallet(ctx context.Context, client *Client) {
	wallet, err := user_svc.Service.Wallet(ctx, &user.WalletRequest{
		Uid: client.UserInfo.Uid,
	})
	msg := model.WsMessage{
		Event: model.WrapEventResponse(ws_consts.Wallet),
		Body:  model.WrapMessage(wallet, err),
	}
	client.WriteChn <- &msg
}
