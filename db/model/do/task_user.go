// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskUser is the golang structure of table a_task_user for DAO operations like Where/Data.
type TaskUser struct {
	g.Meta         `orm:"table:a_task_user, do:true"`
	Id             interface{} //
	Uid            interface{} // 用户ID
	Account        interface{} // 账号
	Status         interface{} // ：未开始，1进行中，2 已完成
	CurrentValue   interface{} // 当前值
	StartTime      *gtime.Time // 开始时间
	TaskId         interface{} // 任务ID
	TaskItemId     interface{} //
	TaskCategoryId interface{} //
	EndTime        *gtime.Time // 结束时间
	GiveTime       *gtime.Time // 领取时间
	TypeCode       interface{} // 类型编号
	Type           interface{} // 类型
	StartValue     interface{} // 开始值
	EndValue       interface{} // 结束值
	LevelNum       interface{} // 关卡数
	Title          interface{} // 标题
	SuccessTime    *gtime.Time // 成功时间
}
