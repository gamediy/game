package user

import (
	"context"
	"fmt"
	"game/app/gateway/internal/controller"
	"game/app/gateway/internal/svc/user_svc"
	"game/app/gateway/internal/ws"
	"game/app/user/api/user/user"
	"game/consts/event/user_event"
	"game/consts/event/user_event/deposit_event"
	"game/consts/event/user_event/mailbox_event"
	"game/consts/event/user_event/sys_event"
	"game/consts/event/user_event/wallet_event"
	"game/consts/event/user_event/withdraw_event"
	"github.com/gogf/gf/v2/util/gconv"

	"game/model"
	"github.com/gogf/gf/v2/frame/g"
)

func UserControllerInit() {
	controller.Ctrl[user_event.Login] = login
	controller.Ctrl[user_event.Heartbeat] = heartbeat
	controller.Ctrl[user_event.UpdateLoginPass] = updateLoginPass
	controller.Ctrl[wallet_event.Wallet] = walletInfo
	controller.Ctrl[wallet_event.ListChangeLog] = listChangeLog
	controller.Ctrl[wallet_event.ListTransType] = listTransType
	controller.Ctrl[deposit_event.DepositAmountItems] = depositAmountItems
	controller.Ctrl[deposit_event.CreateDeposit] = createDeposit
	controller.Ctrl[deposit_event.ListDeposit] = listDeposit
	controller.Ctrl[mailbox_event.ListMailBox] = listMailBox
	controller.Ctrl[mailbox_event.MailBoxTotal] = countMailBoxTotal
	controller.Ctrl[withdraw_event.PayPassStatus] = payPassStatus
	controller.Ctrl[withdraw_event.SetPayPass] = setPayPass
	controller.Ctrl[withdraw_event.BindWithdrawAccount] = bindWithdrawAccount
	controller.Ctrl[withdraw_event.DelWithdrawAccount] = delWithdrawAccountById
	controller.Ctrl[withdraw_event.ListWithdrawAccount] = listWithdrawAccount
	controller.Ctrl[withdraw_event.CreateWithdraw] = createWithdraw
	controller.Ctrl[withdraw_event.ListWithdrawMethod] = listWithdrawMethod
	controller.Ctrl[withdraw_event.ListWithdraw] = listWithdraw
	controller.Ctrl[withdraw_event.ListPublicWithdraw] = listPublicWithdraw
	controller.Ctrl[sys_event.GetAnnouncement] = getAnnouncement
	controller.Ctrl[sys_event.SendSmsCode] = sendSmsCode
}

func sendSmsCode(ctx context.Context, wsclient *ws.Client, query g.Map) (*model.WsMessage, error) {
	res, err := user_svc.Service.SendSmsCode(ctx, &user.SendMsgCodeReq{
		Phone: gconv.String(query["phone"]),
	})
	return &model.WsMessage{
		Event: model.WrapEventResponse(sys_event.SendSmsCode),
		Body:  model.WrapMessage(res, err),
	}, nil
}

func updateLoginPass(ctx context.Context, wsclient *ws.Client, query g.Map) (*model.WsMessage, error) {
	res, err := user_svc.Service.UpdateLoginPass(ctx, &user.UpdateLoginPassReq{
		Uid:     wsclient.UserInfo.Uid,
		OldPass: gconv.String(query["oldPass"]),
		NewPass: gconv.String(query["newPass"]),
	})
	return &model.WsMessage{
		Event: model.WrapEventResponse(user_event.UpdateLoginPass),
		Body:  model.WrapMessage(res, err),
	}, nil
}

// 登录
func login(ctx context.Context, wsclient *ws.Client, query g.Map) (*model.WsMessage, error) {

	ws.Manager.AddUsers(wsclient)
	wallet, err := user_svc.Service.Wallet(ctx, &user.WalletRequest{
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
		Body:  model.WrapMessage(data, err),
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
