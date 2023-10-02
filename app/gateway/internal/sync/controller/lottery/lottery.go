package lottery

import (
	"fmt"
	"game/model"
)

type OpenResultController struct {
}

func (OpenResultController) Controller(data []byte) (*model.WsMessage, error) {
	fmt.Println(data)
	message := model.WsMessage{}
	message.Event = "lottery_openresult"
	message.Data = data
	return &message, nil
}
