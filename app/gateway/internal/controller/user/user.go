package user

import (
	"context"
	"fmt"
	"game/app/gateway/internal/controller"
	"game/app/gateway/internal/svc/user_svc"
	"game/app/gateway/internal/ws"
	"game/consts/event/user_event/deposit_event"
	"game/consts/event/user_event/withdraw_event"

	"game/app/user/api/user/user"
	"game/consts/event/user_event"
	"game/consts/event/user_event/mailbox_event"
	"game/consts/event/user_event/wallet_event"

	"game/model"
	"github.com/gogf/gf/v2/frame/g"
)

func UserControllerInit() {
	controller.Ctrl[user_event.Login] = login
	controller.Ctrl[user_event.Heartbeat] = heartbeat
	controller.Ctrl[wallet_event.Wallet] = wallet
	controller.Ctrl[deposit_event.DepositAmountItems] = depositAmountItems
	controller.Ctrl[deposit_event.CreateDeposit] = createDeposit
	controller.Ctrl[deposit_event.ListDeposit] = listDeposit
	controller.Ctrl[mailbox_event.ListMailBox] = listMailBox
	controller.Ctrl[mailbox_event.MailBoxTotal] = countMailBoxTotal
	controller.Ctrl[withdraw_event.PayPassStatus] = payPassStatus
}

// 登录
func login(ctx context.Context, wsclient *ws.Client, query g.Map) (*model.WsMessage, error) {

	ws.Manager.AddUsers(wsclient)
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
		Event: model.WrapEventResponse(user_event.Login),
		Body:  model.WrapMessage(data, nil),
	}, nil
}

// 心跳

func heartbeat(ctx context.Context, client *ws.Client, query g.Map) (*model.WsMessage, error) {

	fmt.Println("心跳设置")
	client.Heartbeat()
	return &model.WsMessage{
		Event: user_event.Heartbeat,
		Body:  model.WrapMessage("ping", nil),
	}, nil

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
