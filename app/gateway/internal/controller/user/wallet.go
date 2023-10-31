package user

import (
	"context"
	"game/app/gateway/internal/svc/user_svc"
	"game/app/gateway/internal/ws"
	"game/app/user/api/user/user"
	wallet2 "game/app/user/api/user/wallet"
	"game/consts/event/user_event/wallet_event"
	"game/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
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

// 钱包
func wallet(ctx context.Context, client *ws.Client, query g.Map) (*model.WsMessage, error) {

	wallet, err := user_svc.Service.Wallet(ctx, &user.WalletRequest{
		Uid: client.UserInfo.Uid,
	})
	return &model.WsMessage{
		Event: model.WrapEventResponse(wallet_event.Wallet),
		Body:  model.WrapMessage(wallet, err),
	}, nil
}

func listChangeLog(ctx context.Context, wsclient *ws.Client, query g.Map) (*model.WsMessage, error) {
	res, err := user_svc.Service.ListChangeLog(ctx, &wallet2.ListChangeLogReq{
		Uid:  wsclient.UserInfo.Uid,
		Page: gconv.Int64(query["page"]),
		Size: gconv.Int64(query["size"]),
	})
	return &model.WsMessage{
		Event: model.WrapEventResponse(wallet_event.ListChangeLog),
		Body:  model.WrapMessage(res, err),
	}, nil
}
