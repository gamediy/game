// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Game is the golang structure for table game.
type Game struct {
	Code            int         `json:"code"            description:"游戏编号"`
	Name            string      `json:"name"            description:"名称"`
	Status          int         `json:"status"          description:"状态0：关闭，1开始"`
	CreatedAt       *gtime.Time `json:"createdAt"       description:"创建时间"`
	StartTime       string      `json:"startTime"       description:"开始时间"`
	EndTime         string      `json:"endTime"         description:"结束时间"`
	TotalIssue      int         `json:"totalIssue"      description:"一共有多少期"`
	IntervalSeconds int         `json:"intervalSeconds" description:"每期间隔秒数"`
	GameType        int         `json:"gameType"        description:"code"`
	Sort            int         `json:"sort"            description:"排序"`
	CategoryCode    int         `json:"categoryCode"    description:"分类"`
	Number          string      `json:"number"          description:""`
	Length          int         `json:"length"          description:"号码长度"`
	RandomRange     int         `json:"randomRange"     description:"随机范围"`
	CloseSeconds    int         `json:"closeSeconds"    description:"封盘秒数"`
	PlayType        int         `json:"playType"        description:"玩法类型"`
	ThirdId         int         `json:"thirdId"         description:""`
	ThirdGameCode   int         `json:"thirdGameCode"   description:""`
}
