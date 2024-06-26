// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure of table s_role for DAO operations like Where/Data.
type Role struct {
	g.Meta    `orm:"table:s_role, do:true"`
	Id        interface{} //
	Name      interface{} //
	NameZh    interface{} //
	NameJa    interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	Class     interface{} //
}
