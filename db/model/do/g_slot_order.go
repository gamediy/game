// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GSlotOrder is the golang structure of table g_slot_order for DAO operations like Where/Data.
type GSlotOrder struct {
	g.Meta     `orm:"table:g_slot_order, do:true"`
	OrderNo    interface{} // 编号
	Uid        interface{} // 用户编号
	Account    interface{} // 用户账号
	CreatedAt  *gtime.Time // 时间
	Amount     interface{} // 投注金额
	Profit     interface{} // 赢利
	Status     interface{} // 状态 0 撤单，1:投注成功，2：中奖，3：未中奖
	Odds       interface{} // 赔率
	Fee        interface{} // 手续费
	BetResult  interface{} // 下注结果
	PlayCode   interface{} // PlayId
	Pid        interface{} //
	ParentPath interface{} //
}
