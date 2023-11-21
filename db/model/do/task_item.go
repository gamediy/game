// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskItem is the golang structure of table a_task_item for DAO operations like Where/Data.
type TaskItem struct {
	g.Meta         `orm:"table:a_task_item, do:true"`
	Id             interface{} // 编号
	Title          interface{} // 标题
	LevelNum       interface{} // 级别
	Status         interface{} // 状态 0
	Type           interface{} // 类型 1签到
	TypeCode       interface{} // 类型值
	EndValue       interface{} // 结束值
	TimeType       interface{} // 时间类型
	StartTime      *gtime.Time // 开始时间
	EndTime        *gtime.Time // 结束时间
	RewardRule     interface{} // 奖利规则
	GiveType       interface{} // 领取方式   1,所有用户,2 VIP,3,userId,4,level
	GiveValue      interface{} // 领取值
	CurrentValue   interface{} // 当前值
	StartValue     interface{} // 开始值
	TaskCategoryId interface{} // 分类编号
	Remark         interface{} // 备注
	TaskId         interface{} // 任务编号
}
