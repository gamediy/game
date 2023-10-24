// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Words is the golang structure of table s_words for DAO operations like Where/Data.
type Words struct {
	g.Meta    `orm:"table:s_words, do:true"`
	Id        interface{} //
	K         interface{} //
	En        interface{} //
	Zh        interface{} //
	Ja        interface{} //
	Group     interface{} //
	Desc      interface{} //
	CreatedAt *gtime.Time //
}
