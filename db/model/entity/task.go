// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Task is the golang structure for table task.
type Task struct {
	Id             int64       `json:"id"             description:"编号"`
	TaskCategoryId int64       `json:"taskCategoryId" description:"分类编号"`
	Title          string      `json:"title"          description:"标题"`
	LevelNum       int         `json:"levelNum"       description:"关卡数"`
	Status         int         `json:"status"         description:"状态0：未开始，1:进行中，2：结束"`
	Type           int         `json:"implement"           description:"类型   1:签到"`
	StartTime      *gtime.Time `json:"startTime"      description:"开始时间"`
	EndTime        *gtime.Time `json:"endTime"        description:"结束时间"`
	TimeType       int         `json:"timeType"       description:"1:每天，2：每周，3每月,4:每小时 99:永久"`
	GiveType       int         `json:"giveType"       description:"送取类型    1,所有用户,2 VIP,3,userId,4,level"`
	GiveValue      string      `json:"giveValue"      description:"数据"`
	TimeValue      int         `json:"timeValue"      description:""`
	Remark         string      `json:"remark"         description:""`
}
