package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"game/model"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gorilla/websocket"
	"runtime/debug"
	"time"
)

type Client struct {
	Addr     string
	Sid      string
	Socket   *websocket.Conn
	WriteChn chan *model.WsMessage

	UserInfo      *model.UserInfo
	LoginTime     *gtime.Time
	HeartbeatTime *gtime.Time
	SendClose     bool            // 发送是否关闭
	tags          garray.StrArray // 标签
}

var (
	HeartbeatTimeout = 300
)

// NewClient 初始化
func NewClient(addr string, userInfo *model.UserInfo, socket *websocket.Conn) (client *Client) {

	client = &Client{
		Addr:          addr,
		UserInfo:      userInfo,
		Socket:        socket,
		WriteChn:      make(chan *model.WsMessage, 100),
		SendClose:     false,
		LoginTime:     gtime.Now(),
		HeartbeatTime: gtime.Now(),
	}
	client.Heartbeat()
	return
}
func (c *Client) Read(ctx context.Context, ctrl func(ctx context.Context, msg *model.WsMessage, wsclient *Client, query g.Map)) {
	defer func() {

		if r := recover(); r != nil {
			fmt.Println("read stop", string(debug.Stack()), r)
		}
	}()

	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			c.Close()
			return
		}
		// 处理程序
		fmt.Println(string(message))
		msg := model.WsMessage{}
		err = json.Unmarshal(message, &msg)
		if err != nil {
			return
		}
		m2 := gconv.Map(msg.Query)
		ctrl(ctx, &msg, c, m2)
	}
}

// 向客户端写数据
func (c *Client) Write(ctx context.Context) {
	defer func() {
		_ = c.Close
		if r := recover(); r != nil {
			fmt.Println("write stop", string(debug.Stack()), r)
		}
	}()

	for {
		select {
		case message := <-c.WriteChn:
			if message != nil {
				fmt.Println("uid ", c.UserInfo.Uid)
				fmt.Println("writer msg ", message)
				err := c.Socket.WriteJSON(message)
				if err != nil {
					fmt.Println("write msg error", err)
					c.Close()
					return
				}

			}

		}
	}
}

// SendMsg 发送数据
func (c *Client) WriterMsg(ctx context.Context, msg *model.WsMessage) {
	if c == nil || c.SendClose {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("SendMsg stop:", r, string(debug.Stack()))
		}
	}()
	fmt.Println("writer msg ", msg)
	c.WriteChn <- msg
}

// Heartbeat 心跳更新
func (c *Client) Heartbeat() {
	c.HeartbeatTime = gtime.Now().Add(time.Second * time.Duration(HeartbeatTimeout))
	return
}

// IsHeartbeatTimeout 心跳是否超时
func (c *Client) IsHeartbeatTimeout(currentTime *gtime.Time) (timeout bool) {

	fmt.Println(c.HeartbeatTime, currentTime)
	if c.HeartbeatTime.Before(currentTime) {
		fmt.Println("timeout true")
		timeout = true
	}
	return
}

// 关闭客户端
func (c *Client) Close() {

	fmt.Println("客户端关闭：", c.UserInfo.Uid)
	if c.SendClose {
		return
	}

	c.Socket.Close()
	Manager.DelUsers(c)
	fmt.Println("在线：", Manager.GetUsersLen())
	c.SendClose = true
	close(c.WriteChn)

}

func (c *Client) SetTag(tag string) {
	c.tags.Set(0, tag)
}
