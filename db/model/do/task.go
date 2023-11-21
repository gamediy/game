// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Task is the golang structure of table a_task for DAO operations like Where/Data.
type Task struct {
	g.Meta         `orm:"table:a_task, do:true"`
	Id             interface{} // 编号
	TaskCategoryId interface{} // 分类编号
	Title          interface{} // 标题
	LevelNum       interface{} // 关卡数
	Status         interface{} // 状态0：未开始，1:进行中，2：结束
	Type           interface{} // 类型   1:签到
	StartTime      *gtime.Time // 开始时间
	EndTime        *gtime.Time // 结束时间
	TimeType       interface{} // 1:每天，2：每周，3每月,4:每小时 99:永久
	GiveType       interface{} // 送取类型    1,所有用户,2 VIP,3,userId,4,level
	GiveValue      interface{} // 数据
	TimeValue      interface{} //
	Remark         interface{} //
}
