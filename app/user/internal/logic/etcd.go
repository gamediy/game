package logic

import (
	"context"
	"fmt"
	"game/utility/utils/xetcd"
)

func Test123(ctx context.Context) {
	put, err := xetcd.Client.Put(ctx, "/game/1000/openresult", "game open")

	if err == nil {
		fmt.Println(put)
	}
}
