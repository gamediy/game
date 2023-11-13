package cq9

import (
	"context"
	"game/pure/get"
)

func CheckPlayer(ctx context.Context, account string) bool {
	_, err := get.User(ctx, account)
	if err != nil {
		return false
	}

	return true
}
