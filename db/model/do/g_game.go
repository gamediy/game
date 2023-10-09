// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GGame is the golang structure of table g_game for DAO operations like Where/Data.
type GGame struct {
	g.Meta          `orm:"table:g_game, do:true"`
	Code            interface{} // 游戏编号
	Name            interface{} // 名称
	Status          interface{} // 状态0：关闭，1开始
	CreatedAt       *gtime.Time // 创建时间
	StartTime       interface{} // 开始时间
	EndTime         interface{} // 结束时间
	TotalIssue      interface{} // 一共有多少期
	IntervalSeconds interface{} // 每期间隔秒数
	TypeCode        interface{} // 类型code
	Sort            interface{} // 排序
	CategoryCode    interface{} // 分类
}
