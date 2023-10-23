package cq9

import (
	"context"
	"fmt"
	"game/app/third/internal/logic"
)

type CQ9 struct {
}

func init() {
	logic.AccountSvc["CQ9"] = CQ9{}
	logic.BalanceSvc["CQ9"] = CQ9{}
}
func (CQ9) CheckAccount(ctx context.Context, account string) (data interface{}, err error) {
	fmt.Println("CQ9 account")
	return data, err
}

func (CQ9) GetBalance(ctx context.Context, account string, gameCode string) (data interface{}, err error) {
	fmt.Println("CQ9 balance")
	return nil, err
}
