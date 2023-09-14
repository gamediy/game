package sync

import (
	"context"
	"fmt"
	"game/app/gateway/internal/sync/controller"
	"game/app/gateway/internal/sync/controller/lottery"
	"game/app/gateway/internal/ws"
	"game/utility/xetcd"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	lottery_controller map[string]sync_controller.Controller
)

func init() {
	lottery_controller = make(map[string]sync_controller.Controller, 10)
	lottery_controller["/lottery/openresult"] = lottery.OpenResultController{}
}
func Router(ctx context.Context) {

	fmt.Println("start sync watch")
	xetcd.Watch(ctx, "/lottery/", true, func(response clientv3.WatchResponse) {
		for _, event := range response.Events {
			s, ok := lottery_controller[string(event.Kv.Key)]
			if !ok {
				continue
			}
			message, err := s.Controller(event.Kv.Value)
			if err == nil {
				ws.SendToAll(message)
			}

		}
		fmt.Println(response)
	})
	xetcd.Watch(ctx, "/solt/", true, func(response clientv3.WatchResponse) {
		for _, event := range response.Events {
			fmt.Println(string(event.Kv.Key), string(event.Kv.Value))
		}
		fmt.Println(response)
	})
	xetcd.Watch(ctx, "/user/", true, func(response clientv3.WatchResponse) {
		fmt.Println(response)
	})
}
