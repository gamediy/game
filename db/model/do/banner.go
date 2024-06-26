// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Banner is the golang structure of table c_banner for DAO operations like Where/Data.
type Banner struct {
	g.Meta    `orm:"table:c_banner, do:true"`
	Id        interface{} //
	Title     interface{} //
	Image     interface{} //
	Link      interface{} //
	Desc      interface{} //
	Sort      interface{} //
	Status    interface{} // 1open 2close
	CreatedAt *gtime.Time //
}
