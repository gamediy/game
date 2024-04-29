package model

type TaskItem struct {
	Events  string
	Uid     int64
	Account string
	Data    interface{}
}
