package user

import (
	"context"
	"game/app/gateway/internal/svc/user_svc"
	"game/app/gateway/internal/ws"
	"game/app/user/api/user/user"
	"game/app/user/api/user/wallet"
	"game/consts/event/user_event/wallet_event"
	"game/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func SendWallet(ctx context.Context, client *ws.Client) {
	res, err := user_svc.Service.Wallet(ctx, &user.WalletRequest{
		Uid: client.UserInfo.Uid,
	})
	msg := model.WsMessage{
		Event: model.WrapEventResponse(wallet_event.Wallet),
		Body:  model.WrapMessage(res, err),
	}
	client.WriteChn <- &msg
}

// 钱包
func walletInfo(ctx context.Context, client *ws.Client, query g.Map) (*model.WsMessage, error) {
	res, err := user_svc.Service.Wallet(ctx, &user.WalletRequest{
		Uid: client.UserInfo.Uid,
	})
	return &model.WsMessage{
		Event: model.WrapEventResponse(wallet_event.Wallet),
		Body:  model.WrapMessage(res, err),
	}, nil
}

func listChangeLog(ctx context.Context, wsclient *ws.Client, query g.Map) (*model.WsMessage, error) {
	res, err := user_svc.Service.ListChangeLog(ctx, &wallet.ListChangeLogReq{
		Uid:       wsclient.UserInfo.Uid,
		Page:      gconv.Int64(query["page"]),
		Size:      gconv.Int64(query["size"]),
		TransCode: gconv.String(query["transCode"]),
	})
	return &model.WsMessage{
		Event: model.WrapEventResponse(wallet_event.ListChangeLog),
		Body:  model.WrapMessage(res, err),
	}, nil
}

func listTransType(ctx context.Context, wsclient *ws.Client, query g.Map) (*model.WsMessage, error) {
	res, err := user_svc.Service.ListTransType(ctx, &wallet.ListTransTypeReq{})
	return &model.WsMessage{
		Event: model.WrapEventResponse(wallet_event.ListTransType),
		Body:  model.WrapMessage(res, err),
	}, nil
}
