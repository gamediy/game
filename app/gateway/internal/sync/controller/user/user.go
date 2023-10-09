package user

import (
	"fmt"
	sync_controller "game/app/gateway/internal/sync/controller"
	"game/model"
)

func Init() {
	sync_controller.ControllerC["/user/wallet"] = WalletController{}
}

type WalletController struct {
}

func (WalletController) Controller(data []byte) (*model.WsMessage, error) {
	fmt.Println(data)
	message := model.WsMessage{}
	message.Event = "lottery_openresult"
	message.Body = model.WrapMessage(data, nil)
	return &message, nil
}
