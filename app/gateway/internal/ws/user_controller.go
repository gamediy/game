package ws

import (
	"context"
	"fmt"
	"game/model"
)

// 登录
type LoginController struct {
}

func (LoginController) Controller(ctx context.Context, client *Client, msg *model.WsMessage) (*model.WsMessage, error) {

	Manager.AddUsers(client)
	return &model.WsMessage{
		Event: Login,
		Data:  "login success",
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
		Data:  "ping",
	}, nil

}

type RegisterController struct {
}
