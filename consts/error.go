package consts

import "fmt"

var (
	ErrAccountNotFound = fmt.Errorf("account not found")
	ErrLoginPass       = fmt.Errorf("login pass error")
	ErrDataNotFound    = fmt.Errorf("data not found")
)
