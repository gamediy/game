package implement

import (
	"context"
	"game/core/wallet"
	"game/db/dao"
	"game/db/model/entity"
	"game/model"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

func SignIn(ctx context.Context, task model.TaskItem) (err error) {
	taskInfo := &entity.Task{}
	_ = dao.Task.Ctx(ctx).Scan(&taskInfo, "id =1")
	if taskInfo.Status == 0 {
		return
	}
	now := gtime.Now()
	if !(now.After(taskInfo.StartTime) && now.Before(taskInfo.EndTime)) {
		return
	}

	//  get last
	lastTaskUser := &entity.TaskUser{}
	_ = dao.TaskUser.Ctx(ctx).
		Where("uid", task.Uid).
		Order("id desc").
		Limit(1).
		Scan(&lastTaskUser)
	if lastTaskUser.Id == 0 {
		taskItems, err := dao.TaskItem.ListItems(ctx, taskInfo.Id)
		if err != nil {
			return err
		}
		if err = doSignIn(ctx, taskInfo, taskItems[0], task); err != nil {
			return err
		}
	} else {
		// If already signed today, will not handle.
		if lastTaskUser.SuccessTime.StartOfDay().Equal(now.StartOfDay()) {
			return
		}
		taskItems, err := dao.TaskItem.ListItems(ctx, taskInfo.Id)
		if err != nil {
			return err
		}
		next := lastTaskUser.LevelNum % len(taskItems)
		if err = doSignIn(ctx, taskInfo, taskItems[next], task); err != nil {
			return err
		}
	}
	return
}

func doSignIn(ctx context.Context, task *entity.Task, taskItem *entity.TaskItem, userInfo model.TaskItem) (err error) {
	x := wallet.Balance{}
	x.TransCode = wallet.SignIn
	x.Uid = userInfo.Uid
	x.Amount = gconv.Float64(taskItem.RewardRule)
	now := gtime.Now()
	if err = x.Update(ctx, func(ctx context.Context, tx gdb.TX) error {
		taskUser := entity.TaskUser{
			Uid:            userInfo.Uid,
			Account:        userInfo.Account,
			Status:         2,
			TaskId:         task.Id,
			TaskItemId:     taskItem.TaskId,
			TaskCategoryId: task.TaskCategoryId,
			StartTime:      now,
			EndTime:        now,
			GiveTime:       now,
			TypeCode:       taskItem.TypeCode,
			Type:           taskItem.Type,
			LevelNum:       taskItem.LevelNum,
			Title:          taskItem.Title,
			SuccessTime:    now,
		}
		if _, err = tx.Model(dao.TaskUser.Table()).Insert(&taskUser); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return
	}
	return
}
