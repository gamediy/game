// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TThird is the golang structure of table t_third for DAO operations like Where/Data.
type TThird struct {
	g.Meta   `orm:"table:t_third, do:true"`
	Id       interface{} //
	Name     interface{} // 名称
	Status   interface{} // 状态
	Key      interface{} // key
	Url      interface{} // 地址
	Merchant interface{} // 商户号
	Token    interface{} // 钱包验证token
	Remark   interface{} //
}
