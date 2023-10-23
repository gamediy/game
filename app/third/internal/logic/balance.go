package logic

import "context"

var (
	BalanceSvc = make(map[string]Balance)
)

type Balance interface {
	GetBalance(ctx context.Context, account string, gameCode string) (data interface{}, err error)
}
