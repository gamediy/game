package sync

import (
	"context"
	"encoding/json"
	"game/model"
	"game/utility/utils/xetcd"
)

func Put(ctx context.Context, msg model.WsMessage) {
	marshal, _ := json.Marshal(msg)
	xetcd.Client.Put(ctx, "/sync/", string(marshal))
}
