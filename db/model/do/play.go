// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Play is the golang structure of table g_play for DAO operations like Where/Data.
type Play struct {
	g.Meta        `orm:"table:g_play, do:true"`
	Code          interface{} // 玩法编号
	Name          interface{} // 玩法名
	Status        interface{} // 状态
	Odds          interface{} // 赔率
	Probabilities interface{} // 概率
	PlayType      interface{} // 玩法类型
	BetMin        interface{} // 最小投注金额
	BetMax        interface{} // 最大投注金额
	Number        interface{} //
}
