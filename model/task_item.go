package model

type TaskItem struct {
	Type    int // 1 签到
	Uid     int64
	Account string
	Data    interface{}
}
