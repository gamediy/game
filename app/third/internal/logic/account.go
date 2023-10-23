package logic

import (
	"context"
)

var (
	AccountSvc = make(map[string]Account)
)

type Account interface {
	CheckAccount(ctx context.Context, account string) (data interface{}, err error)
}
