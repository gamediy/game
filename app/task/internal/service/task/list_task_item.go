package task

import (
	"context"
	"fmt"
	"game/app/task/api/task/task"
	"game/db/dao"
	"game/db/model/entity"
)

func ListTaskItem(ctx context.Context, req *task.ListTaskItemReq) (*task.ListTaskItemRes, error) {
	res := &task.ListTaskItemRes{}
	res.List = make([]*task.TaskItem, 0)
	_ = dao.TaskItem.Ctx(ctx).Fields("reward_rule,id task_item_id").Scan(&res.List, "task_id", req.TaskId)

	taskItemIds := make([]int64, 0)
	for _, i := range res.List {
		taskItemIds = append(taskItemIds, i.TaskItemId)
		i.TaskId = req.TaskId
	}

	taskUser := make([]*entity.TaskUser, 0)
	if err := dao.TaskUser.Ctx(ctx).
		Where("uid = ? and task_id = ?", req.Uid, req.TaskId).
		WhereIn("task_item_id", taskItemIds).
		Scan(&taskUser); err != nil {
		return nil, err
	}
	if len(taskUser) != 0 {
		for _, i := range taskUser {
			for _, j := range res.List {
				if i.TaskItemId == j.TaskItemId {
					j.Status = fmt.Sprint(i.Status)
				}
			}
		}
	}
	return res, nil
}
