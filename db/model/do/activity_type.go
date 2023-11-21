// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ActivityType is the golang structure of table a_activity_type for DAO operations like Where/Data.
type ActivityType struct {
	g.Meta `orm:"table:a_activity_type, do:true"`
	Type   interface{} //
	Name   interface{} //
	Status interface{} //
	Rule   interface{} // 规则
	Remark interface{} // 说明
}
