package ws

import (
	"fmt"
	"game/model"
)

// 登录
type LoginController struct {
}

// 心跳
type HeartbeatController struct {
}

func (LoginController) Controller(client *Client, msg *model.Message) (*model.Message, error) {

	Manager.AddUsers(client)
	return &model.Message{
		Event: Login,
		Data:  "login success",
	}, nil
}
func (HeartbeatController) Controller(client *Client, msg *model.Message) (*model.Message, error) {

	fmt.Println("心跳设置")
	client.Heartbeat()
	return &model.Message{
		Event: Heartbeat,
		Data:  "ping",
	}, nil

}
