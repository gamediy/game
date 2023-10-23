package cq9

import (
	"context"
	"game/app/third/internal/logic"
	"testing"
)

func TestCQ9_CheckAccount(t *testing.T) {
	logic.AccountSvc["CQ9"].CheckAccount(context.TODO(), "account")
}
func TestCQ9_GetBalance(t *testing.T) {
	logic.BalanceSvc["CQ9"].GetBalance(context.TODO(), "account", "112")
}
