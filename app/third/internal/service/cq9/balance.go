package cq9

import (
	"context"
	"game/pure/get"
)

func Balance(ctx context.Context, account string) (float64, error) {
	w, err := get.WalletByAccount(ctx, account)
	if err != nil {
		return 0, err
	}
	return w.Balance, nil
}
