// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TBet is the golang structure of table t_bet for DAO operations like Where/Data.
type TBet struct {
	g.Meta       `orm:"table:t_bet, do:true"`
	Id           interface{} //
	Uid          interface{} //
	Account      interface{} //
	GameTitle    interface{} // 自己的
	GameCode     interface{} // 自己的
	ThirdId      interface{} // 三方ID
	ThirdName    interface{} // 三方名称
	ThirdData    interface{} // 三方数据
	Amount       interface{} // 下注额
	ThirdOrderNo interface{} // 三方订单号
	Status       interface{} // 0:下注成功，1：中奖，2：未中奖，3：未结算
	CalcTime     *gtime.Time // 结算时间
	CreatedAt    *gtime.Time //
}
