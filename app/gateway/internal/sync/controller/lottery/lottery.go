package lottery

import (
	"fmt"
	"game/model"
)

type OpenResultController struct {
}

func (OpenResultController) Controller(data []byte) (*model.Message, error) {
	fmt.Println(data)
	message := model.Message{}
	message.Event = "lottery_openresult"
	message.Data = data
	return &message, nil
}
