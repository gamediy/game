package model

type TaskItem struct {
	TypeCode string
	Uid      int64
	Account  string
	Data     interface{}
}
