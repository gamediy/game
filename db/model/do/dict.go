// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Dict is the golang structure of table s_dict for DAO operations like Where/Data.
type Dict struct {
	g.Meta    `orm:"table:s_dict, do:true"`
	Id        interface{} //
	Title     interface{} //
	K         interface{} //
	V         interface{} //
	Desc      interface{} //
	Group     interface{} //
	Status    interface{} //
	Type      interface{} // 1 文本，2 img
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
