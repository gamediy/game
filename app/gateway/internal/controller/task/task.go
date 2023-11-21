package task

import (
	"context"
	"game/app/gateway/internal/controller"
	tasksvc "game/app/gateway/internal/svc/task"
	"game/app/gateway/internal/ws"
	"game/app/task/api/task/task"
	taskevent "game/consts/event/task"
	"game/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func Init() {
	controller.Ctrl[taskevent.ListTaskItem] = listTaskItem
}

func listTaskItem(ctx context.Context, wsclient *ws.Client, query g.Map) (*model.WsMessage, error) {
	res, err := tasksvc.ListTaskItem(ctx, &task.ListTaskItemReq{
		TaskId: gconv.Int64(query["taskId"]),
		Uid:    wsclient.UserInfo.Uid,
	})
	return &model.WsMessage{
		Event: model.WrapEventResponse(taskevent.ListTaskItem),
		Body:  model.WrapMessage(res, err),
	}, nil
}
