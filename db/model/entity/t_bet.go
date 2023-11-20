// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TBet is the golang structure for table t_bet.
type TBet struct {
	Id           int64       `json:"id"           description:""`
	Uid          int64       `json:"uid"          description:""`
	Account      string      `json:"account"      description:""`
	GameTitle    string      `json:"gameTitle"    description:"自己的"`
	GameCode     int         `json:"gameCode"     description:"自己的"`
	ThirdId      int         `json:"thirdId"      description:"三方ID"`
	ThirdName    string      `json:"thirdName"    description:"三方名称"`
	ThirdData    string      `json:"thirdData"    description:"三方数据"`
	Amount       float64     `json:"amount"       description:"下注额"`
	ThirdOrderNo int64       `json:"thirdOrderNo" description:"三方订单号"`
	Status       int         `json:"status"       description:"0:下注成功，1：中奖，2：未中奖，3：未结算"`
	CalcTime     *gtime.Time `json:"calcTime"     description:"结算时间"`
	CreatedAt    *gtime.Time `json:"createdAt"    description:""`
}
