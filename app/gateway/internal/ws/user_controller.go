package ws

import (
	"context"
	"fmt"
	"game/app/gateway/internal/svc/user_svc"
	"game/app/user/api/user/user"
	"game/consts/ws_consts"
	"game/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func UserControllerInit() {
	Ctrl[ws_consts.Login] = login
	Ctrl[ws_consts.Heartbeat] = heartbeat
	Ctrl[ws_consts.Wallet] = wallet
	Ctrl[ws_consts.DepositAmountItems] = depositAmountItems
	Ctrl[ws_consts.ListMailBox] = listMailBox
}

func listMailBox(ctx context.Context, wsclient *Client, query g.Map) (*model.WsMessage, error) {
	req := user.ListMailBoxReq{}
	if query == nil {
		query = g.Map{"size": 1, "page": "10"}
	}
	req.Size = gconv.Int64(query["size"])
	req.Page = gconv.Int64(query["page"])
	req.Read = gconv.String(query["read"])
	req.Receiver = gconv.String(query["receiver"])
	res, err := user_svc.Service.ListMailBox(ctx, &req)
	if err != nil {
		return nil, err
	}

	return &model.WsMessage{
		Event: model.WrapEventResponse(ws_consts.ListMailBox),
		Body:  model.WrapMessage(res, nil),
	}, nil
}

func depositAmountItems(ctx context.Context, wsclient *Client, query g.Map) (*model.WsMessage, error) {
	items, err := user_svc.Service.ListDepositAmountItems(ctx)
	if err != nil {
		return nil, err
	}

	return &model.WsMessage{
		Event: model.WrapEventResponse(ws_consts.DepositAmountItems),
		Body:  model.WrapMessage(items, nil),
	}, nil
}

// 登录
func login(ctx context.Context, wsclient *Client, query g.Map) (*model.WsMessage, error) {

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

func heartbeat(ctx context.Context, client *Client, query g.Map) (*model.WsMessage, error) {

	fmt.Println("心跳设置")
	client.Heartbeat()
	return &model.WsMessage{
		Event: ws_consts.Heartbeat,
		Body:  model.WrapMessage("ping", nil),
	}, nil

}

// 钱包
func wallet(ctx context.Context, client *Client, query g.Map) (*model.WsMessage, error) {

	wallet, err := user_svc.Service.Wallet(ctx, &user.WalletRequest{
		Uid: client.UserInfo.Uid,
	})
	return &model.WsMessage{
		Event: model.WrapEventResponse(ws_consts.Wallet),
		Body:  model.WrapMessage(wallet, err),
	}, nil
}
