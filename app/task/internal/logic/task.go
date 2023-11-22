package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"game/app/task/internal/logic/implement"
	"game/model"
	"game/utility/utils/xetcd"
	"github.com/gogf/gf/v2/frame/g"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func TaskRun(ctx context.Context) {
	go xetcd.Watch(ctx, "/task/", true, func(response clientv3.WatchResponse) {
		for _, event := range response.Events {
			model := model.TaskItem{}
			err := json.Unmarshal(event.Kv.Value, &model)
			if err != nil {
				g.Log().Error(ctx, err)
				continue
			}
			if model.Type == 1 {
				implement.SignIn(ctx, model)
			}
		}
		fmt.Println(response)
	})
	return
}
