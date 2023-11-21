// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskItem is the golang structure for table task_item.
type TaskItem struct {
	Id             int64       `json:"id"             description:"编号"`
	Title          string      `json:"title"          description:"标题"`
	LevelNum       int         `json:"levelNum"       description:"级别"`
	Status         int         `json:"status"         description:"状态 0"`
	Type           int         `json:"implement"           description:"类型 1签到"`
	TypeCode       string      `json:"typeCode"       description:"类型值"`
	EndValue       float64     `json:"endValue"       description:"结束值"`
	TimeType       int         `json:"timeType"       description:"时间类型"`
	StartTime      *gtime.Time `json:"startTime"      description:"开始时间"`
	EndTime        *gtime.Time `json:"endTime"        description:"结束时间"`
	RewardRule     string      `json:"rewardRule"     description:"奖利规则"`
	GiveType       int         `json:"giveType"       description:"领取方式   1,所有用户,2 VIP,3,userId,4,level"`
	GiveValue      string      `json:"giveValue"      description:"领取值"`
	CurrentValue   float64     `json:"currentValue"   description:"当前值"`
	StartValue     float64     `json:"startValue"     description:"开始值"`
	TaskCategoryId int64       `json:"taskCategoryId" description:"分类编号"`
	Remark         string      `json:"remark"         description:"备注"`
	TaskId         int64       `json:"taskId"         description:"任务编号"`
}
