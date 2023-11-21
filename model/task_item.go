package model

type TaskItem struct {
	Type    int
	Uid     int64
	Account string
	Data    interface{}
}
