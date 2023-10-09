// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GGame is the golang structure for table g_game.
type GGame struct {
	Code            int         `json:"code"            description:"游戏编号"`
	Name            string      `json:"name"            description:"名称"`
	Status          int         `json:"status"          description:"状态0：关闭，1开始"`
	CreatedAt       *gtime.Time `json:"createdAt"       description:"创建时间"`
	StartTime       string      `json:"startTime"       description:"开始时间"`
	EndTime         string      `json:"endTime"         description:"结束时间"`
	TotalIssue      int         `json:"totalIssue"      description:"一共有多少期"`
	IntervalSeconds int         `json:"intervalSeconds" description:"每期间隔秒数"`
	TypeCode        int         `json:"typeCode"        description:"类型code"`
	Sort            int         `json:"sort"            description:"排序"`
	CategoryCode    int         `json:"categoryCode"    description:"分类"`
}
