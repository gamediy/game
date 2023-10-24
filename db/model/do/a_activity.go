// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AActivity is the golang structure of table a_activity for DAO operations like Where/Data.
type AActivity struct {
	g.Meta  `orm:"table:a_activity, do:true"`
	Id      interface{} //
	Title   interface{} //
	Type    interface{} // 类型
	Content interface{} //
	Images  interface{} //
	Status  interface{} //
	Event   interface{} // 事件
	Rule    interface{} //
}
