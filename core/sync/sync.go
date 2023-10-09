package sync

import (
	"context"
	"encoding/json"
	"game/utility/utils/xetcd"
)

type SyncInfo struct {
	WsMessage interface{}
}

func (s *SyncInfo) Put(ctx context.Context) {
	marshal, _ := json.Marshal(s)
	xetcd.Client.Put(ctx, "/sync/", string(marshal))
}
