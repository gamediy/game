package model

type WsMessage struct {
	Event string      `json:"event"`
	Body  *Message    `json:"body"`
	Query interface{} `json:"query"`
	Tag   string      `json:"tag"`
	Uid   int64       `json:"_"`
}

func WrapEventResponse(string string) string {
	return string + "_response"
}

type Message struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func WrapMessage(data interface{}, err error) *Message {
	message := Message{}
	if err != nil {
		message.Code = 50
		message.Msg = err.Error()
	}
	message.Data = data
	return &message
}
