package model

type Message struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
	Tag   string      `json:"tag"`
	Uid   int         `json:"_"`
}
