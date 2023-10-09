package ws

import (
	"context"
	"fmt"
	"game/app/gateway/internal/svc/user_svc"
	"game/app/user/api/user/user"
	"game/model"
)

const (
	Error     = "/error"
	Heartbeat = "/heartbeat"

	Login  = "/user/login"
	Join   = "/user/join"
	Quit   = "/user/quit"
	Wallet = "/user/wallet"

	Enter = "/game/enter"
)

func UserControllerInit() {

	ControllerC[Login] = LoginController{}
	ControllerC[Heartbeat] = HeartbeatController{}
	ControllerC[Wallet] = WalletController{}
}

// 登录
type LoginController struct {
}

func (LoginController) Controller(ctx context.Context, wsclient *Client, msg *model.WsMessage) (*model.WsMessage, error) {

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
		Event: model.WrapEventResponse(Login),
		Body:  model.WrapMessage(data, nil),
	}, nil
}

// 心跳
type HeartbeatController struct {
}

func (HeartbeatController) Controller(ctx context.Context, client *Client, msg *model.WsMessage) (*model.WsMessage, error) {

	fmt.Println("心跳设置")
	client.Heartbeat()
	return &model.WsMessage{
		Event: Heartbeat,
		Body:  model.WrapMessage("ping", nil),
	}, nil

}

type RegisterController struct {
}
type WalletController struct {
}

func (WalletController) Controller(ctx context.Context, client *Client, msg *model.WsMessage) (*model.WsMessage, error) {

	wallet, err := user_svc.Service.Wallet(ctx, &user.WalletRequest{
		Uid: client.UserInfo.Uid,
	})
	return &model.WsMessage{
		Event: model.WrapEventResponse(Wallet),
		Body:  model.WrapMessage(wallet, err),
	}, nil
}
