package sync

import (
	"context"
	"encoding/json"
	"fmt"
	"game/app/gateway/internal/ws"
	"game/model"
	"game/utility/utils/xetcd"
	"github.com/gogf/gf/v2/frame/g"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func Router(ctx context.Context) {
	fmt.Println("start sync watch")
	go xetcd.Watch(ctx, "/sync/", true, func(response clientv3.WatchResponse) {
		for _, event := range response.Events {
			model := model.WsMessage{}
			err := json.Unmarshal(event.Kv.Value, &model)
			if err != nil {
				g.Log().Error(ctx, err)
				continue
			}
			ws.Send(&model)
		}
		fmt.Println(response)
	})
}
