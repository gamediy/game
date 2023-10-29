// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Payment is the golang structure of table p_payment for DAO operations like Where/Data.
type Payment struct {
	g.Meta   `orm:"table:p_payment, do:true"`
	Merchant interface{} // 商户
	Key      interface{} //
	Name     interface{} //
	Status   interface{} //
	Url      interface{} //
	Code     interface{} // 编号
}
