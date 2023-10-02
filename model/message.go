package model

type WsMessage struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
	Tag   string      `json:"tag"`
	Uid   int         `json:"_"`
}

type HttpMessage struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func WrapHttpMessage(data interface{}, err error) *HttpMessage {
	message := HttpMessage{}
	if err != nil {
		message.Code = 50
		message.Msg = err.Error()
	}
	message.Data = data
	return &message
}
