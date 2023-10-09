// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GLotteryOrder is the golang structure of table g_lottery_order for DAO operations like Where/Data.
type GLotteryOrder struct {
	g.Meta     `orm:"table:g_lottery_order, do:true"`
	OrderNo    interface{} // 编号
	Uid        interface{} // 用户编号
	Account    interface{} // 用户账号
	CreatedAt  *gtime.Time // 时间
	GameCode   interface{} // 游戏编号
	GameName   interface{} // 游戏名称
	Amount     interface{} // 投注金额
	Profit     interface{} // 赢利
	CalcAt     *gtime.Time // 结算时间
	Status     interface{} // 状态 0 撤单，1:投注成功，2：中奖，3：未中奖
	Odds       interface{} // 赔率
	Fee        interface{} // 手续费
	OpenResult interface{} // 开奖结果
	PlayCode   interface{} // PlayId
	GameIssue  interface{} // 期号
	Pid        interface{} //
	ParentPath interface{} //
}
