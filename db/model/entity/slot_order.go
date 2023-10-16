// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SlotOrder is the golang structure for table slot_order.
type SlotOrder struct {
	OrderNo    int64       `json:"orderNo"    description:"编号"`
	Uid        int64       `json:"uid"        description:"用户编号"`
	Account    string      `json:"account"    description:"用户账号"`
	CreatedAt  *gtime.Time `json:"createdAt"  description:"时间"`
	Amount     float64     `json:"amount"     description:"投注金额"`
	Profit     float64     `json:"profit"     description:"赢利"`
	Status     int         `json:"status"     description:"状态 0 撤单，1:投注成功，2：中奖，3：未中奖"`
	Odds       float64     `json:"odds"       description:"赔率"`
	Fee        float64     `json:"fee"        description:"手续费"`
	BetResult  string      `json:"betResult"  description:"下注结果"`
	PlayCode   int         `json:"playCode"   description:"PlayId"`
	Pid        int64       `json:"pid"        description:""`
	ParentPath string      `json:"parentPath" description:""`
	GameCode   int         `json:"gameCode"   description:""`
}
