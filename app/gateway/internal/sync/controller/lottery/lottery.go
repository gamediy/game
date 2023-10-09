package lottery

import (
	"fmt"
	sync_controller "game/app/gateway/internal/sync/controller"
	"game/model"
)

func Init() {
	sync_controller.ControllerC["/lottery/open_result"] = OpenResultController{}
}

type OpenResultController struct {
}

func (OpenResultController) Controller(data []byte) (*model.WsMessage, error) {
	fmt.Println(data)
	message := model.WsMessage{}
	message.Event = "lottery_openresult"
	message.Body = model.WrapMessage(data, nil)
	return &message, nil
}
