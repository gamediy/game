// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskUser is the golang structure for table task_user.
type TaskUser struct {
	Id             int64       `json:"id"             description:""`
	Uid            int64       `json:"uid"            description:"用户ID"`
	Account        string      `json:"account"        description:"账号"`
	Status         int         `json:"status"         description:"：未开始，1进行中，2 已完成"`
	CurrentValue   float64     `json:"currentValue"   description:"当前值"`
	StartTime      *gtime.Time `json:"startTime"      description:"开始时间"`
	TaskId         int64       `json:"taskId"         description:"任务ID"`
	TaskItemId     int64       `json:"taskItemId"     description:""`
	TaskCategoryId int64       `json:"taskCategoryId" description:""`
	EndTime        *gtime.Time `json:"endTime"        description:"结束时间"`
	GiveTime       *gtime.Time `json:"giveTime"       description:"领取时间"`
	TypeCode       string      `json:"typeCode"       description:"类型编号"`
	Type           int         `json:"type"           description:"类型"`
	StartValue     float64     `json:"startValue"     description:"开始值"`
	EndValue       float64     `json:"endValue"       description:"结束值"`
	LevelNum       int         `json:"levelNum"       description:"关卡数"`
	Title          string      `json:"title"          description:"标题"`
	SuccessTime    *gtime.Time `json:"successTime"    description:"成功时间"`
}
