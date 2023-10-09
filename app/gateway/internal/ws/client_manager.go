package ws

import (
	"context"
	"fmt"
	"game/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/os/gcron"
	"sync"
)

var (
	Manager = NewClientManager() // 管理者

)

func StartWebSocket(ctx context.Context) {
	g.Log().Info(ctx, "启动：WebSocket")
	go Manager.start(ctx)
	go Manager.ping(ctx)
}

type ClientManager struct {
	Clients    map[int64]*Client // 全部的连接
	Lock       sync.RWMutex      // 读写锁
	Unregister chan *Client      // 断开连接处理程序
	Login      chan *Client
	Broadcast  chan *model.WsMessage // 广播 向全部成员发送数据

	TagBroadcast  chan *model.WsMessage
	UserBroadcast chan *model.WsMessage
}

func NewClientManager() (clientManager *ClientManager) {
	clientManager = &ClientManager{
		Clients:       make(map[int64]*Client),
		Unregister:    make(chan *Client, 1000),
		Broadcast:     make(chan *model.WsMessage, 1000),
		TagBroadcast:  make(chan *model.WsMessage, 1000),
		UserBroadcast: make(chan *model.WsMessage, 1000),
	}
	return
}

func (manager *ClientManager) InClient(client *Client) (ok bool) {
	manager.Lock.RLock()
	defer manager.Lock.RUnlock()
	_, ok = manager.Clients[client.UserInfo.Uid]
	return
}

func (manager *ClientManager) GetClientsLen() (clientsLen int) {
	clientsLen = len(manager.Clients)
	return
}

func (manager *ClientManager) AddUsers(client *Client) {
	manager.Lock.Lock()
	defer manager.Lock.Unlock()
	manager.Clients[client.UserInfo.Uid] = client
	fmt.Println("add client uid ", client.UserInfo.Uid)
	fmt.Println("当前在线：", manager.GetClientsLen())
}

func (manager *ClientManager) DelUsers(client *Client) {
	manager.Lock.Lock()
	defer manager.Lock.Unlock()
	if _, ok := manager.Clients[client.UserInfo.Uid]; ok {
		delete(manager.Clients, client.UserInfo.Uid)
	}

}

// GetUsersLen 已登录用户数
func (manager *ClientManager) GetUsersLen() (userLen int) {
	userLen = len(manager.Clients)
	return
}

// EventLogin 用户登录事件
func (manager *ClientManager) EventLogin(client *Client) {
	if manager.InClient(client) {
		manager.AddUsers(client)
	}
}

// EventUnregister 用户断开连接事件
func (manager *ClientManager) EventUnregister(client *Client) {
	// 删除用户连接
	manager.DelUsers(client)

}

// ClearTimeoutConnections 定时清理超时连接
func (manager *ClientManager) clearTimeoutConnections() {

	time := gtime.Now()
	for _, client := range manager.Clients {
		if client.IsHeartbeatTimeout(time) {
			fmt.Println("心跳时间超时 关闭连接", client.Addr, client.UserInfo.Uid, client.LoginTime, client.HeartbeatTime)
			client.Close()
		}
	}
}

func (manager *ClientManager) ping(ctx context.Context) {

	gcron.Add(ctx, "*/30 * * * * *", func(ctx context.Context) {
		manager.clearTimeoutConnections()
	})

}

func (manager *ClientManager) start(ctx context.Context) {
	for {
		select {

		case login := <-manager.Login:
			// 用户登录
			manager.EventLogin(login)

		case conn := <-manager.Unregister:
			// 断开连接事件
			manager.EventUnregister(conn)
		case message := <-manager.Broadcast:
			fmt.Println("广播：", message)
			for _, conn := range manager.Clients {
				conn.WriterMsg(ctx, message)
			}
		case message := <-manager.TagBroadcast:
			// 标签广播事件
			for _, conn := range manager.Clients {
				if conn.tags.Contains(message.Tag) {
					conn.WriterMsg(ctx, message)
				}
			}
		case message := <-manager.UserBroadcast:
			// 用户广播事件

			for _, conn := range manager.Clients {
				if conn.UserInfo.Uid == message.Uid {
					conn.WriterMsg(ctx, message)
				}
			}

		}

	}
}

// SendToAll 发送全部客户端
func SendToAll(response *model.WsMessage) {
	Manager.Broadcast <- response
}

// SendToUser 发送单个用户
func SendToUser(response *model.WsMessage) {
	Manager.UserBroadcast <- response
}

// SendToTag 发送某个标签
func SendToTag(response *model.WsMessage) {
	Manager.TagBroadcast <- response
}
