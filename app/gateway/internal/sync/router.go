package sync

import (
	"context"
	"fmt"
	sync_controller "game/app/gateway/internal/sync/controller"
	"game/app/gateway/internal/sync/controller/lottery"
	"game/app/gateway/internal/ws"
	"game/utility/utils/xetcd"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func InitController() {
	sync_controller.ControllerC = make(map[string]sync_controller.Controller, 10)
	lottery.Init()
}
func Router(ctx context.Context) {
	InitController()
	fmt.Println("start sync watch")
	go xetcd.Watch(ctx, "/sync/", true, func(response clientv3.WatchResponse) {
		for _, event := range response.Events {
			s, ok := sync_controller.ControllerC[string(event.Kv.Key)]
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

}
