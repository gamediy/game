package ws

import (
	"context"
	"fmt"
	"game/app/gateway/internal/svc/user_svc"
	"game/app/user/api/user/user"
	"game/consts/ws_consts"
	"game/model"
)

var (
	Ctrl map[string]func(ctx context.Context, wsclient *Client, msg *model.WsMessage) (*model.WsMessage, error)
)

func UserControllerInit() {
	Ctrl[ws_consts.Login] = login
	Ctrl[ws_consts.Heartbeat] = heartbeat
	Ctrl[ws_consts.Wallet] = wallet
}

// 登录
func login(ctx context.Context, wsclient *Client, msg *model.WsMessage) (*model.WsMessage, error) {

	Manager.AddUsers(wsclient)
	wallet, _ := user_svc.Service.Wallet(ctx, &user.WalletRequest{
		Uid: wsclient.UserInfo.Uid,
	})
	data := struct {
		UserInfo *model.UserInfo   `json:"userInfo"`
		Wallet   *user.WalletReply `json:"wallet"`
	}{
		UserInfo: wsclient.UserInfo,
		Wallet:   wallet,
	}
	return &model.WsMessage{
		Event: model.WrapEventResponse(ws_consts.Login),
		Body:  model.WrapMessage(data, nil),
	}, nil
}

// 心跳

func heartbeat(ctx context.Context, client *Client, msg *model.WsMessage) (*model.WsMessage, error) {

	fmt.Println("心跳设置")
	client.Heartbeat()
	return &model.WsMessage{
		Event: ws_consts.Heartbeat,
		Body:  model.WrapMessage("ping", nil),
	}, nil

}

// 钱包
func wallet(ctx context.Context, client *Client, msg *model.WsMessage) (*model.WsMessage, error) {

	wallet, err := user_svc.Service.Wallet(ctx, &user.WalletRequest{
		Uid: client.UserInfo.Uid,
	})
	return &model.WsMessage{
		Event: model.WrapEventResponse(ws_consts.Wallet),
		Body:  model.WrapMessage(wallet, err),
	}, nil
}
