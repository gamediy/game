// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Menu is the golang structure of table s_menu for DAO operations like Where/Data.
type Menu struct {
	g.Meta    `orm:"table:s_menu, do:true"`
	Id        interface{} //
	Pid       interface{} //
	Name      interface{} //
	NameZh    interface{} //
	NameJa    interface{} //
	Url       interface{} //
	Desc      interface{} //
	Sort      interface{} //
	Type      interface{} // 1 group 2 menu
	Status    interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
